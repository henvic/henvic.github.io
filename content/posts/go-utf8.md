---
title: "UTF-8 strings with Go: len(s) isn't enough"
type: post
description: "In this post, I show you the bare minimum you need to know how to do UTF-8 string manipulation in Go safely."
date: "2022-03-07"
image: "/img/posts/go/go-logo-blue.png"
hashtags: "golang"
---
In this post, I show you the bare minimum you need to know how to do UTF-8 string manipulation in Go safely.

<small>Read also: [Back to basics: Writing an application using Go and PostgreSQL](/posts/go-postgres/) and [Homelab: Intel NUC with the ESXi hypervisor](/posts/homelab/).</small>

# tl;dr
Use the `unicode/utf8` [package](https://pkg.go.dev/unicode/utf8) to:

1. Validate if string isn't in another encoding or corrupted:

```go
fmt.Println(utf8.ValidString(s))
```

2. Get the right number of runes in a UTF-8 string:

```go
fmt.Println(utf8.RuneCountInString("é um cãozinho")) // returns 13 as expected
fmt.Println(len("é um cãozinho")) // returns 15 because 'é' and 'ã' are represented by two bytes each
```

1. Strings might get corrupted if you try to slice them incorrectly:

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var dog = "é um cãozinho"
	dog = dog[1:]
	fmt.Printf("got: %q (valid: %v)\n", dog, utf8.ValidString(dog))
}

// Output:
// got: "� um cãozinho" (valid: false)
```

To slice them correctly, use `utf8.DecodeRune` or `utf8.DecodeRuneInString` to get the first rune and its size:

```go
func main() {
	var dog = "é um cãozinho"
        _, offset := utf8.DecodeRuneInString(dog)
	dog = dog[offset:]
	fmt.Printf("got: %q (valid: %v)\n", dog, utf8.ValidString(dog))
}

// Output:
// got: " um cãozinho" (valid: true)
```

# Background

In the early days of the web, websites from different regions used different [character encodings](https://en.wikipedia.org/wiki/Character_encoding), accordingly to their demographic region. Nowadays, most websites use the [Unicode](https://en.wikipedia.org/wiki/Unicode) implementation known as UTF-8. Unicode defines 144,697 characters.
Here are some of the most popular encodings:

| Encoding | Use |
| ------------- |-------------|
| [UTF-8](https://en.wikipedia.org/wiki/UTF-8) | Unicode Standard (International) |
| [ISO-8859-1](https://en.wikipedia.org/wiki/ISO/IEC_8859-1) | Western European languages (includes English) |
| [ISO-8859-2](https://en.wikipedia.org/wiki/ISO/IEC_8859-2) | Eastern European languages |
| [ISO-8859-5](https://en.wikipedia.org/wiki/ISO/IEC_8859-5) | Cyrillic languages |
| [GB 2312](https://en.wikipedia.org/wiki/GB_2312) | Simplified Chinese |
| [Shift JIS](https://en.wikipedia.org/wiki/Shift_JIS) | Japanese |
| [Windows-125x series](https://en.wikipedia.org/wiki/Windows_code_page) | Windows code pages: characters sets for multiple languages |
| ... | ... |

UTF-8 was created in 1992 by Ken Thompson and Rob Pike as a variable-width character encoding, originally implemented for the [Plan 9](https://en.wikipedia.org/wiki/Plan_9_from_Bell_Labs) operating system. It is backward-compatible with ASCII. As of 2022, more than 97% of the content on the web is encoded with UTF-8.
See:

* [Unicode over 60 percent of the web](https://googleblog.blogspot.com/2012/02/unicode-over-60-percent-of-web.html) (Google, February 3, 2012)
* [Historical yearly trends in the usage statistics of character encodings for websites](https://w3techs.com/technologies/history_overview/character_encoding/ms/y) (since 2011)

# Why should you care?
Take a language with just a few extra glyphs like Portuguese or Spanish, and you'll quickly notice the importance of handling encodings properly when writing software.
To show that, let me write a small program that will iterate over a string rune by rune [assuming it is UTF-8](https://go.dev/blog/strings) and print its representation:

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var examples = "abcdeáãàâéíóõúûüç世界😁"
	for _, c := range examples {
		fmt.Printf("%#U\tdecimal: %d\tbinary: %b \tbytes: %d\n", c, c, c, utf8.RuneLen(c))
	}
}

// Output:
// U+0061 'a'	decimal: 97	binary: 1100001 		bytes: 1
// U+0062 'b'	decimal: 98	binary: 1100010 		bytes: 1
// U+0063 'c'	decimal: 99	binary: 1100011 		bytes: 1
// U+0064 'd'	decimal: 100	binary: 1100100 		bytes: 1
// U+0065 'e'	decimal: 101	binary: 1100101 		bytes: 1
// U+00E1 'á'	decimal: 225	binary: 11100001 		bytes: 2
// U+00E3 'ã'	decimal: 227	binary: 11100011 		bytes: 2
// U+00E0 'à'	decimal: 224	binary: 11100000 		bytes: 2
// U+00E2 'â'	decimal: 226	binary: 11100010 		bytes: 2
// U+00E9 'é'	decimal: 233	binary: 11101001 		bytes: 2
// U+00ED 'í'	decimal: 237	binary: 11101101 		bytes: 2
// U+00F3 'ó'	decimal: 243	binary: 11110011 		bytes: 2
// U+00F5 'õ'	decimal: 245	binary: 11110101 		bytes: 2
// U+00FA 'ú'	decimal: 250	binary: 11111010 		bytes: 2
// U+00FB 'û'	decimal: 251	binary: 11111011 		bytes: 2
// U+00FC 'ü'	decimal: 252	binary: 11111100 		bytes: 2
// U+00E7 'ç'	decimal: 231	binary: 11100111 		bytes: 2
// U+4E16 '世'	decimal: 19990	binary: 100111000010110 	bytes: 3
// U+754C '界'	decimal: 30028	binary: 111010101001100 	bytes: 3
// U+1F601 '😁'	decimal: 128513	binary: 11111011000000001 	bytes: 4
```

