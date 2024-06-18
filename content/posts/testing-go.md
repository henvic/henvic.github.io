---
title: "On testing Go code using the standard library"
type: post
description: "Most programming language ecosystems provide assert functions in their testing libraries but not Go's. Go's standard testing package follows a more direct and to-the-point approach."
date: "2024-06-18"
image: "/img/posts/go/go-logo-blue.png"
hashtags: "golang,testing"
---

Most modern programming language ecosystems provide assert functions in their testing libraries but not Go's.
Its [standard testing package](https://pkg.go.dev/testing) follows a more direct and to-the-point approach.
In fact, there isn't even a single assertion function in the testing package, and writing idiomatic tests in Go isn't that different from writing application code.

You mainly use the `t.Errorf` and `t.Fatalf` functions, which borrows the idioms of the [fmt package](https://pkg.go.dev/fmt) to format output, as shown in this code, meaning you get to use the helpful printing verbs of the fmt package, such as:

```txt
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%v	the value in a default format when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
```

For example:

<div class="grid-x">
<div class="cell medium-12 large-6">

```go
package abs

import "testing"

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
```

</div>
<div class="cell auto">

```shell
$ go test # on success
PASS

# On failure
$ go test
--- FAIL: TestAbs (0.00s)
    code_test.go:14: Abs(-1) = 3; want 1
FAIL
exit status 1
FAIL	github.com/henvic/exp	0.114s
```

</div>
</div>

To me, this provides a vastly superior experience than writing a test using an assertion library that follows the [XUnit](https://en.wikipedia.org/wiki/XUnit)-style:

<div class="grid-x">
<div class="cell medium-12 large-6">

```go
package abs

import (
	"testing"

	// testify is very used in the Go ecosystem.
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(-1))
}
```

</div>
<div class="cell auto">

```shell
$ go test
--- FAIL: TestAbs (0.00s)
    code_test.go:14:
        	Error Trace:	/exp/code_test.go:14
        	Error:      	Not equal:
        	            	expected: 1
        	            	actual  : 3
        	Test:       	TestAbs
FAIL
exit status 1
FAIL	github.com/henvic/exp	0.152s
​
```

</div>
</div>

Using an assertion library is often seen as a way to reduce the effort in writing testing code.
Sure, we can save three lines of code using [assert.Equal](https://pkg.go.dev/github.com/stretchr/testify/assert#Equal) instead of t.Error, but is this really a good idea?
For me, this is a distraction.

> There are many things in the Go language and libraries that differ from modern practices, simply because we feel it's sometimes worth trying a different approach.
>
> Source: [Go FAQ: Why does Go not have assertions?](https://go.dev/doc/faq#assertions)

## Optimizing for reading vs. optimizing for writing

Rather than focusing on small gains in speed when writing the initial code, you want to make the intention of your code clear in the long run.
Writing idiomatic Go code reduces the amount of time you spend chasing defects and maintaining your software.
On [Go Testing By Example](https://research.swtch.com/testing), Russ Cox provides excellent tips on how to do that.

<div class="grid-x">
<div class="cell large-4 medium-6">
{{< youtube X4rxi9jStLo >}}
</div>
</div>

**Sidenote:** You can also pass additional arguments to Testify's `assert.Equal` (and others functions) to add a line with why a failure occurred, but this isn't common.


## On t.FailNow() and testify/require
Another widespread problem I see with assertion libraries for Go is that they indiscriminately call [t.FailNow](https://pkg.go.dev/testing#T.FailNow) whenever an error happens is widespread.

This is the effect of doing this:

```go
t.Error("this doesn't stops the execution")
t.Fatal("this kills a test")
t.Error("not printed")

// $ go test ./...
// --- FAIL: TestAbs (0.00s)
//     code_test.go:13: this doesn't stop the execution
//     code_test.go:14: this kills a test
// FAIL
// FAIL	github.com/henvic/exp	0.115s
// FAIL
```

Which, when using Testify, might look be hidden in something like this:

```go
assert.Equal(t, 1, 2, "this doesn't stop the execution")
require.Equal(t, true, false, "this kills a test") // any errors after this won't be printed
assert.Equal(t, "a", "b", "not printed")
```

The effect is that the test doesn't print the third error ("not printed") as the preceding `t.Fatal` invocation terminates the goroutine executing the test by calling `t.FailNow()`, which is basically a call to `t.Fail()` followed by a call to `runtime.Goexit()`.

As you can see, Testify has two packages: [assert](https://pkg.go.dev/github.com/stretchr/testify/assert) and [require](https://pkg.go.dev/github.com/stretchr/testify/require), where the difference is that the functions of the second call the functions of the first but also stop the test execution when a test fails by calling `t.FailNow()` once the first returns.
I can't fathom their design decison of separating this into two packages.

¯\\\_(ツ)\_/¯

```go
package require
// ...
func Equal(t Testing.TB, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
        t.Helper() // When printing file and line information, this function will be skipped.
        if !assert.Equal(t, expected, actual, msgAndArgs...) {
		t.FailNow()
	}
}
```

> FailNow marks the function as having failed and stops its execution by calling `runtime.Goexit` (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

Some developers have a strong risky preference for always using `t.Fatal` or Testify's require, going as far as enforcing this linter logic or during code review process, usually in the name of consistency.
Arguing against this, for many times I heard that the Go team had "fixed this already," and it was now safe to call `it.FailNow` anywhere. Well, not really.
Besides, they are missing the point.

By calling `t.Fatal` indiscriminately you more often than not hides useful error messages from a check that comes after your initial check that failed.
I've been hit by this particular issue on many codebases multiple times, such as when something such as a deferred function executes a `t.FailNow()`, masking test panics (see [Go issue #29207](https://github.com/golang/go/issues/29207)).

## Helper functions vs. assertion function
From time to time you might find yourself trying to check the same logic over and over.
For example, you might want to check if all positive numbers of a slice are even.

The following code is a reasonable option:

```go
func assertPositiveEvens(t testing.TB, numbers []int) {
	for pos, num := range numbers {
		if num > 0 && num%2 != 0 {
			t.Fatalf("number at index %d = %d is not a positive even", pos, num)
		}
	}
}
```

However, as you can see in this [other example](https://google.github.io/styleguide/go/decisions#assertion-libraries), it might be more interesting to have a function that returns *a value or an error that can be used in the test's failure message instead*.

```go
func checkPositiveEvens(numbers []int) error {
	for pos, num := range numbers {
		if num > 0 && num%2 != 0 {
			return fmt.Errorf("index(%d) = %d is not a positive even", pos, num)
		}
	}
        return nil
}

// Usage:
if err := checkPositiveEvens(vector); err != nil {
	t.Error("invalid vector:", err)
}
```

**Why?** This makes composing errors much easier:

```go
func TestVectors(t *testing.T) {
	if err := checkPositiveEvens([]int{2, 4, 6, 8, 10, 7}); err != nil {
		t.Error("invalid vector:", err)
	}
}

// exp/code_test.go:27: invalid vector: index(5) = 7 is not a positive even
```

P.S. Have you noticed that `assertPositiveEvens` masked a call to `t.FailNow` by using `t.Fatal` rather than `t.Error`?

## Comparing full structures
The package [github.com/google/go-cmp](https://github.com/google/go-cmp) is a package for the equality of Go values.
Using it, you have a powerful approach for comparing whether two values are semantically equal.

Suppose you've a list of animals on a database and need to verify if the animal you retrieved after calling it matches what you expect.
So, you have a struct with:

```go
type Animal struct {
	Name  string
	Class string
	Sound string
}
```

And a set of values like this:

```go
var gecko = Animal{
	Name:  "Gecko",
	Class: "Reptile",
	Sound: "gecko",
}

var dog = Animal{
	Name:  "Dog",
	Class: "Mammal",
	Sound: "Bark",
}
```

One naïve strategy might be to use:

```go
if dog != gecko {
	t.Error("dog and gecko are not the same animal")
}
```

This would work so far, but not for too long as:

> Struct types are comparable if all their field types are comparable. Two struct values are equal if their corresponding non-blank field values are equal. The fields are compared in source order, and comparison stops as soon as two field values differ (or all fields have been compared).
>
> Source: [Go spec: Comparison operators](https://go.dev/ref/spec#Comparison_operators)

For comparing struct with such fields you need at least [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual):

```go
if !reflect.DeepEqual(dog, gecko) {
	t.Error("dog and gecko are not the same animal")
}
```

Great! Now, it seems to be working as intended.
However, you still can't know exactly **why** the values are different.

Can you use the following testify/assert to rescue you?

```go
assert.Equal(t, want, getAnimal("dog"))
```

If you now change the type of the `Sound` field to a slice of strings, you get something like:

```shell
$ go test
--- FAIL: TestAnimals (0.00s)
    code_test.go:29:
        	Error Trace:	/Users/henvic/projects/gocode/src/github.com/henvic/exp/animalia_test.go:29
        	Error:      	Not equal:
        	            	expected: animalia.Animal{Name:"Dog", Class:"Mammal", Sound:[]string{"Bark"}}
        	            	actual  : animalia.Animal{Name:"Gecko", Class:"Reptile", Sound:[]string{"Click"}}

        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -1,6 +1,6 @@
        	            	 (animalia.Animal) {
        	            	- Name: (string) (len=3) "Dog",
        	            	- Class: (string) (len=6) "Mammal",
        	            	+ Name: (string) (len=5) "Gecko",
        	            	+ Class: (string) (len=7) "Reptile",
        	            	  Sound: ([]string) (len=1) {
        	            	-  (string) (len=4) "Bark"
        	            	+  (string) (len=5) "Click"
        	            	  }
        	Test:       	TestAnimals
FAIL
exit status 1
FAIL	github.com/henvic/exp	0.146s
```

It looks like it did the trick!

However, that won't work for too long either, as you might want to skip fields or check dynamic data, but let's talk about it later.

### Can we do better?

> If your function returns a struct, don’t write test code that performs an individual comparison for each field of the struct. Instead, construct the struct that you’re expecting your function to return, and compare in one shot using diffs or deep comparisons. The same rule applies to arrays and maps.
> Source: [Go Wiki: Go Test Comments](https://go.dev/wiki/TestComments)

Instead of using testify/assert, we can use go-cmp and have a much clearer and to-the-point error message:

```go
if !cmp.Equal(dog, gecko) {
	t.Errorf("animal is not a dog: %v", cmp.Diff(dog, gecko))
}

// or even:

if diff := cmp.Diff(dog, gecko); diff != "" {
	t.Errorf("animal is not a dog: %v", diff)
}
```

Which should print:

```shell
$ go test
--- FAIL: TestAnimals (0.00s)
    code_test.go:37: animal is not a dog:   animalia.Animal{
        - 	Name:  "Dog",
        + 	Name:  "Gecko",
        - 	Class: "Mammal",
        + 	Class: "Reptile",
        - 	Sound: []string{"Bark"},
        + 	Sound: []string{"Click"},
          }
FAIL
exit status 1
FAIL	github.com/henvic/exp	0.146s
```

I wonder if developers using Testify often use require rather than assert just because Testify needlessly prints too many lines.

Now, what if you want to skip fields or check dynamic data?
What do you do?
This is quite a common problem I often have to deal with.
One strategy is to extract such fields you want to verify explicitly, and then your testing code becomes "a wall" of assertion calls, like this:

```go
got := getAnimal("dog")

assert.Equal(t, "Dog", got.Name)
assert.Equal(t, "Mammal", got.Class)
assert.Equal(t, []string{"Bark"}, got.Sound)
assert.WithinDuration(t, time.Now(), got.CreatedAt, 10*time.Second)
assert.GreaterOrEqual(t, got.Age, 3)
```

Once you've enough fields, this can quickly get out of control, so you might find yourself starting to mix strategies of checking some individual fields, then preparing the bigger struct for an equality assertion.
So you end up with:

```go
assert.WithinDuration(t, time.Now(), got.CreatedAt, 10*time.Second)
assert.GreaterOrEqual(t, 3, got.Age)
// Cleaning dynamic values.
got.Age = 0
got.CreatedAt = time.Time{}
got.Location = nil
want.Location = nil // Location is always initialized for [some reason], so also needs to be reset.
// Test the rest.
assert.Equal(t, want, got)
```

With go-cmp, you could do the same using the [cmpopts package](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts) to help you a bit in reducing this amount of preparation before checking two structs:

```go
if !cmp.Equal(got, dog,
	cmpopts.EquateApproxTime(time.Second), // Check if recorded was just created.
        cmpopts.IgnoreFields(Animal{}, "Age", "Location")) {
	t.Errorf("animal is not a dog: %v", cmp.Diff(dog, got))
}

if got.Age < 3 {
        t.Errorf("animal age should be at least 3, got %d instead", got.Age)
}
```

### More complex comparisons
If you need to make more complex comparisons, look at the [go-cmp documentation](https://pkg.go.dev/github.com/google/go-cmp/cmp) to learn how you can check and transform the values.
Here is a simple example:

```go
func compareAgeDelta(delta int) cmp.Option {
	return cmp.FilterPath(func(p cmp.Path) bool {
		return p.GoString() == ".Age"
	}, cmp.Comparer(func(x, y int) bool {
		return Abs(x, y) <= delta
	}))
}

func TestAnimals(t *testing.T) {
	// ...
	if diff := cmp.Diff(want, got, compareAgeDelta(3)); diff != "" {
		t.Errorf("animals mismatch: %v", diff)
	}
}
```

You should also know that go-cmp will panic in some cases, such as:

* if you try to write a non-deterministic comparer (why we used the Abs function inside compareAgeDelta)
* if you try to compare structs with unexported fields without ignoring them
* if you use an invalid transformer or sorting function
* if it detects incomparable values or anonymous structs
* if it detects an unexported field and you didn't explicitly ignore it (might want to see [cmpopts.IgnoreUnexported](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts#IgnoreUnexported)).

While I don't use Testify on my projects, I understand its charms to newcomers to Go who are already used to xUnit-style tests from other ecosystems.
This testing approach is just one of the few Go design decisions that deviate from modern practice, and I'm fine with that.

## Can we do worse?
Yes.
**Much worse.**
Ginkgo's approach to Behavior-driven Development (BDD) hinders productivity beyond acceptable for me, both from an objective point-of-view, considering [mechanical sympathy](https://wa.aws.amazon.com/wellarchitected/2020-07-02T19-33-23/wat.concept.mechanical-sympathy.en.html), and from a developer experience expectation.
The hardest to maintain and slowest tests I have witnessed and had to tolerate in my career used Ginkgo and Gomega to test a web platform built using GORM (slow and buggy [ORM]((https://blog.codinghorror.com/object-relational-mapping-is-the-vietnam-of-computer-science/))).
By some back-of-the-envelope calculation, I can estimate they were at least 100 times slower than they could be (over 10min for something that should never take longer than half a minute in any circumstances) and at least a ten-fold order of magnitude harder to maintain due to their design choices getting in the way of Go tooling and completely ignoring the language idioms.

```go
package torture_test

import (
	// ...

	// The fun starts by not using qualified identifiers on imports
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var client = music.NewClient()

func TestRequests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Something speccial")
}

var _ = BeforeSuite(func() {
	client = service.NewClient()
})

var _ = Describe("Service test", func() {
	Context("Receiving a request", func() {
		It("Returning a response", func() {
			resp, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			body, err := io.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(body)).To(Equal(wantResp))
		})
	})
})
```

### Do you want another wild BDD tests against the Go idioms that people use?
¯\\\_(ツ)\_/¯

```go
type eateryStage struct {
	t *testing.T
        // ...
}

func (s *eateryStage) and() { return s }

func (s *eateryStage) IAmHungry() { require.Equal(s.you.food, 0) }
func (s *eateryStage) thereIsFood() { s.fridge.food = 10 }
func (s *eateryStage) IEat() {
        s.you.food++
        s.fridge.food--
        require.GreaterThan(t, s.fridge.food, 0)
}
func (s *eateryStage) IGetUnhungry() {
        require.GreaterThan(t, s.you.food, 0)
}

func newTestEateryStage(t *testing.T) (*eateryStage, *eateryStage, *eateryStage) {
	s := &eateryStage{
		// blah blah blah...
	}
	return s, s, s // 🤢 lol
}

func TestEat(t *testing.T) {
	given, when, then := newTestEateryStage(t)

	given.IAmHungry().and().thereIsFood()
	when.IEat()
	then.IGetUnhungry()
}
```

## Closing thoughts
> Are you familiar with writing benchmark tests and Fuzz tests?

A presenter often asks this question when discussing either topic. Despite having plenty of experience writing tests with the language, attendees often won't be familiar with it.

Sometimes, they perceive this as very different than writing regular, mundane tests.
While there is some truth to that, if you use the standard library directly rather than abstractions from assertion libraries, you are already familiar with 90% of what it takes to do either.
You'll also find it easier to debug a problem whenever something goes terribly wrong.
And when you find an opportunity to casually write a benchmark of fuzzing, you'll be able to do so in no time, maybe confidently reusing existing code.

I hope you enjoyed this blog post showing the value of the standard testing library.
If you don't, that's fine. Please take it easy as [someone is wrong on the Internet](https://xkcd.com/386/), and it might as well be me.
