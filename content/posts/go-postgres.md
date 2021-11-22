---
title: "Back to basics: Writing an application using Go and PostgreSQL"
type: post
description: "Learn how to use PostgreSQL with the Go programming language using the pgx driver and toolkit in a very productive manner. Furthermore, with the provided source code, you'll be able to learn how to write efficient and sound unit and integration tests, ready to be run locally or on a Continuous Integration environment, such as GitHub Actions."
date: "2021-11-22"
image: "/img/posts/go/go-logo-blue.png"
hashtags: "golang,postgres,postgresql"
---
By reading this tutorial, you'll learn how to use [PostgreSQL](https://www.postgresql.org/) with the [Go](https://www.golang.org/) programming language using the [pgx](https://github.com/jackc/pgx) driver and toolkit in a very productive manner.
Furthermore, with the provided source code, you'll be able to learn how to write efficient and sound unit and integration tests, ready to be run locally or on a Continuous Integration environment, such as GitHub Actions.
Use the Table of Contents to skip to a specific part of this long post.
Don't forget to check out the accompanying repository [github.com/henvic/pgxtutorial](https://github.com/henvic/pgxtutorial).

<!-- Place this tag where you want the button to render. -->
<a class="github-button" href="https://github.com/henvic/pgxtutorial" data-size="large" data-show-count="true" aria-label="Star henvic/pgxtutorial on GitHub">Star</a>
<a class="github-button" href="https://github.com/henvic/pgxtutorial/fork" data-icon="octicon-repo-forked" data-size="large" data-show-count="true" aria-label="Fork henvic/pgxtutorial on GitHub">Fork</a>
<a class="github-button" href="https://github.com/sponsors/henvic" data-size="large" aria-label="Sponsor @henvic on GitHub">Sponsor</a>
<script async defer src="https://buttons.github.io/buttons.js"></script>

