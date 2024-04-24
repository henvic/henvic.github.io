---
title: "The ecosystem of the Go programming language"
type: post
description: "Go is one of the most prominent general-purpose programming languages nowadays. Google, Apple, Microsoft, Amazon, and Adobe, to name a few, have been using the language extensively. It’s the language of choice behind multiple cloud computing projects such as Kubernetes, and it’s steadily expanding towards numerous areas of software development. In this article, you'll find resources to learn about Go and its ecosystem."
date: "2021-03-22"
image: "/img/posts/go/go-logo-blue.png"
hashtags: "golang"
---
Go is one of the most prominent general-purpose programming languages nowadays.
Google, Apple, Microsoft, Amazon, and Adobe, [to name a few](https://github.com/golang/go/wiki/GoUsers), have been using the language extensively.
It's the language of choice behind multiple cloud computing projects such as [Kubernetes](https://kubernetes.io/), and it's steadily expanding towards numerous areas of software development.
In this article, you'll find resources to learn about Go and its ecosystem.

If you want to see how people are using Go, check out the [Go Developer Survey 2020 Results](https://blog.golang.org/survey2020-results).

<a href="https://golang.org/"><img src="/img/posts/go/go-logo-blue.svg" width="200" alt="Gopher"></a>

{{< table_of_contents >}}

<small>Read also: [Back to basics: Writing an application using Go and PostgreSQL](/posts/go-postgres/) and [My Go mistakes](/posts/my-go-mistakes/).</small>

| Fact sheet |
| ------------- |-------------|
| Paradigm | Multi-paradigm: mostly imperative, and concurrent. |
| Designed by | [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer), [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike), and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson). [Russ Cox](https://swtch.com/~rsc/) and [Ian Lance Taylor](https://airs.com/ian/) quickly joined them. |
| Sponsored by | [Google](https://www.google.com/) |
| Dates | Design began in late 2007; Publicly Announced [November 10, 2009](https://opensource.googleblog.com/2009/11/hey-ho-lets-go.html) by the Go Team. |
| Influenced by | C, Pascal, Modula-2, Oberon-2, CSP, Occam, Newsqueak, Limbo, Alef, BCPL, Smalltalk, APL, etc. |
| Typing | Inferred, static, strong, structural. |
| Known for being | Modern, readable, concise, garbage-collected, fast to compile, statically linked. |
| Site | [golang.org](https://golang.org/) |
| Release cycle | [Major release every 6 months](https://github.com/golang/go/wiki/Go-Release-Cycle) |
| Essentials | [Effective Go](https://golang.org/doc/effective_go), [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments), [Specification](https://golang.org/ref/spec), and [FAQ](https://golang.org/doc/faq) |

## Hello World in Go
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
```

[Run this code in The Go Playground](https://play.golang.org/)

<script type="text/javascript">
amzn_assoc_placement = "adunit0";
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_linkid = "dc91c08424afa81183b8cf582406ddac";
amzn_assoc_search_bar = "true";
amzn_assoc_title = "Go";
amzn_assoc_asins = "0134190440,1617291781,1617297003,B001QI4UJW";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

{{< tweet user=henriquev id=1252360422683897856 >}}

## A stable platform

Go 1 has a compatibility promise at the source level with the aim of building a stable platform for the growth of program and projects made with the language:

> It is intended that programs written to the Go 1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification. At some indefinite point, a Go 2 specification may arise, but until that time, Go programs that work today should continue to work even as future "point" releases of Go 1 arise (Go 1.1, Go 1.2, etc.).
> 
> [Go 1 and the Future of Go Programs](https://golang.org/doc/go1compat)

The _go1compat_ promise by the Go Team is worth highlighting in a world where most languages swiftly roll major versions with breaking changes and ever-changing APIs almost yearly. I highly suggest you read their compatibility document to understand the expectations you can have and how they might apply to topics such as bugs, spec errors, security issues, the use of unkeyed struct literals, and the [unsafe](https://pkg.go.dev/unsafe) package.

## A simple language
Consulting the [spec](https://golang.org/ref/spec), you can find out that Go has only 25 [reserved keywords](https://en.wikipedia.org/wiki/Reserved_word).

```text
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

You'll also discover it is a concise language with regular syntax. In fact, one with the goals the Go team had when designing it was [to increase the productivity of software engineering at scale](https://talks.golang.org/2012/splash.article) at Google.

| Influence |
| ------------- |-------------|
| Statement and expression syntax | C |
| Declaration syntax | Pascal |
| Packages | Modula-2, Oberon-2 |
| Concurrency | CSP, Occam, Newsqueak, Limbo, Alef |
| The semicolon rule | BCPL |
| Methods | Smalltalk |
| <-, := | Newsqueak |
| iota | APL |
| Lessons good and bad | C++, C#, Java, JavaScript, LISP, Python, Scala, etc. |

Source: [Hello, Gophers!](https://talks.golang.org/2014/hellogophers.slide)

## Learning
[Effective Go](https://golang.org/doc/effective_go), [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments), and [Specification](https://golang.org/ref/spec) are must-read documents if you want to be serious about Go.
I always recommend anyone using the language to follow the recommendations of **Effective Go** and **Code Review Comments** for the sake of consistency – even when not 100% sold to them.

In the end of 2019, the Go Team launched [go.dev](https://go.dev/) website, a *hub for Go users providing centralized and curated resources from across the Go ecosystem*. Go to [learn.go.dev](https://learn.go.dev/) to see a list of **Learning Resources** with a range of tactics and focus subjects.

I started learning Go six years ago and mainly used the **official documentation**, **A Tour of Go**, and **Go by Example** to learn the language.
No matter if you're an experienced developer or a newcomer to programming, Go is a fun language to learn, and the community is friendly and respectful.

### Quick start guides
* [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
* [A Tour of Go](https://tour.golang.org/)
* [Go by Example](https://gobyexample.com/)
* [Learn Go: The complete course](https://www.karanpratapsingh.com/blog/learn-go-the-complete-course)

### Books
* [The Go Programming Language](https://amzn.to/2ONS33T) ([website](https://www.gopl.io/)) was written by [Alan A. A. Donovan](https://alandonovan.net/) and [Brian W. Kernighan](https://www.cs.princeton.edu/~bwk/) (co-author of [The C Programming Language](https://amzn.to/3qFOpX6) with [Dennis Ritchie](https://www.bell-labs.com/usr/dmr/www/) and [The Unix Programming Environment](https://amzn.to/3rJTGyk) with [Rob Pike](http://www.herpolhode.com/rob/))
* For a concise and comprehensive guide, try [Go in Action](https://www.manning.com/books/go-in-action) ([online version](https://livebook.manning.com/book/go-in-action/) and [on Amazon](https://amzn.to/2OfIbjR)) by [Bill Kennedy](https://twitter.com/goinggodotnet) and [Brian Ketelsen](https://www.brian.dev/) (known for co-organizing GopherCon)
* [Learning Go: An Idiomatic Approach to Real-World Go Programming](https://amzn.to/2QmB89s) by [Jon Bodner](https://twitter.com/jonbodner) ([published by O'Reilly](https://www.oreilly.com/library/view/learning-go/9781492077206/))
* [Practical Go Lessons](https://www.practical-go-lessons.com/) by [Maximilien Andile](https://twitter.com/maximilienald)

<script type="text/javascript">
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_design = "enhanced_links";
amzn_assoc_asins = "0134190440";
amzn_assoc_placement = "adunit";
amzn_assoc_linkid = "9da8cfa3d8c20690df5251108220866d";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

<script type="text/javascript">
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_design = "enhanced_links";
amzn_assoc_asins = "1617291781";
amzn_assoc_placement = "adunit";
amzn_assoc_linkid = "020cf6e02222213b0e94a5466dc12cf6";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

### Articles, tutorials, talks, opinions, etc.
* [Go wiki](https://github.com/golang/go/wiki/)
* [Go talks](https://talks.golang.org/)
* [The Go Memory Model](https://golang.org/ref/mem)
* [Data Race Detector](https://golang.org/doc/articles/race_detector)
* [research!rsc](https://research.swtch.com/) by Russ Cox
* [Articles](https://peter.bourgon.org/articles/) and [talks](https://peter.bourgon.org/talks/) by [Peter Bourgon](https://peter.bourgon.org/)
* [Practical Go](https://dave.cheney.net/practical-go) and [High Performance Go](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html) by [Dave Cheney](https://dave.cheney.net/)
* [Writing Accessible Go 2018 GopherCon talk](https://www.youtube.com/watch?v=cVaDY0ChvOQ) ([slides](https://www.juliaferraioli.com/presos/writing-accessible-go/)) by [Julia Ferraioli](https://www.juliaferraioli.com/)
* [Rust vs. Go: Why They’re Better Together](https://thenewstack.io/rust-vs-go-why-theyre-better-together/)
* [GopherCon 2018: Bryan C. Mills - Rethinking Classical Concurrency Patterns](https://www.youtube.com/watch?v=5zXAHh5tJqQ) ([slides](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view) and [transcript](https://about.sourcegraph.com/go/gophercon-2018-rethinking-classical-concurrency-patterns/))
* [TutorialEdge Go Development](https://tutorialedge.net/golang/) and [course](https://tutorialedge.net/course/golang/)
* [Darker Corners of Go](https://rytisbiel.com/2021/03/06/darker-corners-of-go/) by [Rytis Biel](https://rytisbiel.com/)

### Training and workshops
* Many conferences, such as [GopherCon](https://gophercon.com/) offers hands-on workshops before or after the main event.
* [Ultimate Go](https://www.ardanlabs.com/ultimate-go/) is a workshop by [Bill Kennedy](https://twitter.com/goinggodotnet)'s [Ardan Labs](https://www.ardanlabs.com/) consulting company and praised by many!

## Community
Popular communications channels about the language include:

* [Go Forum](https://forum.golangbridge.org/)
* [Gophers Slack](https://invite.slack.golangbridge.org/)
* [golang-nuts general discussion list](https://groups.google.com/g/golang-nuts)
* [#golang](https://twitter.com/hashtag/golang) hashtag @ Twitter
* [Reddit's Go community](https://reddit.com/r/golang)
* [#go-nuts](https://freenode.logbot.info/go-nuts/) IRC channel @ [Libera Chat](https://libera.chat/)) (you'll need an IRC client or webchat like [IRCCloud](https://www.irccloud.com/))

### Conferences & Meetups
There are many Go meetups and conferences nowadays. Here are some of the most popular conferences. *Due to the current epidemic, many were moved online for the time being or have announced they'll follow a dual online/presential approach.*

* [GopherCon](https://gophercon.com) ([videos](https://www.youtube.com/channel/UCx9QVEApa5BKLw9r8cnOFEA))
* [GopherCon EU](https://gophercon.eu/) ([videos](https://www.youtube.com/c/GopherConEurope))
* [GopherCon UK](https://www.gophercon.co.uk/) ([videos](https://www.youtube.com/channel/UC9ZNrGdT2aAdrNbX78lbNlQ))
* [GopherCon Brasil](https://gopherconbr.org/) ([videos](https://www.youtube.com/channel/UCGFVA_XvkUoMWpKVH0IrjUA/))
* [dotGo](https://www.dotgo.eu/) ([videos](https://www.youtube.com/user/dotconferences))

Kudos to initiatives such as [GoBridge](https://gobridge.org), making it possible for many underrepresented developers to attend the conferences – as the cost of tickets and travel can get as much as a few thousand dollars!


Look up Awesome Go's [conferences](https://github.com/avelino/awesome-go#conferences) and [meetups](https://github.com/avelino/awesome-go#conferences) list or the Gophers Slack to see if a local meetup group exists close to your location.

### News
* [Golang Weekly](https://golangweekly.com/), a weekly newsletter about the Go programming language.

### Podcasts
* [Go Time](https://changelog.com/gotime), a weekly podcast with diverse discussions from around the Go community
* [On The Metal](https://oxide.computer/podcast/)

### YouTube and Twitch channels
Besides the [aforementioned](#conferences--meetups) conference channels, plenty of Gophers have Go channels.

* [justforfunc: Programming in Go](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw) by [Francesc Campoy](https://www.campoy.cat/)
* [Filippo Valsorda's Twitch](https://twitch.tv/filosottile)
* [Ardan Labs channel](https://www.youtube.com/channel/UCCgGRKeRM1b0LTDqqb4NqjA)
* [Rob Evans a.k.a. @deadprogram Twitch](https://www.twitch.tv/lapipatv)

## Text editors and IDEs
Thanks to the [Language Server Protocol](https://microsoft.github.io/language-server-protocol/), a protocol originally designed for Microsoft Visual Studio Code, which allows programming language support to be implemented and distributed independently of any given editor or IDE, Go's support on development environments is widespread.

### Visual Studio Code
[Visual Studio Code](https://code.visualstudio.com/) is an open-source source code editor by Microsoft with a minimal GUI that doesn't get in your own.
I use it daily to work with Go code, and in my opinion, it's the best option for most people.
The first time you open a Go file with it, it should recognize the file and suggest you install the official [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go).
[Read more about Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go) to learn about shortcuts, [debug Go code](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) with [Delve](https://code.visualstudio.com/docs/languages/go), and more.

<a href="https://raw.githubusercontent.com/golang/vscode-go/"><img src="/img/posts/go/toggletestfile.gif" width="700" alt="Go extension"></a>

### Acme
[Acme](http://acme.cat-v.org/) is a text editor and graphical shell created by [Rob Pike](http://www.herpolhode.com/rob/) for the Plan&nbsp;9 operating system (more on that below). Famously used by C's designer [Dennis Ritchie](https://www.bell-labs.com/usr/dmr/www/), it's Vim's antithesis: it heavily relies on [mouse chording](https://en.wikipedia.org/wiki/Mouse_chording).

<div class="grid-x">
        <div class="medium-6 small-12">
                {{< youtube dP1xVpMPn8M >}}
        </div>
</div>

[A Tour of the Acme Editor – Russ Cox](https://www.youtube.com/watch?v=dP1xVpMPn8M) ([details](https://research.swtch.com/acme))

### Vim
[Vim](https://www.vim.org/) is a [text-based user interface](https://en.wikipedia.org/wiki/Text-based_user_interface) editor.
It requires time and effort to learn and memorize its many commands.
I use it primarily for punctual changes on configuration files, and when accessing remote machines.

For using it with Go, you'll want to install the [vim-go](https://github.com/fatih/vim-go) plugin by [Fatih Arslan](https://github.com/fatih).

<a href="https://github.com/fatih/vim-go"><img src="/img/posts/go/vim-go.png" width="320" alt="vim-go plugin"></a>

### GoLand
[GoLand](https://www.jetbrains.com/go/) is an [Integrated Development Environment](https://en.wikipedia.org/wiki/Integrated_development_environment) (IDE) by JetBrains (the authors of IntelliJ).
It's a full-featured IDE, and the obvious choice if you must do some heavy refactoring.
I use it occasionally – whenever I find myself doing a repetitive refactoring task.

<a href="https://www.jetbrains.com/go/"><img src="/img/posts/go/GoLand_6_Extensibility.png" width="700" alt="GoLand"></a>

## Coding style and code quality
Go has a very consistent style. Thanks to its simplicity, it is fairly easy to write tools to do static analysis of Go code, especially since the [x/tools/go/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) package appeared.

I recommend taking a light approach to enforcing code style.
Thanks to `go fmt`, the fight between [Tabs versus Spaces](https://www.youtube.com/watch?v=SsoOG6ZeyUI) is over.
You don't have an option, and have to use TABs.
And this is a good thing, even if you prefer spaces: you make a small concession of style for the great benefit of consistency.
You'll also notice most Go static analyzers also don't let users configure their preferences, and in my opinion this is a good thing.
It's hard to make everyone happy, so why not try to achieve a consistent style instead?
For me this is an easy trade-off.

<div class="grid-x">
        <div class="medium-6 small-12">
                {{< youtube SsoOG6ZeyUI >}}
        </div>
</div>
&nbsp;

### Testing
Go [testing library](https://pkg.go.dev/testing) is pretty straight-forward.
Go has two types of tests: `*testing.T` (regular tests) and `*testing.B` (benchmark tests).

For people used to other programming languages, the lack of assertion functions might seem a bit odd at first. Still, it's pretty good as it makes it crystal clear what you're testing, and your testing code looks more like production code.

* Heavy assertion libraries slows down testing due to increased build time.
* BDD testing libraries make it especially hard to debug code using common tooling.

```go
func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		// You almost always want to use t.Error* functions instead of t.Fatal*.
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
```

To run the tests, I use an alias I called `gotest`:

```shell
$ go test -race -coverprofile=coverage.out && go tool cover -html coverage.out -o coverage.html
```

* `-race` flag enables the [Data Race Detector](https://golang.org/doc/articles/race_detector)
* `-coverprofile` flag enables code coverage for my code

Most of the time, I've been running this all from within Visual Studio Code, in any case.
The `-run` flag enables you to run functions matching a specific regex – really useful when you need to run a specific test or group of related tests.
Use `go test helpflag` to learn more about the available flags.

```shell
$ go test -run=TestValidUser
```

{{< tweet user=henriquev id=1370366670787780614 >}}

#### httptest
The [http](https://pkg.go.dev/net/http) package is one of the best packages in Go.
You don't need a web framework to write web services effectively.
To write integration tests for your endpoints, the [http/httptest](https://pkg.go.dev/net/http/httptest) package might be useful.

#### Fuzzy testing
[Fuzz testing](https://en.wikipedia.org/wiki/Fuzzing) is a technique that involves providing invalid, unexpected, or random data as inputs to a computer program.
There is a [proposal](https://github.com/golang/go/issues/44551) to add fuzz test support to the standard library.
Meanwhile, you can use [gofuzz](https://github.com/google/gofuzz) and [go-fuzz](https://github.com/dvyukov/go-fuzz).

### Static analysis tools
Before pushing changes upstream, I like to run a script that executes some programs to check code quality.
I also find it very productive to have a [Continuous Integration](https://en.wikipedia.org/wiki/Continuous_integration) system such as [GitHub Actions](https://github.com/features/actions) set up to run most of them once code is submitted upstream.

#### go vet
> Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.
>
> Source: [documentation](https://golang.org/cmd/vet/)

The `go vet` is part of the Go toolchain, and have a list of analyzers you can see with `go tool vet help`.
By default, all analyzers are run.

```shell
$ # ./... to walk the working directory recursively
$ go vet ./...
```

As with typical Unix-like commands, no output or exit error = 0 means you're good.

#### Staticcheck
[Staticcheck](https://staticcheck.io/) is a state of the art linter for Go by [Dominik Honnef](https://dominik.honnef.co/).
You don't really need any other linter besides `go vet` and `staticcheck`.

```shell
$ staticcheck ./...
```

#### GitHub Security code scanning
Suppose you’re using GitHub to host an open source project's repository or willing to pay extra for their Enterprise plan. In that case, [GitHub Security](https://github.com/features/security) has you covered with its [CodeQL](https://securitylab.github.com/tools/codeql) security scanner.
[Sonar](https://rules.sonarsource.com/go/RSPEC-138) is another option.

#### gosec
[gosec](https://github.com/securego/gosec) is a security checker for Go.
It scans your code and pinpoints unsafe security usage.
By default, I find it quite noisy as it reports unhandled errors ignoring what I see as many false-positive cases, so I exclude this specific rule when I decide to use this tool on one of my projects.

```shell
$ # Ignoring gosec unhandled errors warning due to many false-positives.
$ gosec -quiet -exclude G104 ./...
```

#### unparam
[unparam](https://github.com/mvdan/unparam) helps you find unused function parameters and results in your code. **I must warn that I use this tool manually**, not in a script or CI system, as I find it distracting there, especially at the beginning of a project.

#### mispell
[mispell](https://github.com/client9/misspell) is a great tool for checking common English mistakes on your source code.
I also often use [Grammarly](https://www.grammarly.com/) to review my writing on documentation.

```shell
$ misspell cmd/**/*.{go,sh} internal/**/*.{go} README.md
```

I also use other static analyzer tools from time-to-time, manually, and that I'd advise against adding to a Continuous Integration pipeline.

#### gocyclo
[gocyclo](https://github.com/fzipp/gocyclo) calculates [cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) of functions in your Go source code.

I've used this in the past to find complicated code and to decide where to start refactoring code. I also prefer to use this manually, but I see some value in adding it as a step of a CI pipeline – as long as you err by being generous and configure the tool to scream only in case of undeniable high values.

```shell
$ gocyclo -over 15 -ignore "generated.go|pb.go|_test|vendor/" .
```

## Standard library
Go's standard library is outstanding both in terms of covering a wide range of needs and in terms of quality.
You'll find that with Go, you can rely on the standard library for many things where you'd typically require adding external dependencies to your project.

I always suggest to people starting with Go to take some time to read both the [docs](https://pkg.go.dev/net/http), and the source code for [net/http](https://github.com/golang/go/tree/master/src/net/http), as it is the best example of a well-organized and successful package I can think of.

### x packages
Packages inside the /x/ have looser compatibility requirements than the rest of the standard library.
A package that shows to be helpful throughout the ecosystem might be promoted from /x/ to the status of a regular package of the standard library, like [context](https://pkg.go.dev/context).

> Code in sub-repositories of the main go tree, such as [golang.org/x/net](https://pkg.go.dev/golang.org/x/net), may be developed under looser compatibility requirements. However, the sub-repositories will be tagged as appropriate to identify versions that are compatible with the Go 1 point releases.
>
> [Go 1 and the Future of Go Programs: Sub-repositories](https://golang.org/doc/go1compat):

## Third-party packages
You can go really far with the standard library, and it's not uncommon to write production-quality programs relying solely on it.
Before you jump and start importing a plethora of external packages on your code, do yourself a favor and remember the [Go proverb](https://go-proverbs.github.io/):

> [A little copying is better than a little dependency.](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=568s)

It's important to be aware of the cost of importing dependencies to avoid pitfalls like the [leftpad fiasco](https://www.theregister.com/2016/03/23/npm_left_pad_chaos/), security vulnerabilities, quality problems, or [worse](http://sunnyday.mit.edu/papers/therac.pdf).

> I’ve also noticed that in Go I need fewer dependencies, and my dependencies themselves have fewer dependencies. Go doesn’t have a culture of exporting as-much-logic-as-possible to external dependencies. Code duplication is more acceptable in the Go community than elsewhere. This can be frustrating. Sometimes, you just want a good library that performs some type of sanitation or parsing. Many times, you’ll need to write that functionality yourself, or copy/paste it from a StackOverflow answer. Generally I think this is a positive. Fewer dependencies means fewer things that break while you let a project sit idle for a couple months.
> 
> [The Value in Go's Simplicity](https://benjamincongdon.me/blog/2019/11/11/The-Value-in-Gos-Simplicity/)

You can search for Go packages in [pkg.go.dev](https://pkg.go.dev/), and discover commonly used ones in the  [go.dev](https://go.dev/) pages.
If you want an extensive and comprehensive list of Go packages with a short one-line description, [Awesome Go](https://awesome-go.com/) is what you want!

> Did you know Go programs can only build so fast because it doesn't allow circular imports?
> This way, the linker can be simpler, and we mitigate one of the common problems of [dependency hell](https://en.wikipedia.org/wiki/Dependency_hell).

### cobra
[cobra](https://cobra.dev/) is a framework for Modern CLI Apps by [Steve Francia](https://spf13.com/).

### pgx (PostgreSQL driver)
[pgx](https://github.com/jackc/pgx) is the best PostgreSQL driver and toolkit for Go. While normally you want to use database/sql interfaces for connecting to a database, you probably want to use pgx without it due to a number of advantages thanks to PostgreSQL's binary protocol.

### gRPC-Go
[gRPC-Go](https://github.com/grpc/grpc-go) is the Go implementation of the *high performance, open source universal RPC framework* [gRPC](https://grpc.io/), which has a strict specification compared to the more popular HTTP APIs with JSON.

### paseto (alternative to JWT)
[paseto](https://github.com/o1egl/paseto) is a Platform-Agnostic Security Tokens implementation of [PASETO](https://paseto.io/) tokens. This standard is <em>everything you love about JOSE (JWT, JWE, JWS) without any of the [many design deficits that plague the JOSE standards](https://paragonie.com/blog/2017/03/jwt-json-web-tokens-is-bad-standard-that-everyone-should-avoid)</em>.

### GoReleaser
[GoReleaser](https://goreleaser.com/) is a powerful packaging tool. With GoReleaser, you can:

* Cross-compile your Go project
* Release to GitHub, GitLab, and Gitea
* Create Docker images and manifests
* Create Linux packages and Homebrew taps
* ... and much more!

### equinox.io
[equinox.io](https://equinox.io/) is a tool and service that you can use to package and distribute your Go programs ([library](https://github.com/equinox-io/equinox)). It has a slightly more limited packaging scope than GoReleaser, but provides some killer features such as different distribution channels (i.e., you can create channels stable, unstable, etc.) and code signing.

To learn more about distributing Go binaries for multiple platforms safely, read my post [Counter-Strike code leaked: should you worry?](https://henvic.dev/posts/cs-security/)

### goexpect
[goexpect](https://github.com/google/goexpect) is an implementation of [Expect](https://core.tcl-lang.org/expect/) in Go (I created [pseudoterm](https://github.com/henvic/pseudoterm), something quite similar).

### subcommands
[subcommands](https://github.com/google/subcommands) is a Go package that implements a simple way for a single command to have many subcommands.
I created [clino](https://github.com/henvic/clino) about a year ago, unaware of this.

### avo
[avo](https://github.com/mmcloughlin/avo) makes high-performance Go assembly easier to write, review and maintain.

### httpretty
[httpretty](https://github.com/henvic/httpretty) is a package I created to help me debug HTTP requests on CLI applications.

<div class="asciicast-carrier">
    <script id="asciicast-297429" src="https://asciinema.org/a/297429.js" data-autoplay="true" data-loop="true"
        async></script>
</div>

### Lip Gloss
[lipgloss](https://github.com/charmbracelet/lipgloss) is [Charm](https://charm.sh/)'s package for writing Text-based user interfaces (TUIs) (just like [ncurses](https://invisible-island.net/ncurses/), but better).

### Bubble Tea
[Bubble Tea](https://github.com/charmbracelet/bubbletea) is a framework to build terminal apps.

### Bubbles
[Bubbles](https://github.com/charmbracelet/bubbles) is a package containing TUI components for Bubble Tea.

### SQL Databases
* Writing SQL yourself is always the best option to write sane and fast queries. Use [sqlc](https://sqlc.dev/) if you want to reduce the amount of code you need to write yourself, instead of [using an ORM and making your application slower and your program more error-prone and harder to debug](https://blog.codinghorror.com/object-relational-mapping-is-the-vietnam-of-computer-science/).

## Hardware integration
[The Hybrid Group](https://hybridgroup.com/) has some quite exciting projects.
If you're interested in interfacing with hardware, running Go in embedded computers, or computer vision, you want to check what they're up to!

<a href="https://gobot.io/"><img src="/img/posts/go/gobot-logo-small.png" width="160" alt="Gobot"></a>
<a href="https://gocv.io/"><img src="/img/posts/go/gocvlogo.jpg" width="160" alt="GoCV"></a>
<a href="https://tinygo.org/"><img src="/img/posts/go/tinygo-logo-small.png" width="160" alt="TinyGo"></a>

<div class="grid-x">
        <div class="medium-6 small-12">
                {{< youtube EiB9ZVrvrz0 >}}
        </div>
</div>

[Small is Going Big: Go on Microcontrollers](https://www.youtube.com/watch?v=EiB9ZVrvrz0) by [Ron Evans](https://twitter.com/deadprogram) at GopherCon 2019.

* [Gobot](https://gobot.io/) is a framework for robotics, physical computing, and the Internet of Things (IoT) with support for dozens of different hardware/software platforms.
* [GoCV](https://gocv.io/) is a package for computer vision using [OpenCV 4](https://opencv.org/opencv-4-0/) and beyond.
* [TinyGo](https://tinygo.org/) is a Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on [LLVM](https://llvm.org/).

## Go modules and vendoring
> A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.
>
> [Go Modules Reference](https://golang.org/ref/mod)

Go modules appeared in 2019 after the language and its ecosystem were mature and the community already experimented with a diverse number of approaches to solve the problem of [managing dependencies](https://golang.org/doc/modules/managing-dependencies).

**See also:**
* [Managing dependencies](https://golang.org/doc/modules/managing-dependencies)
* [Go Modules Reference](https://golang.org/ref/mod)
* [go.mod file reference](https://golang.org/doc/modules/gomod-ref)
* [Using Go Modules](https://blog.golang.org/using-go-modules)

### Auditing
Once you've go mod set up, you might want to audit your dependencies. There are two ways:

* You could set up your own Go modules proxy and whitelist what dependencies to allow.
* You can vendor your dependencies.

I don't have any suggestions regarding auditing with a go proxy, but I audit dependencies when vendoring (or vendorizing) code I control – more on that in a minute!
Auditing is an important step to verify the quality of your dependencies and to check if no malicious code was injected somewhere.

> Keeping the number of dependencies low also means you've less to audit!

### Vendoring
Vendoring (or vendorizing) is copying dependencies to your own project, versioning them along with your codebase.
Not everyone is comfortable with the idea of copying external code to their projects, but I consider this the safest approach to guarantee long-term access to your dependencies.
You avoid the risk of losing access to your dependencies for a variety of reasons, including the evil <abbr title="Digital Millenium Copyright Act">DMCA</abbr> takedown notices.

* Go dependencies are mostly small textual .go files, so they don't take up much space.
* It mitigates the risk of losing access to the exact version of the dependency you use in case a dependency is taken down.
* Faster binary search for bad commits when using `git bisect` to debug code.

### Go modules proxy
Alternatively, suppose you want control over what dependencies your team uses. In that case, you might consider setting up your own [Go proxy](https://proxy.golang.org/).
This might work well for enterprises with a strict requirement for tight control.
Maybe solutions like [JFrog Artifactory](https://jfrog.com/integration/go-registry/) might help.

## Tools for software development
### godoc
[godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) is a program that extracts and generates documentation for Go programs.
You can use it to browse the documentation for your code and the standard library off-line easily, using a web browser.

It is also what originally powered godoc.org, before it was sunset by [pkg.go.dev](https://pkg.go.dev), which is powered by the more complex [pkgsite](https://github.com/golang/pkgsite).

```shell
$ go install golang.org/x/tools/cmd/godoc@latest
$ godoc
$ open http://localhost:6060
```

Use `godoc` flags `-play` to enable playground and `-index` to enable the search index.

* [Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html)

### swag
[swag](https://github.com/swaggo/swag) is a tool to automatically generate RESTful API documentation with Swagger 2.0 (it still doesn't support the newest [OpenAPI Specification](https://swagger.io/specification/), though).

### Goda
[Goda](https://github.com/loov/goda) is a Go dependency analysis toolkit. It contains tools to figure out what your program is using. For example, you can print out dependency trees or draw a dependency graph, and see the impact of cutting a package from your program.

### depaware
[depaware](https://github.com/tailscale/depaware) is a Go dependencies analysis tool by [Tailscale](https://tailscale.com/).

### gops
[gops](https://github.com/google/gops) is a Google tool to list and diagnose Go processes currently running on your system.

### goversion
[goversion](https://rsc.io/goversion) scans a directory tree and, for every executable it finds, prints the Go version used to build that executable.
It also has a flag to show module info.

### sqlc
[sqlc](https://sqlc.dev/) generates fully-type safe idiomatic Go code for use with database/sql from SQL.
Take a look at their [Playground](https://play.sqlc.dev/) to see some examples.

### hey
[hey](https://github.com/rakyll/hey) is a tiny program by [Jaana Dogan (@rakyll)](https://rakyll.org/) that sends some load to a web application (similar to Apache's [ab](https://en.wikipedia.org/wiki/ApacheBench)).

### vegeta
[vegeta](https://github.com/tsenart/vegeta) is another HTTP load testing tool.

### Go Guru
[guru](https://pkg.go.dev/github.com/golang/tools/cmd/guru) is a tool for answering questions such as the ones below about your source code. Read the [documentation](http://golang.org/s/using-guru) for more information.

> * Where is this identifier declared?
> * Where are all the references to this declaration?
> * What are the fields and methods of this this expression?
> * What is the API of this package?
> * Which concrete types implement this interface?
> * What are the possible callees of this dynamic call?
> * What are the possible callers of this function?
> * Where might a value sent on this channel be received?

### Pythia
[Pythia](https://github.com/fzipp/pythia) is a browser-based user interface for [guru](#go-guru).

<a href="https://github.com/fzipp/pythia"><img src="/img/posts/go/pythia_screenshot.png" width="800" alt="Pythia screenshot"></a>

### Wails
[Wails](https://wails.app/) is a framework for building applications using Go and Web Technologies.
I haven't had the opportunity to try it out yet, but I'm excited about the possibilities!

### mkcert
[mkcert](https://github.com/FiloSottile/mkcert) is a simple tool for making locally-trusted development certificates.

## Productivity tools
### present
[present](https://pkg.go.dev/golang.org/x/tools/present) is a tool that you can use to create articles or slides.
It's useful for writing simple and direct presentations when you don't care about setting a specific visual style, as it uses a simple black & white theme that you cannot configure.

* [Go talks](https://github.com/golang/talks) created with present
* [Tutorial](https://medium.com/@drashti.ved_84172/level-up-your-go-presentations-b4d06fc495e5)

```shell
$ go install golang.org/x/tools/cmd/present@latest
```

### Hugo
[Hugo](https://gohugo.io/) is the world's fastest framework for building websites by [Bjørn Erik Pedersen](https://bep.is/) and [Steve Francia](https://spf13.com/), and it's able to compile whole websites in a fraction of a second!
It's a natural choice if you want to write your own website or the documentation for a project – using Go templates!

[I've been using it](https://github.com/henvic/henvic.github.io) for almost two years for generating this website.
You can create your own theme or use one of the many available online.
I decided to create my own, so I used a minimal one as the base in the very beginning to understand how Hugo works.

<a href="https://gohugo.io/"><img src="/img/posts/go/hugo-logo-wide.svg" width="400" alt="Hugo"></a>

[Hugo in Action](https://amzn.to/3c3je3W) is an upcoming book about Hugo (Summer 2021). You can learn more about it on the [publisher's page](https://www.manning.com/books/hugo-in-action).

<div id="amzn-assoc-ad-28708ac8-a880-4aff-b5fd-2649b98d4954"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=28708ac8-a880-4aff-b5fd-2649b98d4954"></script>

### Caddy
[Caddy](https://github.com/caddyserver/caddy) is an extensible HTTP server platform written in Go.

I've been using it for local software development instead of [nginx](https://www.nginx.com/) to serve static files and to expose multiple HTTP services on the same address using a reverse proxy.

### gophernotes
[gophernotes](https://github.com/gopherdata/gophernotes) is a Go kernel for [Jupyter](https://jupyter.org/) notebooks and [nteract](https://nteract.io/).

> The Jupyter Notebook is an open-source web application that allows you to create and share documents that contain live code, equations, visualizations and narrative text. Uses include: data cleaning and transformation, numerical simulation, statistical modeling, data visualization, machine learning, and much more.

### xbar
[xbar](https://xbarapp.com/) is a macOS program written in Go that lets you put the output of any script/program in your macOS menu bar.
It's like [iStat Menus](https://bjango.com/mac/istatmenus/), but for your menus! It’s an open-source project written primarily in Go by [Mat Ryer](https://twitter.com/matryer)!

### ghz
[ghz](https://ghz.sh) is a gRPC benchmarking and load testing tool.

### Evans
[Evans](https://github.com/ktr0731/evans) is gRPC client with REPL and CLI modes.

## Reporting bugs and proposing feature requests
* The project uses publicly accessible [GitHub issues](https://github.com/golang/go/issues) to track and discuss those.
* Before you submit a proposal, search to see if what you've in mind was already discussed in the past. Use your findings wisely.
* If you want to report a security bug, please follow the [Go Security Policy](https://golang.org/security).

## Code review process
* See [go-review](https://go-review.googlesource.com/). The Go project and several other Google products like Chromium (see [chromium-review](https://chromium-review.googlesource.com/)) use the [Gerrit Code Review](https://www.gerritcodereview.com/) to manage collaboration: it's like GitHub's Pull-Request, but on steroids!

### Contributing back
* You probably want to start by drafting a proposal.
* You should read the [Contribution Guide](https://golang.org/doc/contribute).
* Sign and submit a simple and straight-forward Contributor License Agreement (CLA).
* Please make sure you check at least a few proposals and reasons why they were either approved or rejected before submitting your own proposal.
* In my blog post [signal.NotifyContext: handling cancelation with Unix signals using context](/posts/signal-notify-context/), I share a bit of my experience adding a feature to the signal package.

## Gopher
<a href="https://blog.golang.org/gopher"><img src="/img/posts/go/gopher.jpg" width="600" alt="The Go gopher"></a>

The gopher is an iconic mascot of the language. It was created by [Renée French](https://www.instagram.com/reneefrench/), who also created the cute Glenda, the Plan&nbsp;9 Bunny.
You can create your own Gopher at [Gopherize.me](https://gopherize.me/). But what does Plan&nbsp;9 has to do with Go? Well, a lot!

## Plan&nbsp;9 influence

[Plan&nbsp;9](https://en.wikipedia.org/wiki/Plan_9_from_Bell_Labs) is a distributed operating system created at Bell Labs by members of the same group that originally developed Unix and the C programming language.
Part of the Plan&nbsp;9 team included [Rob Pike](http://www.herpolhode.com/rob/), Ken Thompson, and Russ Cox – an avid contributor to the [Plan 9 from User Space](https://9fans.github.io/plan9port/) fork – and they brought many ideas and concepts of the operating system to Go.

<a href="https://9p.io/plan9/glenda.html"><img src="/img/posts/go/plan9bunnysmblack.jpg" width="190" alt="Glenda, the Plan 9 Bunny"></a>

Plan&nbsp;9 is an interesting topic on its own, but my experience with it is fairly limited. Besides the aforementioned [Acme](#acme), two interesting things about Plan&nbsp;9:

* Files are key objects in Plan&nbsp;9 with their [9P](https://en.wikipedia.org/wiki/9P_(protocol)) protocol (even more than in UNIX), and even a mouse is a file.
* [UTF-8](https://en.wikipedia.org/wiki/UTF-8), a variable-width character encoding that is de facto the core of [Unicode](https://home.unicode.org/), was also created by Ken Thompson and Rob Pike, and the original implementation was for Plan&nbsp;9. Not surprisingly, Go has great UTF-8 support!

```go
// Quadratic takes a general quadratic equation of the form
// ax² + bx + c = 0.
func Quadratic(a, b, c float64) (x1, x2, Δ float64) {
	Δ = math.Pow(b, 2) - (4 * a * c)
	x1 = (-b + math.Sqrt(Δ)) / (2 * a)
	x2 = (-b - math.Sqrt(Δ)) / (2 * a)
	return x1, x2, Δ // Just don't go this wild!!!
}
```

Thanks! If you liked what I wrote, maybe you want to check out my other posts, such as my experience creating a [homelab](/posts/homelab/) to play with different operating systems or learn about [my Go mistakes](/posts/my-go-mistakes/).

{{< tweet user=henriquev id=1373980655906852867 >}}

<div id="amzn-assoc-ad-e4e6eccf-8b48-4046-a4d9-37f587a481a3"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=e4e6eccf-8b48-4046-a4d9-37f587a481a3"></script>

<small>If you click and buy any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/).</small>
