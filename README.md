go-scoro
========

[![GoDoc](https://godoc.org/github.com/lxmx/go-scoro?status.png)](https://godoc.org/github.com/lxmx/go-scoro)

go-scoro is an Go client library for [Scoro API](https://api.scoro.com/api/)

### Getting started

#### Installation

`go get -u gopkg.in/lxmx/go-scoro.v1` or `govendor fetch gopkg.in/lxmx/go-scoro.v1`

#### Documentation

API documentation and examples are available via [godoc](https://godoc.org/github.com/lxmx/go-scoro).

If you need to check documentation locally before commit, run `godoc -http ":6060"` and open [http://localhost:6060/pkg/github.com/lxmx/go-scoro/](http://localhost:6060/pkg/github.com/lxmx/go-scoro/) in browser.

#### Examples

The [examples](./examples) directory contains more elaborate example applications.

`go run examples/<path>/<file>.go -company <company_id> -api_key <api_key> -subdomain <subdomain>`

#### Note:

The library does not have implementations of all Scoro resources. PRs for new resources and endpoints are welcome.