[![Go logo](/img/posts/go/go-logo-blue.png)](https://www.golang.org/)
[![PostgreSQL mascot Slonik](/img/posts/go-postgres/slonik_with_black_text_and_tagline.gif)](https://www.postgresql.org/)

> Check out the [api.proto file](https://github.com/henvic/pgxtutorial/blob/main/internal/api/api.proto) to see the gRPC API created for this tutorial.

{{< table_of_contents >}}

<small>Read also: [The ecosystem of the Go programming language](/posts/go/).</small>

## Context
PostgreSQL, also known as Postgres, is an extendible feature-rich [Object-Relational Database Management System](https://arctype.com/blog/postgres-ordbms-explainer/) that is almost 100% [SQL standards-compliant](https://en.wikipedia.org/wiki/SQL_compliance) and released as open source software under a permissive license.

Much of the content in this tutorial is based on experience I acquired working for [HATCH Studio](https://hatchstudio.co/), even though I'm just shy of ten months there.
My first assignments involved improving some data structures we used internally and led to a team discussion about moving from a document-based database to a more traditional relational database in the face of some challenges.
Next, we held a brainstorming session where we assembled our backend team to analyze our situation and discuss options. I already had some limited experience using PostgreSQL with pgx for a pet project and was pleased to discover that the rest of the team also saw PostgreSQL as an excellent choice for meeting our needs: great developer experience, performance, reliability, and scalability.

For our search infrastructure, we started using [Amazon OpenSearch Service](https://aws.amazon.com/opensearch-service/).
We listen to PostgreSQL database changes via its [Logical Streaming Replication Protocol](https://www.postgresql.org/docs/current/protocol-logical-replication.html) and ingest data into our [OpenSearch](https://opensearch.org/)/[Elasticsearch](https://www.elastic.co/elastic-stack/) search engine through a lightweight connector built in-house.

It works similar to hooking [Apache Kafka](https://kafka.apache.org/), but is easier to use and allows us to move faster without breaking things: running integration tests on a developer machine takes only seconds: much of which is an overkill `time.Sleep()` so we never waste time with flaky tests caused by the eventual consistency characteristics of the search engine.
This solution will not be presented now but in a future opportunity.

## tl;dr
To play with it install [Go](https://go.dev/) on your system.
You'll need to connect to a [PostgreSQL](https://www.postgresql.org/) database.
You can check if a connection is working by calling `psql`.

```sh
# Clone my repository with any of the following commands:
$ gh repo clone henvic/pgxtutorial
$ git clone https://github.com/henvic/pgxtutorial.git
$ git clone git@github.com:henvic/pgxtutorial.git
# then:
$ cd pgxtutorial
# Create a database
$ psql -c "CREATE DATABASE pgxtutorial;"
# Set the environment variable PGDATABASE
$ export PGDATABASE=pgxtutorial
# Run migrations
$ tern migrate -m ./migrations
# Run all tests passing INTEGRATION_TESTDB explicitly
$ INTEGRATION_TESTDB=true go test -v ./...
# Execute application
$ go run ./cmd/pgxtutorial
2021/11/22 07:21:21 HTTP server listening at localhost:8080
2021/11/22 07:21:21 gRPC server listening at 127.0.0.1:8082
```

Then, use [Evans](https://github.com/ktr0731/evans) to explore its gRPC endpoints by running the following command:

```sh
$ evans repl --host localhost --port 8082 -r
```

<script id="asciicast-1bVgY061isn2EHtk2wYwNXJtg" src="https://asciinema.org/a/1bVgY061isn2EHtk2wYwNXJtg.js" async></script>

<script type="text/javascript">
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_design = "enhanced_links";
amzn_assoc_asins = "B0859PF5HB";
amzn_assoc_placement = "adunit";
amzn_assoc_linkid = "14f3adb449e0071a86e28d28a1a33996";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

## SQL is alive and well
What if there is a powerful domain-specific structured query language designed to manage data out there?
It turns out there is one created almost half a century ago, in 1974: [SQL](https://en.wikipedia.org/wiki/SQL), created by Donald D. Chamberlin and Raymond F. Boyce.

Considering trade-offs, for a typical project, I prefer being able to define sound data structures up-front and having a low cost of maintenance over time, rather than a reduced cost of prototyping promised by "no-SQL" databases.

## Choosing PostgreSQL
Some technical benefits from using PostgreSQL include:
* You can use [inheritance between tables](https://www.postgresql.org/docs/current/tutorial-inheritance.html).
* You can use [JSON types](https://www.postgresql.org/docs/current/datatype-json.html) when your requirements are fluid.
* [Custom data types](https://www.postgresql.org/docs/current/sql-createtype.html) such as enum (to enumerate or restrict values for a given column) and composite (list of attribute names and data types, essentially the same as the row type of a table).

**Why not MySQL?**
MySQL is [merely relational](https://developer.okta.com/blog/2019/07/19/mysql-vs-postgres) and plagued by Oracle's controversial control over it.
In comparison, PostgreSQL is truly community-driven (see [The PostgreSQL Global Development Group](https://www.postgresql.org/developer/)).

**Why not SQLite?**
[SQLite](https://sqlite.org/) is the [most used](https://sqlite.org/mostdeployed.html) database engine in the world: from smartphones to jet engines.
It also excels in [performance](https://www.sqlite.org/speed.html), and it has [aviation-grade quality and testing](https://sqlite.org/testing.html).
However, it is not a client/server SQL database, and it is trying to solve [a different problem](https://www.sqlite.org/whentouse.html).

## Running PostgreSQL
If you haven't used PostgreSQL yet, you can start by [downloading it](https://www.postgresql.org/download/) and reading its official [documentation](https://www.postgresql.org/docs/current/) and [wiki](https://wiki.postgresql.org/).

There are several ways to run it (including Docker).
For your development machine, just do whatever is more convenient for you.
For deployments, you might want to take a look at the following:

* [Server Administration (PostgreSQL documentation)](https://www.postgresql.org/docs/current/admin.html)
* [EnterpriseDB](https://www.enterprisedb.com/)
* [Quickstart for Google Cloud SQL for PostgreSQL](https://cloud.google.com/sql/docs/postgres/quickstart)
* [Amazon RDS for PostgreSQL](https://aws.amazon.com/rds/postgresql/)
* [Azure Database for PostgreSQL](https://azure.microsoft.com/en-us/services/postgresql/)
* [Managed PostgreSQL from Heroku](https://www.heroku.com/postgres)

## PostgreSQL clients
To connect to PostgreSQL, you can use the official terminal-based front-end [psql](https://www.postgresql.org/docs/current/app-psql.html) or something else.

* [pgcli](https://www.pgcli.com/) is a CLI or [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) similar to psql, but has auto-completion, syntax highlighting and displays the docstring of functions as you type.
* [Postico](https://eggerapps.at/postico/) is a macOS GUI for PostgreSQL
* [pgAdmin](https://www.pgadmin.org/) is a feature-rich open source web-based client for PostgreSQL
* [Navicat for PostgreSQL](https://www.navicat.com/en/products/navicat-for-postgresql)

The following [terminal screen recording](https://asciinema.org/a/429446) shows me using pgcli on a half-baked personal project.

<script id="asciicast-429446" src="https://asciinema.org/a/429446.js" data-autoplay="true" data-loop="true" async></script>

### Environment variables for configuring PostgreSQL
The common and straightforward way to configure your PostgreSQL clients, tools, and applications is [to use environment variables](https://www.postgresql.org/docs/current/libpq-envars.html).
Alternatively, you can use [connection strings](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) on your application for most everything but a few options.

Example of environment variables-based configuration accepted by most tools that use a PostgreSQL database:

```shell
export PGHOST="localhost"
export PGPORT=5432
export PGUSER="username"
export PGPASSWORD="your-secret-password"
export PGDATABASE="test_whatever"

# Set a timeout for acquiring a connection to the database:
export PGCONNECT_TIMEOUT=5
```

The `PGCONNECT_TIMEOUT` (or [connect_timeout](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNECT-CONNECT-TIMEOUT) parameter) is the maximum time (in seconds) to wait for establishing a connection per host or IP address.

Generally speaking, please be advised that it's important to set it to free resources and avoid resource starvation when things are not working correctly.
For example, if you expect a request that requires a database connection to be fulfilled in 100ms on the worst case, but it already took, say, 3s because the application server cannot establish a connection to the database, it's probably already time to return an error, and let the consumer of the API decide what to do next, instead of holding the connection and taking forever to respond.

If you've read my article [Environment variables, config, secrets, and globals](/posts/env/) you might be wondering why I'm recommending environment variables here, given that I'm not a big fan of them.
In short, I don't think it's worth fighting it in this case, especially with the convenience of using the very same configuration on your application and tooling, meaning you don't need to configure your PostgreSQL connection settings in multiple places.

#### direnv
To make things more manageable, you might want to use [direnv](https://direnv.net/) to load and unload environment variables depending on what current directory you're working on at the moment.
For doing that:

1. `cd` into a directory where you want the context to be a specific database connection.
2. Create a file called .envrc with your list of environment variables (and [add it to your global .gitginore](https://docs.github.com/en/get-started/getting-started-with-git/ignoring-files)?)
3. Run the command `direnv allow` to load environment variables on your .envrc file.
4. Run the command `direnv reload` to apply changes whenever you update your .envrc file.

## Choosing a database driver for PostgreSQL
To connect to a Postgres database in Go, you'll need to use a third-party library as the standard library package doesn't provide official drivers for any databases.

However, the standard library provides interfaces for building or using database drivers (this has a good set of [goals](https://golang.org/src/database/sql/doc.txt)).

* Package `database/sql` defines a generic interface around SQL (or SQL-like) databases.
* Package `database/sql/driver` defines interfaces to be implemented by **database drivers** to be accessed via package sql, respectively.

Now, the best [SQL driver for Go](https://golang.org/s/sqldrivers) is [github.com/jackc/pgx](https://github.com/jackc/pgx).
Another good PostgreSQL driver is [github.com/lib/pq](https://github.com/lib/pq), which is now effectively in maintenance mode.

### pgx driver and toolkit
pgx provides two communication interfaces for connecting to a PostgreSQL database:

* pgx native interface to PostgreSQL
* Go's `database/sql` interface

There are some advantages of using this native interface instead of `database/sql`, as you can verify in its [README](https://github.com/jackc/pgx/blob/master/README.md).
The most interesting ones probably are:

* Faster binary format instead of a textual representation at the protocol level for datatypes different than `int64, float64, bool, []byte, string, time.Time, or nil`.
* JSON and BJSON support.
* Conversion of PostgreSQL arrays to Go slice mappings for integers, floats, and strings.

<script type="text/javascript">
amzn_assoc_tracking_id = "henvic-20";
amzn_assoc_ad_mode = "manual";
amzn_assoc_ad_type = "smart";
amzn_assoc_marketplace = "amazon";
amzn_assoc_region = "US";
amzn_assoc_design = "enhanced_links";
amzn_assoc_asins = "B0184N7WWS";
amzn_assoc_placement = "adunit";
amzn_assoc_linkid = "10528842874d7a1789b89f3c8652d0ea";
</script>
<script src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US"></script>

## Database migrations and SQL schema
[tern](https://github.com/jackc/tern) is a standalone migration tool for PostgreSQL that is part of the pgx toolkit.
You can use it to write sequential migration (even for your tests, as we'll see later).

To start using migrations, you can do something like this:

```shell
$ mkdir migrations
$ cd migrations
$ tern new <name>
```

Running `tern new product` will generate a file named `001_initial_schema.sql` with the following:

```sql
-- Write your migrate up statements here

---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
```

You can then replace this with a simple schema like the following for your first migration:

```sql
-- product table
CREATE TABLE product (
	id text PRIMARY KEY CHECK (ID != '') NOT NULL,
	name text NOT NULL CHECK (NAME != ''),
	description text NOT NULL,
	price int NOT NULL CHECK (price >= 0),
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	modified_at timestamp with time zone NOT NULL DEFAULT now()
	-- If you want to use a soft delete strategy, you'll need something like:
	-- deleted_at timestamp with time zone DEFAULT now()
	-- or better: a product_history table to keep track of each change here.
);

COMMENT ON COLUMN product.id IS 'assume id is the barcode';
COMMENT ON COLUMN product.price IS 'price in the smaller subdivision possible (such as cents)';
CREATE INDEX product_name ON product(name text_pattern_ops);

---- create above / drop below ----

DROP TABLE product;
```

The **exact** line `---- create above / drop below ----` should be present to indicate what part of the SQL code should be executed on a migration up and on a migration down.

The following commands might be useful:

* `tern new` to create a migration file prefixed by the next available sequence number (i.e., `015_add_user_account_type.sql`)
* `tern migrate` to apply your migrations on your database.
* `tern migrate -d -1` to go back one migration.
* `tern migrate -h` for help.

[tern](https://github.com/jackc/tern) keeps track of your migration version by creating and maintaining a `schema_version` table on your database.

> Tip: Be extra careful doing data migrations! It can be destructive!

<script id="asciicast-Diy7Nutmq3TWCCFdNFozbY9tl" src="https://asciinema.org/a/Diy7Nutmq3TWCCFdNFozbY9tl.js" async></script>

## main package
Let's see how we can connect to PostgreSQL using pgx from a program written in [Go](https://golang.org/).

First, to use pgx in a concurrency-safe and reliable manner, we must use a connection pool.
We can do this by using [pgx/pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool) to manage a number of low-level pgconn database driver connections in a pool.

```go
// NewPGXPool is a PostgreSQL connection pool for pgx.
//
// Usage:
// pgPool := database.NewPGXPool(context.Background(), "", &PGXStdLogger{}, pgx.LogLevelInfo)
// defer pgPool.Close() // Close any remaining connections before shutting down your application.
//
// Instead of passing a configuration explictly with a connString,
// you might use PG environment variables such as the following to configure the database:
// PGDATABASE, PGHOST, PGPORT, PGUSER, PGPASSWORD, PGCONNECT_TIMEOUT, etc.
// Reference: https://www.postgresql.org/docs/current/libpq-envars.html
func NewPGXPool(ctx context.Context, connString string, logger pgx.Logger, logLevel pgx.LogLevel) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(connString) // Using environment variables instead of a connection string.
	if err != nil {
		return nil, err
	}

	conf.ConnConfig.Logger = logger

	// Set the log level for pgx, if set.
	if logLevel != 0 {
		conf.ConnConfig.LogLevel = logLevel
	}

	// pgx, by default, does some I/O operation on initialization of a pool to check if the database is reachable.
	// Comment the following line if you don't want pgx to try to connect pool once the Connect function is called,
	//
	// If comment it, and your application seems stuck, you probably forgot to set up PGCONNECT_TIMEOUT,
	// and your code is hanging waiting for a connection to be established.
	conf.LazyConnect = true

	// pgxpool default max number of connections is the number of CPUs on your machine returned by runtime.NumCPU().
	// This number is very conservative, and you might be able to improve performance for highly concurrent applications
	// by increasing it.
	// conf.MaxConns = runtime.NumCPU() * 5

	pool, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}
	return pool, nil
}
```

We're passing an empty connString to [pgxpool.ParseConfig](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#ParseConfig) as we're using environment variables for configuration.
The pgx package will read the environment variables directly.
An equivalent [connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING) would be something along the lines of:

```
host=localhost port=5432 dbname=mydb connect_timeout=5
```

## Logging
In the preceding example, I passed a [pgx.Logger](https://pkg.go.dev/github.com/jackc/pgx#Logger) as the second parameter to the newPostgres function.
The pgx driver supports a few loggers [out of the box](https://github.com/jackc/pgx/tree/master/log).

If you use [@sirupsen](https://sirupsen.com/)'s [logrus](https://github.com/sirupsen/logrus) or Uber's [zap](https://github.com/uber-go/zap):
```go
// For github.com/jackc/pgx/v4/log/zapadapter:
pool, err := NewPGXPool(ctx, "", zapadapter.NewLogger(logger), pgx.LogLevelInfo)
defer pool.Close()

// For github.com/jackc/pgx/v4/log/logrusadapter:
pool, err := NewPGXPool(ctx, "", logrusadapter.NewLogger(logger), pgx.LogLevelInfo)
defer pool.Close()
```

If you use a logger that isn't supported out of the box, you must implement the [pgx.Logger interface](https://pkg.go.dev/github.com/jackc/pgx#Logger).

For example, for the sake of keeping this exercise simple, I'm going to create a log for Go package "log".

```go
// PGXStdLogger prints pgx logs to the standard logger.
// os.Stderr by default.
//
// This satisfies the following pgx.Logger interface:
// type Logger interface {
// 	// Log a message at the given level with data key/value pairs. data may be nil.
// 	Log(level pgx.LogLevel, msg string, data map[string]interface{})
// }
type PGXStdLogger struct {}

func (l *PGXStdLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

// For this logger:
pool, err := NewPGXPool(ctx, "", &PGXStdLogger{}, pgx.LogLevelInfo)
defer pool.Close()
```

### Log level
The following [pgx log levels](https://pkg.go.dev/github.com/jackc/pgx#LogLevel) are supported: `trace, debug, info, warn, error, and none`.

For development, the default `info` level provides a significant level of introspection into what is going on, but its verbosity is overwhelming for a deployment, and you certainly want to tune it down to `pgx.LogLevelWarn` on your production deployment.

### About pgxpool and concurrency
In Go, you've to manage concurrent access to resources yourself when you create goroutines.

[pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool) manages a pool of connections to your Postgres database, which you can safely use concurrently from multiple Goroutines, without having to worry about the thread-safety aspects with regards to individual database connections.
*pgxpool* works by managing a dynamic number of individual TCP connections to the database and controlling access, ensuring each goroutine acquiring the database connection accesses them safely (without data races).

> Remember that each request handled by net/http is handled by its own goroutine.

By reusing open connections, you also avoid the expensive round-trip costs of establishing a new TCP connection and (perhaps) performing a TLS handshake for each connection to the database.
*pgxpool's* default max number of connections is the number of CPUs on your machine.
This setting is very conservative, and you might be able to gain performance for highly concurrent applications by increasing [*pgxpool.Config.MaxConns](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool#Config).

Running benchmark tests or stress tests might help you find the sweet spot for your application.
That said, avoid the urge to do premature optimization: the default setting is perfect for most applications and use cases.

## Limited "safe" interface
I've found it mildly reassuring to create a limited [Postgres interface](https://github.com/henvic/pgxtutorial/blob/4bffea842db23624e2c1faca52d4cdde95a2b440/internal/database/interface.go) with only a subset of whitelisted methods safe to be used on the business logic of my application instead of passing `pgxpool.Pool` everywhere.
By doing so, we reduce the risk of introducing a low-level pgx API method call (such as to `pool.Close()`) without some explicit effort.

* This might live in an [internal package](https://golang.org/doc/go1.4#internalpackages), and you can probably use it most everywhere.
* `LISTEN / NOTIFY` doesn't work with it (see [chat example](https://github.com/jackc/pgx/tree/master/examples/chat)).

If you decide to use this interface, you can bypass it (helpful when debugging or testing) with the following code:

```go
// Assuming db is your postgres.PGXInterface
if pool, ok := db.(*pgxpool.Pool); ok {
	// Variable pool is a *pgxpool.Pool
	// So you can do whatever you need to do bypassing the interface now.
	
	// Print stats from the connection pool.
	fmt.Printf("%+v\n", pool.Stat())
}
```

## Database layer
Now, you might want to create an interface responsible for bridging the communication of your business logic with the database.
People call it multiple names (<abbr title="Data access object">DAO</abbr>, <abbr title="Data access layer">DAL</abbr>, repository, database, data source, etc.), and those who love to recite design patterns as poetry might say they're all different stuff.
I don't care about the naming or exact definition of such patterns.
For me, the end goal is to create a sane layer for the separation of concerns of your application and database.
In Go, the idiomatic thing to do is to [define interfaces where they're used](https://golang.org/doc/effective_go#interfaces_and_types), so let's start by that.

> Now is an excellent time to remind you that creating thousands of packages is awful! Keep things simple.

```go
// DB layer.
//go:generate mockgen --build_flags=--mod=mod -package inventory -destination mock_db_test.go . DB
type DB interface {
	// CreateProduct creates a new product.
	CreateProduct(ctx context.Context, params CreateProductParams) error

	// UpdateProduct updates an existing product.
	UpdateProduct(ctx context.Context, params UpdateProductParams) error

	// GetProduct returns a product.
	GetProduct(ctx context.Context, id string) (*Product, error)

	// SearchProducts returns a list of products.
	SearchProducts(ctx context.Context, params SearchProductsParams) (*SearchProductsResponse, error)

	// DeleteProduct deletes a product.
	DeleteProduct(ctx context.Context, id string) error

	// CreateProductReview for a given product.
	CreateProductReview(ctx context.Context, params CreateProductReviewDBParams) error

	// UpdateProductReview for a given product.
	UpdateProductReview(ctx context.Context, params UpdateProductReviewParams) error

	// GetProductReview gets a specific review.
	GetProductReview(ctx context.Context, id string) (*ProductReview, error)

	// GetProductReviews gets reviews for a given product or from a given user.
	GetProductReviews(ctx context.Context, params ProductReviewsParams) (*ProductReviewsResponse, error)

	// DeleteProductReview deletes a review.
	DeleteProductReview(ctx context.Context, id string) error
}
```

See the finalized implementation here: [internal/postgres/postgres.go](https://github.com/henvic/pgxtutorial/blob/4bffea842db23624e2c1faca52d4cdde95a2b440/internal/postgres/postgres.go).

While you can just call the database directly from your business logic, having this in a separate layer:

* Is going to be easier to maintain.
* Simplifies testing, whether you're going to be using mocks or a real implementation.

The preceding [go:generate](https://blog.golang.org/generate) comment uses mockgen to generate mocks for you can use to test code importing this interface.

For now, let's create a package named `postgres` containing an implementation of this interface.

* Alternatively, a postgres.go file on your business logic package is good enough for a small application.
* Having a strict separation in terms of interfaces and data structures has some value here.
* However, breaking down your application into too many packages might hurt readability.

Anyhow, be advised that [there are disadvantages](https://en.wikipedia.org/wiki/Data_access_object#Disadvantages) to doing this separation too:

* [Leaky abstractions](https://en.wikipedia.org/wiki/Leaky_abstraction)
* [Abstraction inversion](https://en.wikipedia.org/wiki/Abstraction_inversion)
* [Code duplication](https://en.wikipedia.org/wiki/Duplicate_code)

One way to reduce the burden of leaky abstractions is to ensure our <abbr title="Data Access Object">DAO</abbr> imports our business logic types, and not the other way around* – keeping the footprint of where our postgres package is imported to a minimum: hopefully, only the main package.
*One good thing is that Go doesn't allow import cycles.*

Two leaky abstractions to watch out for are:

1. pgx returns the `pgx.ErrNoRows` error when rows are expected but none are returned.
I've found it practical to consume this error on the database layer and return a `nil` reference pointer (or empty slice) to represent this situation.
Alternatively, you can replace the error with one defined on your service layer.
2. Reading [Rows](https://pkg.go.dev/github.com/jackc/pgx#Rows) one-by-one calling `rows.Next()` is a memory-efficient way to do things if you're doing an operation that'll require you to read a large number of rows. You've two options here: accept the leaky abstraction and deal directly with `pgx` from your service layer or do "_the right thing_", and have abstraction inversion and code duplication to solve this in a more limited manner.
Either way, you'll have to remember to `defer rows.Close()` to free resources.

## Implementation
I typically start by satisfying the interface I want to imeplement by just creating a DB struct with the methods I defined in the interface.
Next, I call panic from them.
This way I can start the implementation already starting fast integration tests right away.

```go
// DB handles database communication with PostgreSQL.
type DB struct {
	Postgres *pgxpool.Pool // Alternatively, a postgres.PGXInterface.
}

// CreateReview for a given product.
func (db *DB) CreateReview(ctx context.Context, params CreateReviewParams) error {
	panic("not implemented")
}

// UpdateReview for a given product.
func (db *DB) UpdateReview(ctx context.Context, params UpdateReviewParams) error {
	panic("not implemented")
}

// GetReview gets a specific review.
func (db *DB) GetReview(ctx context.Context, id string) (*Review, error) {
	panic("not implemented")
}

// GetReviews gets reviews for a given product or from a given user.
func (db *DB) GetReviews(ctx context.Context, params GetReviewsParams) ([]*Review, error) {
	panic("not implemented")
}

// DeleteReview from the database.
func (db *DB) DeleteReview(ctx context.Context, id string) error {
	panic("not implemented")
}

// CreateReviewFeedback with a score for a review.
func (db *DB) CreateReviewFeedback(ctx context.Context, params CreateReviewFeedbackParams) error {
	panic("not implemented")
}

// DeleteReviewFeedback removes a review feedback.
func (db *DB) DeleteReviewFeedback(ctx context.Context, params CreateReviewFeedbackParams) error {
	panic("not implemented")
}
```

<div id="amzn-assoc-ad-28708ac8-a880-4aff-b5fd-2649b98d4954"></div><script async src="//z-na.amazon-adsystem.com/widgets/onejs?MarketPlace=US&adInstanceId=28708ac8-a880-4aff-b5fd-2649b98d4954"></script>

## Packages and tools
* [pgx](https://github.com/jackc/pgx) is the SQL driver.
* [tern](https://github.com/jackc/tern) is a migration tool that is part of the pgx toolkit.
* [scany](https://github.com/georgysavva/scany) is a package for scanning from a database into Go structs and more.
* [pgtools](https://github.com/hatch-studio/pgtools) is a library containing code for testing infrastructure and more.
* [go-cmp](https://github.com/google/go-cmp) is a package for comparing Go values in tests.
* [GoMock](https://github.com/golang/mock) is a mocking framework for the Go programming language.

I've had a bad experience with ORM in the recent past, and I generally recommend against introducing this sort of abstraction to code.
I prefer much more something much closer to using pgx with scany and `pgtools.Wildcard()` than to a full-featured ORM.
I don’t have a strong opinion regarding query builders, but I see them more as a liability on the supply chain.

Inspiration-wise, the pkg.go.dev website is powered by [pkgsite](https://github.com/golang/pkgsite), which uses PostgreSQL as a database.
Nowadays, I regularly check out their repository to learn more and see what they use.

## Testing
Opt-in for tests that require a database connection by means of:

* test flags (`-integration`, which works on local directory mode, and not with list mode)
* test build flags (`-tags=integration`)
* environment variables

At work, we decided to use test build flags. Example: `go test -tags=integration.`
I prefer environment variables instead, as I'm not too fond of the hassle introduced by build tags on tooling, such as text editor or [gopls](https://pkg.go.dev/golang.org/x/tools/gopls), which might stop processing them [as intended](https://github.com/golang/go/issues/29202).


### Running tests without database
You want your tests to pass even when running without a database for the sake of keeping your developer experience smooth.
For example, in the pgxtutorial repository, I resorted to skipping all tests of an entire package at once for some packages if the `INTEGRATION_TESTDB` environment variable is not set.
[I did this](https://github.com/henvic/pgxtutorial/blob/4bffea842db23624e2c1faca52d4cdde95a2b440/internal/inventory/inventory_test.go#L20-L26) thinking about maintainability: it was way easier and safer than to keep tests that didn’t rely on a database passing.

```go
func TestMain(m *testing.M) {
	if os.Getenv("INTEGRATION_TESTDB") != "true" {
		log.Printf("Skipping tests that require database connection")
		return
	}
	os.Exit(m.Run())
}
```

When this general approach is not welcome, we can skip an individual test directly:

```go
if os.Getenv("INTEGRATION_TESTDB") == "true" {
	t.Skip("skipping test that require database connection")
}
```

### Running test with pgtools and tern
The following code shows how we can use pgtools/sqltest to run integration tests against a real database:

```go
var force = flag.Bool("force", false, "Force cleaning the database before starting")

func TestCreateProduct(t *testing.T) {
	t.Parallel()
	migration := sqltest.New(t, sqltest.Options{
		Force: *force,
		Path:  "../../migrations",
	})
	pool := migration.Setup(context.Background(), "")

	db := &DB{
		Postgres: pool,
	}
	// ...
}
```

Check the pgtools/sqltest documentation to learn about all existing options.
By default, a database is temporarily created for the test function that called `migration.Setup()`.
Using pgtools/sqltest, you can feel safe knowing that the created database is prefixed with "test" to ensure it doesn't clash with existing databases.
If you're running tests on multiple packages, set a unique value for the `TemporaryDatabasePrefix` field.

### Table-driven tests
I often use table-driven tests when writing integration tests, and I usually find it easier to call migration.Setup() in the parent test and pass it ahead.
However, you must be aware that this might introduce a problem of reliability: a test that depended on another might fail when running isolated with `go test -run=Foo/bar`.
Finding the right balance is essential here, as we want to identify problematic behaviors that might not be perceived if we were to have fully isolated tests.

### Real implementation vs. using test doubles
Taking a look at the `internal/inventory/inventory` package tests you'll notice it mostly uses a real implementation for its tests.
It contains only a thin layer of test doubles using [GoMock](https://github.com/golang/mock) to verify if parameters are being passed correctly or to simulate database errors.

We generated the mock in the repository by calling the following command:

```sh
$ mockgen --build_flags=--mod=mod -package inventory -destination mock_db_test.go . DB
```

You can add the following directive on your code to do the same when running `go generate`:

```go
//go:generate mockgen --build_flags=--mod=mod -package inventory -destination mock_db_test.go . DB
```

The code for your test cases relying on a mock is a little bit awkward to get used to but are powerful and easy to use once you understand how they work. Example:

```go
ctrl := gomock.NewController(t)
m := inventory.NewMockDB(ctrl)
m.EXPECT().CreateProduct(gomock.Not(gomock.Nil()),
	inventory.CreateProductParams{
		ID:          "simple",
		Name:        "product name",
		Description: "product description",
		Price:       150,
	}).Return(errors.New("unexpected error"))
return m
```

For a full example, see [TestServiceCreateProduct/database_error](https://github.com/henvic/pgxtutorial/blob/4bffea842db23624e2c1faca52d4cdde95a2b440/internal/inventory/inventory_test.go#L123-L147).

### Comparing structures
You can use the package [go-cmp](https://github.com/google/go-cmp) to compare different structs easily.

In some circumstances ([test example](https://github.com/henvic/pgxtutorial/blob/4bffea842db23624e2c1faca52d4cdde95a2b440/internal/inventory/inventory_test.go#L425-L434)), we've set the fields CreatedAt and ModifiedAt for some values returned from the database.
In others, we decided to ignore it.

```go
// Ignoring field generated automatically:
if !cmp.Equal(tt.want, got, cmpopts.IgnoreFields(inventory.ProductReview{}, "ID")) {
	t.Errorf("value returned by Service.GetProductReview() doesn't match: %v", cmp.Diff(tt.want, got))
}

// Several ways to check for equality treating time values as special:
if !cmp.Equal(tt.want, got, cmpopts.EquateApproxTime(time.Minute)) {
	t.Errorf("value returned by Service.GetProduct() doesn't match: %v", cmp.Diff(tt.want, got))
}

if !cmp.Equal(tt.want, got, cmpopts.IgnoreFields(inventory.Product{}, "CreatedAt", "ModifiedAt")) {
	t.Errorf("value returned by DB.GetProduct() doesn't match: %v", cmp.Diff(tt.want, got))
}

if !cmp.Equal(tt.want, got, cmpopts.EquateApproxTime(time.Minute)) {
	t.Errorf("value returned by Service.GetProduct() doesn't match: %v", cmp.Diff(tt.want, got))
}

if !cmp.Equal(tt.want, got, cmpopts.EquateApproxTime(time.Minute)) {
	t.Errorf("value returned by Service.SearchProducts() doesn't match: %v", cmp.Diff(tt.want, got))
}

if !cmp.Equal(tt.want, got, cmpopts.IgnoreTypes(time.Time{})) {
	t.Errorf("value returned by DB.GetProductReviews() doesn't match: %v", cmp.Diff(tt.want, got))
}
```

Alternatively, you can evaluate the returned values and then copy CreatedAt and ModifiedAt before comparing structs.
For example:

```go
if got.CreatedAt.IsZero() {
	t.Error("Service.GetProductReview() returned CreatedAt should not be zero")
}
if !got.CreatedAt.Before(got.ModifiedAt) {
	t.Error("Service.GetProductReview() should return CreatedAt < ModifiedAt")
}
tt.want.CreatedAt = got.CreatedAt
tt.want.ModifiedAt = got.ModifiedAt
// Now, compare structures here.
```

### Failed idea: replay testing
I've experimented with creating a replay testing framework [built on top of pgmock](https://github.com/henvic/pgmock/tree/feat/pgplayback/pgplayback) to run tests without requiring a database, but the idea wasn't practical.
Besides, the main reason for having it wasn't justifiable: the integration tests are fast.
### Running tests
```sh
$ INTEGRATION_TESTDB=true go test -v ./...
```

You can also add to your .zshrc or .envrc and run `go test` directly:

```sh
export INTEGRATION_TESTDB=true
```

For running the integration tests on GitHub Actions, we need first to create a workflow such as the following:

```yml
name: Integration
on:
  pull_request:
    types: [opened, synchronize, reopened, ready_for_review]
  push:
    branches:
      - main
permissions:
  contents: read
  pull-requests: read
jobs:
  # Reference: https://docs.github.com/en/actions/guides/creating-postgresql-service-containers
  postgres-test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres
        env:
          POSTGRES_USER: runner
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: test_pgxtutorial
        options: >-
          --name postgres
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          # Maps tcp port 5432 on service container to the host
          - 5432:5432
    env:
      INTEGRATION_TESTDB: true
      PGHOST: localhost
      PGUSER: runner
      PGPASSWORD: postgres
      PGDATABASE: test_pgxtutorial
    steps:
    - uses: actions/checkout@v1
    - uses: actions/setup-go@v1
      with:
        go-version: '1.17.x'
    - name: Run Postgres tests
      run: go test -v -race -count 1 -covermode atomic -coverprofile=profile.cov ./...
    - name: Code coverage
      if: ${{ github.event_name != 'pull_request' }}
      uses: shogo82148/actions-goveralls@v1
      with:
        path-to-profile: profile.cov
```

{{< tweet henriquev 1462723601703096322 >}}

I hope you enjoyed this tutorial. Now go ahead and checkout the repository [github.com/henvic/pgxtutorial](https://github.com/henvic/pgxtutorial).

<small>If you click and buy any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/).</small>
