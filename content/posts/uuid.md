---
title: "You don't need UUID"
type: post
description: "I've experienced first-hand how using UUID hurts the usability of computer systems, and I want you to understand why you certainly don't need it."
date: "2021-05-30"
image: "/img/posts/uuid/matrix.png"
hashtags: "programming"
---
[UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier), short for Universally Unique Identifier ([spec](https://www.rfc-editor.org/info/rfc4122)), is a 128-bit identifier format widespread on computer systems.
The following is an example, using its prevalent representation: <samp>a73ba12d-1d8b-2516-3aee-4b15e563a835</samp>.
I've experienced first-hand how using UUID hurts the usability of computer systems, and I want you to understand why you certainly don't need it.

Take Amazon. Which do you think is a link to one of their products?

* [amzn.to/3c6n63N](https://amzn.to/3c6n63N)
* amzn.to/a73ba12d-1d8b-2516-3aee-4b15e563a835

A simple ID like <samp>3c6n63N</samp> is more than enough to represent any product while keeping it readable and making communication easier.
A UUID alternative like <samp>a73ba12d-1d8b-2516-3aee-4b15e563a835</samp> is just wasteful from an user's perspective.

In a specific case, I’ve witnessed entities that were supposed to be exposed to end-users using user-friendly ids while UUIDs were in use internally.
For this to work, either a one-to-one mapping or a separate list of unique ids for your entities are necessary.

> You might as well end up with a combination of the worst of both worlds: UUID + sequentially generated (auto increment) numeric identifiers.

I’ve seen all sorts of combinations of this on many systems: some entities using only friendly ids, some using only sequential ids, and most using both.

For example, an entity that is used only internally might be using a UUID.
Then a demand to expose it externally appears, and developers add a new friendly ID so that users won't deal with ridiculously long links.

* When using UUID, it's hard to use, track, and compare data manually.
* When using sequential ids, you might be leaking sensitive business metrics to the competition or hit scalability and syncing limits.
* When using both at the same time, your internal users might have a hard time mapping ids to UUIDs and vice-versa.

So, with all this said, I think I've made a case for trying to just stick to more accessible ids everywhere!
Next, watch this video, and let's see a practical alternative.

<div class="grid-x">
        <div class="medium-6 small-12">
                {{< youtube gocwRvLhDf8 >}}
                <p><small>
                <a href="https://www.youtube.com/watch?v=gocwRvLhDf8">Will YouTube Ever Run Out Of Video IDs?</a> by <a href="https://www.tomscott.com/">Tom Scott</a>
                </small></p>
        </div>
</div>

## Collisions and uniqueness
The UUID textual representation is 36 characters long, being four hyphen separators and 32 hexadecimal digits.
There are four versions.
Version 1 and 2 were date-time and MAC address-based.
Version 3 and 5 are namespace name-based.
Version 4 is completely randomly generated (hence, it has more entropy) and is what most web systems seem to use.

It has 16^32 = 2^128 bits that guarantee uniqueness and has an insignificant risk of collision.

> Hexadecimal is commonly used in computing as a compact representation for the binary numeral system.
> In hexadecimal, 10 is 0xA, 11 is 0xB, 12 is 0xC, 13 is 0xD, 14 is 0xE, 15 is 0xF, 16 is 0x10, and so on.
> Read [Binary number: conversion to and from other numeral systems](https://en.wikipedia.org/wiki/Binary_number#Conversion_to_and_from_other_numeral_systems) if you want to learn or review how it works.

## Practical solutions
As Tom Scott shows in his video, 11 base58-encoded characters are enough for YouTube to serve content even when considering that private videos should be undiscoverable.

Let's see what a simple and elegant solution for generating that in Go might be:

```go
// NewID generates a random base-58 ID.
func NewID() string {
	const (
		alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz" // base58
		size     = 11
	)

        var id = make([]byte, size)
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}
	for i, p := range id {
		id[i] = alphabet[int(p)%len(alphabet)] // discard everything but the least significant bits
	}
	return string(id)
}
```

<p><a href="https://play.golang.org/p/T3wvtUNSJcy" lang="en-US" class="button secondary">Play with this code</a></p>

This solution uses the human-readable [base58 encoding scheme](https://tools.ietf.org/id/draft-msporny-base58-01.html).
I cheated a little by using only the least significant bits to create the ID instead of trying to squeeze performance as this is good enough.

Thank you for reading this article.

{{< tweet 1399330361126031363 >}}

<script type="text/javascript">
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_design = "enhanced_links";
amzn_assoc_asins = "0984782850";
amzn_assoc_placement = "adunit";
amzn_assoc_linkid = "65cdf509fe5c5bb4f9b616660057129c";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

If you click and buy any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/).
