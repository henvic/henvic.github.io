---
title: "Environment variables, config, secrets, and globals"
type: post
description: "Server-side applications are heavily relying on environment variables for holding configuration data, and even credentials like token and password. Is this a good idea?"
date: "2021-01-16"
image: "/img/posts/env/airplane.jpg"
hashtags: "unix,env,config,secrets"
---

At some point in time, perhaps with the advent of [The Twelve-Factor App](https://12factor.net/) methodology, we saw new server-side applications moving from a file-based configuration to an environment variable configuration approach.

Many developers like to use environment variables for credentials because they are ephemeral. By relying on them, you might avoid leaking credentials on the web by accidentally checking them on public repositories, or in case of a [directory traversal attack](https://en.wikipedia.org/wiki/Directory_traversal_attack) vulnerability.

While these concerns might be worth considering, relying on environment variables doesn't come without risks.

**There is no magic way to keep secrets, secret without risks.**

Whatever security model you decide to follow and whether you decide to use environment variables or otherwise, you should at least be aware of the points below.

<small>Read also: [Back to basics: Writing an application using Go and PostgreSQL](/posts/go-postgres/) and [You don't need UUID](/posts/uuid/).</small>

<a data-flickr-embed="true" href="https://www.flickr.com/photos/henriquev/42847303541/in/dateposted-public/" title="Sólheimasandur DC-3 Plane Wreck"><img src="https://live.staticflickr.com/1753/42847303541_73c2267ed8_c.jpg" width="800" height="449" alt="Sólheimasandur DC-3 Plane Wreck"></a><script async src="//embedr.flickr.com/assets/client-code.js" charset="utf-8"></script>

## Don't keep secrets on your code repository
Application code should live in a repository separate from configuration data for your environments (production, staging, testing, etc.).

You should also always consider that your code is a public asset available on the Internet, regardless if it is open source or something you never intend to share with others. By having a strong separation of concerns, you can avoid a hassle down the road – regardless if you decide to open source it one day or you need to rotate credentials because someone left your team.

It might be the case that you want to have some sample configuration along with your application code. Don't reuse the same configuration data that gives you access to your environments or subsystems (internal or third-party).

Some external services you use might provide a testing environment, and you might think it is okay to share that because you have nothing to hide. Don't do that! If for nothing else, think about the lost time you may incur if your tests start breaking because people elsewhere (maybe a former employee) might reuse your keys, making you end up with [Too Many Requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) HTTP errors due to rate limiting.

## Reducing risk with cloud services and integrations
Many developers and companies rely on services "in the cloud". One of the biggest things now is Kubernetes.

Let's suppose for a moment you're using Kubernetes [declaratively](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/) – meaning you're maintaining configuration files to define the state of your systems. You might be versioning them with git, and using GitHub for a review/approval pipeline, and Continuous Integration to roll out a new release.

In this case, your attack surface area includes at least:

* Your employees with direct access.
* GitHub systems and their staff.
* Continuous Integration service.
* Third-party services with read-access to your private repositories.

**How to mitigate this risk?**
You'll want to have all your purportedly sensitive data encrypted.

You also want to use a whitelist approach to grant permissions on your systems and audit them periodically.

**What is sensitive?**
Tokens, passwords, credentials, private keys, SSH keys, etc.

Hosts are not sensitive – 'til they are.

Even though not widely used on the public web anymore (thankfully!), a URI scheme

`
scheme://userinfo@host/path?query#fragment
`

might convey a `userinfo` field carrying a password or token – and to trust your memory to never do this is unreliable.

**What if I encrypt everything on the configuration repository?**
It might work great from a security perspective, although it might doom the developers' code review process.

## What about environment variables?
Given all I said above, it might appear that relying on environment variables is indeed the best you can do. Don't rush to conclusions just yet.

First of all, environment variables need to be stored somewhere that isn't ephemeral. The protection you need to take to keep secrets on them safely is similar to the precautions you'd need anyway if using regular configuration files.

However, it is easier to overlook a series of security risks involving environment variables.

### Globals
If you use environment variables for configuration, where do you use them on your code? Is it on the functions you need to use their values? On a configuration package? Either way, you're relying on your application's global state, and your life might get miserable because this is hard to keep track of.

**It is easier to maintain and audit code if you don't have to worry about globals.**

### Child processes
Take this Go code, for example:

```go
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("env")
	cmd.Stdout = os.Stdout
	log.Printf("Running command and waiting for it to finish...")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command finished with error: %v\n", err)
	}
}
```

**What does this code do?**

If you open the code for [Run](https://golang.org/pkg/os/exec/#Cmd.Run), you'll see it's going to call `cmd.Start()`, which, in essence, [fork a new process](https://github.com/golang/go/blob/e0c3ded337e95ded40eb401e7d9e74716e3a445f/src/os/exec/exec.go#L422-L432) using a low-level [os.StartProcess](https://golang.org/pkg/os/#StartProcess) API:

```go
c.Process, err = os.StartProcess(c.Path, c.argv(), &os.ProcAttr{
	Dir:   c.Dir, // Inheriting working directory.
	Files: c.childFiles,
	Env:   addCriticalEnv(dedupEnv(envv)),
	Sys:   c.SysProcAttr,
})
if err != nil {
	c.closeDescriptors(c.closeAfterStart)
	c.closeDescriptors(c.closeAfterWait)
	return err
}
```

By default, `cmd.Env` (slice of key=value pairs) field will be `nil` when not explicitly initialized. In this case, `c.envv()` copies the environment variable from the parent process similarly to passing `cmd.Env = syscall.Environ()` (except on Windows, where it does slightly more).

If you run this [code in the Go Playground](https://play.golang.org/p/UVYbqhK6Auv), you'll get something like this:

```txt
2009/11/10 23:00:00 Running command and waiting for it to finish...
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=84682c8aa3aa

Program exited.
```

You should expect a similar result on any other language, as inheriting environment variables on child process is the desired default behavior when they are not explicitly set. In other words, **environment variables leak easily** when spawning child processes.

Maybe you don't call any external process yourself, but are you sure no library you depend on does it? Are you sure no one will ever really add an external call on future changes?

### A safer way of using environment variables?
Yes, environment variables are globals. The good (?) news is that you can set and unset them during execution.

If you decide to use environment variables for configuration, you might want to read them on your main package and unset their values immediately afterward.

In Go, this might translate into something along the lines of:

```go
package main

// Your imports...

func main() {
	cfg := loadEnv()
	// From now on, you cannot use os.Getenv("CLOUD_PROVIDER_TOKEN") anymore as it was unset.
	if err := application.Run(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
}

// loadEnv variables used on the application configuration,
// and unsets them so that no one can use them directly as globals elsewhere on your code.
func loadEnv() *config.Settings {
	// Defer clean up calls to just after loadEnv returns.
	defer func() {
		os.Unsetenv("CLOUD_PROVIDER_TOKEN")
		os.Unsetenv("TWITTER_TOKEN")
	}()
	return &config.Settings{
		StorageToken: os.Getenv("CLOUD_PROVIDER_TOKEN"),
		TwitterToken: os.Getenv("TWITTER_TOKEN"),
	}
}
```

**Possible complications:**

* A library might depend on an environment variable, and you won't be able to unset it.
* You might end up unsetting something that shouldn't be unset (for argument's sake, consider the `$PATH` environment variable).

Also, who guarantees that new contributions are going to follow this pattern consistently? In the end, it'd likely just add to the burden code reviewers already have, right?

Static code analysis might do this automagically, but who has the time to write one to do so and decide how to deal with the number of possibilities regarding false positives, false negatives, opt-in, opt-out approaches when you could avoid this hassle by using a configuration file?

> An alternative to this is to make sure you purge any secrets from your child processes environment variables when initializing them, but this might not be feasible everywhere.

### Use environment variables to point to where secrets live
[Google](https://cloud.google.com/docs/authentication/getting-started), [AWS](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html), [Microsoft](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authorization) Software Development Kits (SDK) all provide a way for you to authenticate with their client libraries using a credentials file.

For example, Google uses the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to point to a credentials file by default. Important: this is not security by obfuscation because the likelihood of a credentials file leaking somewhere is smaller than an environment variable leaking.

**Why is that?** The most overlooked problem with credentials in environment variables from the point of view of security is that it is too easy to leak in several ways, such as via a debugging probe or child process.

### Please, no secrets on flags!

```shell
$ ps u
USER     PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND
henvic 90558   0.0  0.4  5302364  63552 s007  S+   Sat03AM   5:31.65 app serve -token th1s-is-4wful
```

Running a virtual server is cheap nowadays, so most likely, you are not using a shared hosting system where many tenants share a single machine without proper isolation unless you live in the past.

Back in the 1990s to 2000s, it was quite common to pay for a multi-tenant machine on a web hosting service where you could create a database, PHP or Perl application, and so on, controlling everything via [cPanel](https://en.wikipedia.org/wiki/CPanel) and SSH.

Serious security issues appeared day and night because the separation between users was still primitive. Any security bug used to have vast consequences (a malicious attacker would usually pay for an account on your web hosting company and run exploits from the host machine).

Containers, virtual machines, and Unix jails mean that security issues that haunted multitenancy systems in the past are mostly gone. Still, we continue to use third-party monitoring tools to watch our operating system metrics and processes and page or ping us in case of problems. If environment variables might leak with these, flags are even way more likely to leak as well.

## Conclusion
Consider the trade-offs of each approach and choose what works for you carefully. I like to keep things simple and have the configuration for my applications living in a separate and well-protected repository.

* If you choose environment variables, try to minimize the risk of unintended globals propagation.
* If you choose files, make sure you don't expose them to the outside world unintentionally, and set & check for proper file-system [permissions](https://en.wikipedia.org/wiki/File-system_permissions) (or access rights).

Oh, if you're using Kubernetes, you can
[use secrets as files from a pod](https://kubernetes.io/docs/concepts/configuration/secret/#using-secrets-as-files-from-a-pod).

## References

* [OWASP Cheat Sheet: Authentication](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
* [Modern Operating Systems, Andrew S. Tanenbaum](https://amzn.to/3swOApT)
* [xkcd: Duty Calls](https://xkcd.com/386/)

<small>If you click and buy any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/).</small>

{{< tweet user=henriquev id=1351293266025709568 >}}