Each of these preceding single characters is represented by what in [character encoding](https://en.wikipedia.org/wiki/Character_encoding) terminology is called a [code point](https://en.wikipedia.org/wiki/Code_point), a numeric value that computers use to map, transmit, and store.
Now, UTF-8 is a [variable-width encoding](https://en.wikipedia.org/wiki/Variable-width_encoding) requiring one to four bytes (that is, 8, 16, 24, or 32 bits) to represent a single ~~character~~ code point.
UTF-8 uses one byte for the first 128 code points (backward-compatible with ASCII), and up to 4 bytes for the rest.
While UTF-8 and many other encodings, such as ISO-8859-1, are backward-compatible with ASCII, their extended codespace aren't compatible between themselves.

> In the Go world, a code point is called a rune.

From Go's [src/builtin/builtin.go](https://cs.opensource.google/go/go/+/master:src/builtin/builtin.go;l=90-92) definition of rune, we can see it uses an int32 internally for each code point:

```go
// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```

# How can this affect you, anyway?
In the early days of the web (though this problem might happen elsewhere), whenever you tried to access content from another demographics for which your computer vendor didn't prepare it to handle, you'd likely get a long series of □ or � replacement characters on your browser.

If you wanted to display, say, Japanese or Cyrillic correctly, not only you'd have to download a new font:
There was also a high chance of the website not setting encoding correctly, forcing you to manually adjust it on your browser (and hope it works).

# unicode/utf8
With Unicode and UTF-8, this became a problem of the past.
Surely, I still cannot read any non-Latin language, but at least everyone's computers now render beautiful Japanese or Chinese calligraphy just fine out of the box.

From a software development perspective, we need to be aware of several problems, such as how to handle strings manipulation correctly, as we don't want to cause data corruption.

For that, when working with UTF-8 in Go, if you need to do any sort of string manipulation, such as truncating a long string, you'll want to use the [unicode/utf8 package](https://pkg.go.dev/unicode/utf8).

**Length of a string vs. the number of runes:**
What is the length of the following words?

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var examples = []string{
		"dog",
		"cão",
		"pa",
		"pá",
		"pà",
		"Hello, World",
		"Hello, 世界",
	}
	for _, w := range examples {
		fmt.Printf("%s\tlen: %d\trunes: %d\n", w, len(w), utf8.RuneCountInString(w))
	}
}

// Output:
dog	len: 3	runes: 3
cão	len: 4	runes: 3
pa	len: 2	runes: 2
pá	len: 3	runes: 2
pà	len: 3	runes: 2
Hello, World	len: 12	runes: 12
Hello, 世界	len: 13	runes: 9
```

As you can see, len cannot be used to count the number of runes properly.
Instead, you must use the following functions:

```go
utf8.RuneCount(p []byte) int
utf8.RuneCountInString(s string) (n int)
```

To validate if a string is consists entirely of valid UTF-8-encoded runes use the following functions:

```go
utf8.Valid(p []byte) bool
utf8.ValidRune(r rune) bool
utf8.ValidString(s string) bool
```

Example:

```go
fmt.Println(utf8.ValidString("isso é um exemplo")) // Should print true.
fmt.Println(utf8.ValidString("\xe0")) // Should print false.
```

# Converting between encodings
To convert to/from other encodings, you might use the [golang.org/x/text/encoding/charmap package](https://pkg.go.dev/golang.org/x/text/encoding/charmap).

Example:

```go
package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	var invalid = "\xe0 \xe1 \xe2 \xe3 \xe9" // this is "à á â ã é" in ISO-8859-1
	fmt.Printf("Not UTF-8: %q (valid: %v)\n", invalid, utf8.ValidString(invalid))

	// If we convert it from ISO8859-1 to UTF-8:
	dec, _ := charmap.ISO8859_1.NewDecoder().String(invalid)
	fmt.Printf("Decoded: %q (valid UTF8: %v)\n", dec, utf8.ValidString(dec))
}

// Output:
// Not UTF-8: "\xe0 \xe1 \xe2 \xe3 \xe9" (valid: false)
// Decoded: "à á â ã é" (valid UTF8: true)
```

If you need help reading a malformed string, or getting runes individually, read the documentation for the [unicode/utf8 package](https://pkg.go.dev/unicode/utf8) and check out its examples.

# References

* [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) by Rob Pike.
* [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/) by Joel Spolsky.
