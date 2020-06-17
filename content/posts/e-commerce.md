---
title: "I'm starting an opensource e-commerce project."
type: post
description: "I'm starting an opensource e-commerce project."
date: "2020-06-17"
image: "/img/posts/e-commerce/market-product-page.png"
hashtags: "e-commerce,opensource,public domain,project"
---
1. Most e-commerce platforms operate in a way that makes stores relying on them hostage due to [vendor lock-in](https://en.wikipedia.org/wiki/Vendor_lock-in) models.
2. E-commerce is something that can get complex fast. Search engine optimization (SEO), inventory management, order status, and user experience are essential and can get hard quickly. Especially when in light of the danger of trying to make everything customizable.
3. Substantial transaction fees drive away businesses operating on tight profit margins.

[market](https://github.com/henvic/market) is an opensource e-commerce software dedicated to the public domain to help small and medium businesses to establish their online presence with a high-quality website without gotchas.

<a href="/img/posts/e-commerce/market-product-page.png"><img src="/img/posts/e-commerce/market-product-page_small.png" alt="market product page" width="600"></a>

## Architecture
market has a core application (read: monolith) that handles all requests to its API and regular web pages - both served by separate hosts.
Requests received via `www.` are handled by the web pages HTTP handlers. The API HTTP endpoints handle requests received via `api.`

[Go](https://golang.org/) has been my language of choice for years, and it is perfect for this kind of project thanks to its simplicity and straight-forward 'no magic' approach to software engineering.

For the frontend of the application, I am going to rely on [React](https://reactjs.org/) and, to be on the safe side of types, I am going to be using [TypeScript](https://www.typescriptlang.org/). Having had a long pause from frontend development, I still don't know how well these two play together, but I am not worried too much.

<table class="stack">
        <caption>External services</caption>
        <tr>
                <td>PostgreSQL</td>
                <td>Relational database to store everything safely.</td>
        </tr>
        <tr>
                <td>ElasticSearch</td>
                <td>Search engine for the search, autocompletion, and category browsing.</td>
        </tr>
        <tr>
                <td>imaginary (or thumbor)</td>
                <td>Photo thumbnail service for resizing images for distribution.</td>
        </tr>
        <tr>
                <td>MinIO / AWS S3 API</td>
                <td>Object storage service compatible with Amazon S3 API.</td>
        </tr>
</table>

### Databases
[PostgreSQL](https://www.postgresql.org/) is a trustworthy ACID-compliant relational database, probably the best choice for this kind of project.
My knowledge regarding databases is limited, but I know this is the right choice. I recommend you read [Things I Wished More Developers Knew About Databases](https://medium.com/@rakyll/things-i-wished-more-developers-knew-about-databases-2d0178464f78) by [Jaana Dogan](https://twitter.com/rakyll) if you are interested in databases.

Keeping track of changes is valuable for an online store. However, unfortunately, PostgreSQL, as most relational databases, does not provide versioning natively. We can use a table for storing logs of changes. I could probably look into using something like a [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree) and something more complex to store a register, but it would be hard and have its problems (example: should purging unwanted data be possible?), so this seems like a good enough trade-off.

### Search
ElasticSearch has excellent search features such as faceted filters, and I am going to take advantage of it. The inventory should be synced from PostgreSQL to ElasticSearch somehow. I am not sure yet what would be the best approach for this today. Perhaps using a PostgreSQL extension to automate this is the best option.

About seven years ago, I was working on a marketplace for used vehicles and achieved eventual consistency with a best-effort synchronization strategy:

* On updates, a message would be sent to a background worker to reindex the altered document.
* An additional crontab would run regularly to fix any consistency issues.

<div class="grid-x">
        <div class="medium-6 small-12">
                {{< youtube dML0FQIUcTY >}}
        </div>
</div>

### Photo thumbnail service
I had a great experience with <a href="https://github.com/thumbor/thumbor">thumbor</a> in the past, and I believe it's the most battle-tested thumbnail service out there.

I'm giving <a href="https://github.com/h2non/imaginary">imaginary</a> a try for this, though, because it is a Go project, and I'd love to see how well it does.

<em>The first program I wrote with Go was a photo thumbnail service, too: <a href="https://github.com/henvic/picel">picel</a>.</em>

Images files are stored in the storage service and served through the photo thumbnail service where resizing and optimization should happen. Caching might happen either at this layer or in the load-balancer layer (more likely).

### Compatibility guarantee
Releases will follow [Semantic versioning](https://semver.org/). I intend new versions to be backward compatible unless there exists an excellent reason why not to.
Hopefully, version 1 is going to have a long run. A critical aspect of this is dealing with data structures wisely.

This guarantee assumes no modifications were made to the software, in any case. While this software is intended to be easy to read, and easy to contribute to, it is not intended to be fully customized, but used as-is. Some general guidelines about customization should be written about this later on. Contributions respecting this goal are always welcome.

## Usage
Ideally, a grocery store or car shop could:

1. Buy a domain;
2. Create an account on a cloud computing platform;
3. Point his DNS to the Software-as-a-Service provider;
4. Deploy a free market appliance and have a store up and running in minutes with HTTPS and everything else out-of-the-box;
5. Import inventory data using a computer or smartphone;
6. Configure payment options (payment gateway, gift cards, shop credit, etc.);
7. Start selling online;

Moreover, if they are small enough, they can get away doing it almost for free. I admit this is a long shot, but it can happen. Cloud providers such as Google Cloud, AWS, and Azure already provide some products that make this easier to achieve, such as [Cloud Functions](https://cloud.google.com/functions/).

Due to latency concerns, though, one day, they might decide to use a lightweight <a href="/posts/homelab">Intel NUC server</a> to run everything on-premises. For example, if they are opening up shop in a rocket flying to Mars. Thanks to the 'no strings attached' concept of the public domain, they can do it without having to ask.

## A few more points
* Seach Engine Optimization using techniques such as structured data to drive high conversion rates.
* Freedom of payment gateway: it should be easy to integrate with the likes of PayPal, Square, Stripe, and so on. Or even no gateway.
* Bye-bye vendor lock-in: if the solution is not for you, just take over.
* Rely on Go simplicity to make it simple to write high-quality e-commerce software.
* Follow the Go proverbs.
* Safety first: third-party code and dependencies audited before usage.
* Single sign-on is very interesting. So it is a traditional login. So is Multi-Factor Authentication (MFA). It is okay not to have everything from day one.
