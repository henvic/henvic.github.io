---
title: "signal.NotifyContext: handling cancelation with Unix signals using context"
type: post
description: "Today my first contribution to the Go standard library was merged. It simplifies handling operating system signals in Go for certain common cases."
date: "2020-09-16"
image: "/img/posts/signal-notify-context/gopher-frontpage.png"
hashtags: "go,golang,os,signal"
---
From Go 1.16 onwards, you'll be able to use

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()
```

to control context cancelation using context, simplifying handling operating system signals in Go for certain common cases. This is my first contribution to the Go standard library, and I am very excited!

## Why
When writing <abbr title="command-line interface">CLI</abbr> code, I often needed to handle cancellation – for instance, when a user presses CTRL+C producing an interrupt signal. Another possible use case is to handle graceful termination of HTTP servers using `http.*Server.Shutdown(ctx)`.

I had to write code using the signal package often and in multiple places. I didn't want that, so I wrote the [ctxsignal](https://github.com/henvic/ctxsignal) package to solve this common problem two years ago.

However, I kept finding it in other places and noticing people would often not handle proper termination correctly or at all, so I got motivated to submit a proposal and try to improve this situation.

## How to use
```go
// This example passes a context with a signal to tell a blocking function that
// it should abandon its work after an operating system signal is notified.
func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		stop()
		fmt.Println("signal received")
	}
}
```

This program will print `"missed signal"` after 10 seconds or will print `"signal received"` when the user sends an interruption signal, such as by pressing CTRL+C on a keyboard.

Please note that the second returned value of the `signal.NotifyContext` function is a function called `stop` instead of `cancel`. Once you're done handling a system signal, you should call `stop`. We call it `stop` instead of `cancel` because you need to call it to stop capturing any further system signal you registered with `NotifyContext`.

> The stop function unregisters the signal behavior, which, like signal.Reset, may restore the default behavior for a given signal.

## Timeline
* I submitted a [proposal](https://github.com/golang/go/issues/37255) and [implementation](https://golang.org/cl/219640) on 17 February (as WithContext in the signal package)
* After some discussions, a somewhat modified proposal was accepted on 1 April as a WithCancelSignal in the context package
* It was moved back to the proposal stage on 15 April after some concerns were presented about having it in the context package
* After more discussions, it was accepted again on 20 May as NotifyContext in the signal package
* A couple of days ago, I restarted working on it.
* Today it got merged and is available in the tip
* [Go 1.16](https://tip.golang.org/doc/go1.16) is expected to be released in February

## Other points
* I delayed working on it and missed the Go 1.15 release due to its [code freeze window](https://github.com/golang/go/wiki/Go-Release-Cycle)
* Go uses [Gerrit](https://www.gerritcodereview.com/) to track changes and code reviews.
* I found Gerrit's patchsets better over GitHub's pull-requests to keep track of changes without losing context (pun intended).

Thanks a lot to everyone who provided ideas, feedback, or code-reviewed my code. In particular, the Go team for the patience in helping me out.