---
title: "Counter-Strike code leaked: should you worry? What if your code leaks?"
type: post
description: "In this blog post I talk about what it takes to delivery software securely to the general public."
date: "2020-04-22"
image: "/img/posts/cs-security/cs.png"
hashtags: "counterstrike,cs,valve,game,security,leak,software"
---
[![Counter-Strike code leaked!](/img/posts/cs-security/cs.png)](https://blog.counter-strike.net)

The [source code](https://en.wikipedia.org/wiki/Source_code) for the game [Counter-Strike: GO](https://www.igdb.com/games/counter-strike-global-offensive) has leaked today. I'm sad to see many concerned users decided not to play the game anymore, for now, worried about [Remote Code Execution](https://en.wikipedia.org/wiki/Arbitrary_code_execution) (RCE). [Valve](https://www.valvesoftware.com/) told them they've reviewed the code, and there's no reason for users to worry.

Here's I discuss what everyone should be doing about distributing software to the general public safely, and why I trust this statement:

{{< tweet 1253075594901774336 >}}
{{< tweet 1252961862058205184 >}}

The reason why people are concerned is because bad practices such as [security through obscurity](https://en.wikipedia.org/wiki/Security_through_obscurity) or blindly trusting "secure servers" to invoke remote code are widespread in the software industry.

If you have an iPhone, you probably know applications are sandboxed. In essence, files of a given application aren't blindly shared with another. You have some permissions settings and fine controls. Some people complain this makes the iOS operating system really closed, but this is actually one of its greatest strengths.

## Who gets it right
Apple is definitely [a leader](https://support.apple.com/en-us/HT210897) when talking about respecting user security and privacy, especially in the smartphone ecosystem. It's [Differential Privacy](https://www.apple.com/privacy/docs/Differential_Privacy_Overview.pdf) approach leverages processing private user data on their own devices . Its [Privacy-Preserving Contact Tracing](https://www.apple.com/covid19/contacttracing/) protocol jointly developed with Google to fight [coronavirus (COVID-19)](https://simple.wikipedia.org/wiki/Coronavirus_disease_2019) has user privacy as a first-class citizen. **It is no different when we talk about software distribution.**

Application sandboxing appeared in iOS since its very beginning. Android is also [moving towards this direction ](https://developer.android.com/training/data-storage#scoped-storage) as users are getting more concerned over the topic.

[FreeBSD](https://www.freebsd.org) and Linux distributions have a history of mirroring packages in public mirrors safely thanks to package signature. By the way, FreeBSD has the concept of jails, that allows you to isolate processes and what computer resources they can access. Something useful to mitigate attack vectors.

Google also has been doing great work bringing safer machines to the general public with its [Chrome OS](https://www.google.com/chromebook/chrome-os/) too.

<div id="amzn-assoc-ad-650fc264-1310-445e-ba75-6bf6849026cc"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=650fc264-1310-445e-ba75-6bf6849026cc"></script>

## Traditional computers
However, the situation for "general" computers is not good, with [sandboxing](https://en.wikipedia.org/wiki/Sandbox_(computer_security)) still flourishing. It is an essential responsibility of any company that develops and publishes software for the public at large to enforce boundaries and limits to deal with their practically unrestricted trust and risks.

Browsers are, perhaps, the only understood example of a successful sandbox on traditional computers.

## Code signing
Code signing is a technique used to confirm software authorship and to provide an indication of provenance, guaranteeing that a program has not been tampered or corrupted.

If you sign your software, you can:

* Distribute your software over unsecured channels
* Guarantee that the end-user of your software get what was distributed
* Make your builds auditable if [reproducible](https://reproducible-builds.org)

If you don’t, you must have control over all the chain of custody of your application and its communication from the distribution channel to the end-user to assure its security.

## Notarization
[Notarization](https://developer.apple.com/documentation/xcode/notarizing_macos_software_before_distribution) is a process of sending your application to Apple so it can sign that your software is free from malware or malicious components. You can automate this process on your delivery pipeline, so there is no reason why not to use it if you distribute applications to end-users.

## It starts from the firmware
[Jessie Frazelle](https://jess.dev/) from the [Oxide Computer Company](https://oxide.computer) has a great article about [Why open source firmware is important for security](https://blog.jessfraz.com/post/why-open-source-firmware-is-important-for-security/).

## curl | sh: shame on you! and me!
If you use any Unix-like computer and are reading this post, there is a high chance you installed something with:

`$ curl http://example.com/unsafe | sh`

If not created something like this yourself! I'm guilt of this myself.

There [are](https://news.ycombinator.com/item?id=12766049) just [too](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/) many [reasons](https://sandstorm.io/news/2015-09-24-is-curl-bash-insecure-pgp-verified-install) why this is, overall, a bad idea.

### Why even bother signing?
If you don't sign and your server is compromised, your users are at risk, and this is unfair.

You don't want to let your users down by exposing or destroying their data, and you don't want to face legal consequences os lose their trust.

This risk can only be partly mitigated by using HTTPS. However, this lacks a defense in depth or Castle Approach. You want to minimize your attack surface.

### Why even bother code notarizing?
If you notarize your application, your Apple users will likely trust you more and have a smoother user experience when installing your application.

<div id="amzn-assoc-ad-867f8b7d-ea56-4a22-8a65-086e6d53aa5a"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=867f8b7d-ea56-4a22-8a65-086e6d53aa5a"></script>

## Apple's Gatekeeper
Apple tightened up software installation making the operating system more likely to stop users from executing untrusted code. Some people complain that they are closing down the platform and that this is a bad thing, and I must disagree. For me, they are trying to make macOS as safe to use as iOS.

[![Gatekeeper screenshot](/img/posts/cs-security/gatekeeper_small.png)](/img/posts/cs-security/gatekeeper.png)

This is what you see if you only allow downloading from the App Store and try to install something from outside:

[![Failure installing from outside App Store](/img/posts/cs-security/cannot-download_small.png)](/img/posts/cs-security/cannot-download.png)

This is what happens if you try to install an application that wasn't product signed now:

[![SourceTree cannot be installed](/img/posts/cs-security/sourcetree-cannot-be-open_small.png)](/img/posts/cs-security/sourcetree-cannot-be-open.png)

You can quickly get around installing apps from outside the App Store that are product signed:

[![SourceTree cannot be installed](/img/posts/cs-security/go-installer-certificate_small.png)](/img/posts/cs-security/go-installer-certificate.png)

However, things look slightly more hostile to applications that aren't product signed:

[![Atlassian needs to sign SourceTree urgently](/img/posts/cs-security/sourcetree-blocked_small.png)](/img/posts/cs-security/sourcetree-blocked.png)

_[Atlassian](https://www.atlassian.com/) evidently forgot to request an Apple certificate to sign SourceTree._

## Writing an application for macOS
Please notice that Apple has two programs with similar names:

* codesign is for signing code and the like
* productsign is for signing packages and the like

For this subject, productsign is what you need.

* [Distributing Your Mac Apps](https://developer.apple.com/macos/distribution/)
* [Notarizing macOS Software Before Distribution](https://developer.apple.com/documentation/xcode/notarizing_macos_software_before_distribution)

Please don't do the curl | sh workaround if you have a public user base even though I did it in the past. You can do better than me.

Tip: verify macOS .pkg installers with the native `installer` program or with [Suspicious Package](https://www.mothersruin.com/software/SuspiciousPackage/).

### Further reading list

* [Code signing](https://developer.apple.com/support/code-signing)
* [About Code Signing](https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/Introduction/Introduction.html)
* [How to sign your Mac OS X App for Gatekeeper](https://successfulsoftware.net/2012/08/30/how-to-sign-your-mac-os-x-app-for-gatekeeper/)
* [Panic: About Gatekeeper](https://panic.com/blog/about-gatekeeper/)
* [How to use the Apple Product Security PGP Key](https://support.apple.com/en-sg/HT201601)
* [How To Sign macOS PKGs for Deployment with MDM](https://simplemdm.com/certificate-sign-macos-packages/)

## Writing an application for Windows
For Windows, you can request a Windows Code Signing certificate from multiple [Certificate Authorities](https://en.wikipedia.org/wiki/Certificate_authority) (CAs). Make sure you don't make the mistake of getting a certificate for a domain, though. Both rely on [Public-key infrastructure](https://en.wikipedia.org/wiki/Public_key_infrastructure), but are not compatible with each other.

[DigiCert](https://www.digicert.com/) is a reliable partner. I used it in the past. I recommend you get a certificate for the maximum amount of time possible, especially if you don't use Windows regularly.

By the way, if you don’t codesign, your users might get weird security risk message by Windows antiviruses.

### Microsoft documentation

* [Cryptography Tools](https://docs.microsoft.com/en-us/windows/win32/seccrypto/cryptography-tools)
* [SignTool](https://docs.microsoft.com/en-us/windows/win32/seccrypto/signtool)

Time-stamping the signature with a remote server is recommended.

> Time-stamping was designed to circumvent the trust warning that will appear in the case of an expired certificate. In effect, time-stamping extends the code trust beyond the validity period of a certificate. In the event that a certificate has to be revoked due to a compromise, a specific date and time of the compromising event will become part of the revocation record. In this case, time-stamping helps establish whether the code was signed before or after the certificate was compromised.

[Source: Wikipedia article on code signing](https://en.wikipedia.org/wiki/Code_signing#Time-stamping).

## Updating your application
Your updates should be safe too. You should not directly download a binary from a server and replace your file without validation.

Your application must validate if the downloaded files are legit before installing them.

With the [Go](https://www.golang.org/) programming language, I used [equinox.io](https://equinox.io) for managing my updates safely and successfully in the past, and I recommend it. However, you don't need any 3rd party provider to do this.

The essentials are:

* You must have a secure device to generate releases and sign them.
* You should keep your private keys safe.

When I was working on releasing a [CLI](https://en.wikipedia.org/wiki/Command-line_interface) used by hundreds of users, I used to connect to a remote server via [SSH](https://en.wikipedia.org/wiki/Secure_Shell). It was protected with public-key infrastructure through my SSH key + password + [one-time password](https://github.com/google/google-authenticator-libpam) in a secure physical location. Access to both the physical location (okay, I admit! Amazon Web Services datacenters!) and to the server was on a need-to basis.

I'd SSH to it whenever a release was ready, and do it from there, to avoid exposing the private key.

Things I slacked on doing but would be nice to have:

* Verification for revoked keys
* Kill switch to stop users from downgrading/updating the application to unsafe versions, if I ever released something with a security bug
* A second key pair to serve as a backup if the first leaked
* Friendly key rotation (had no rotation)

Things I would do if more users were using it and the perceived risk was greater:

Multiple keys so multiple people could validate an update before they would be allowed to be used by users. I know Apple uses something like this to roll their operating system updates, but I couldn't find a reference. If you know, please let me know to update here.

## Chimera: auto update + run command
Equinox didn't allow me to code sign a package, so I had to find a way around it.

I created what I called a chimera: a binary for my program that seems to be a regular but is actually…

An installer that will automatically download and install a new version, replacing itself, and executing the new command from inside, passing environment variables, on the same current working directory, and piping the standard input, output, and error, and exiting with the very same error code.

Here is the [implementation](https://github.com/henvic/wedeploycli/tree/master/update/internal/installer), if you are curious.

The good thing about it was that I was free of having to use Windows to release a new code signed version each time I decided to do it.

[GitHub's official command line tool](https://cli.github.com) (gh; but not related to the [NodeGH](http://nodegh.io) I used to maintain a long time ago!) does the code signing on every release. If you want it, take a look at [their code](https://github.com/cli/cli/blob/master/.github/workflows/releases.yml).

<div id="amzn-assoc-ad-6763ebda-8a65-4dde-84c8-de9b85de01d7"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=6763ebda-8a65-4dde-84c8-de9b85de01d7"></script>

## Summary
If you are a Valve consumer, I'd say you can trust their words and expect nothing bad is going to happen. While many people resort to obfuscation when working on closed-source software, I'd say Valve probably doesn't or they would already be screwed. In theory, issues might exist and get discovered by bad actors easier but I wouldn't worry about it either.

If you're worried about this, you might consider to buy a dedicated gaming machine to mitigate the risks, and this is a good measure if you can afford.

If you are a software developer and found this post useful, please share your experiences and ideas too!
