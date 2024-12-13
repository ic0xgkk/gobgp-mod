# GoBGP: BGP implementation in Go

Forked from [osrg/gobgp](https://github.com/osrg/gobgp).

[![Go Report Card](https://goreportcard.com/badge/github.com/osrg/gobgp)](https://goreportcard.com/report/github.com/osrg/gobgp)
[![Tests](https://github.com/osrg/gobgp/actions/workflows/ci.yml/badge.svg)](https://github.com/osrg/gobgp/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/osrg/gobgp/v3.svg)](https://pkg.go.dev/github.com/osrg/gobgp/v3)
[![Releases](https://img.shields.io/github/release/osrg/gobgp/all.svg?style=flat-square)](https://github.com/osrg/gobgp/releases)
[![LICENSE](https://img.shields.io/github/license/osrg/gobgp.svg?style=flat-square)](https://github.com/osrg/gobgp/blob/master/LICENSE)

GoBGP is an open source Border Gateway Protocol (BGP) implementation designed from scratch for
modern environment and implemented in a modern programming language,
[the Go Programming Language](http://golang.org/).

----

## Features

Compared to the upstream repository, this fork adds the following features. Common features and bug fixes will be attempted to be merged upstream, other features will be mantained within this repository.

* Support Dynamic Multipoint WireGuard.

## Releases

Compared to the upstream repository, this fork will prepend "100" to release step number, starting from the second digit. 

For example, a new release based on version `v3.32.0` from the upstream repository will be versioned as `v3.32.1000`, `v3.32.1100`, `v3.32.1200`, etc., in this fork.

## Install

Try [a binary release](https://github.com/osrg/gobgp/releases/latest).

## Documentation

### Using GoBGP

- [Getting Started](docs/sources/getting-started.md)
- CLI
  - [Typical operation examples](docs/sources/cli-operations.md)
  - [Complete syntax](docs/sources/cli-command-syntax.md)
- [Route Server](docs/sources/route-server.md)
- [Route Reflector](docs/sources/route-reflector.md)
- [Policy](docs/sources/policy.md)
- Zebra Integration
  - [FIB manipulation](docs/sources/zebra.md)
  - [Equal Cost Multipath Routing](docs/sources/zebra-multipath.md)
- [MRT](docs/sources/mrt.md)
- [BMP](docs/sources/bmp.md)
- [EVPN](docs/sources/evpn.md)
- [Flowspec](docs/sources/flowspec.md)
- [RPKI](docs/sources/rpki.md)
- [Metrics](docs/sources/metrics.md)
- [Managing GoBGP with your favorite language with gRPC](docs/sources/grpc-client.md)
- Go Native BGP Library
  - [Basics](docs/sources/lib.md)
  - [BGP-LS](docs/sources/lib-ls.md)
  - [SR Policy](docs/sources/lib-srpolicy.md)
- [Graceful Restart](docs/sources/graceful-restart.md)
- [Additional Paths](docs/sources/add-paths.md)
- [Peer Group](docs/sources/peer-group.md)
- [Dynamic Neighbor](docs/sources/dynamic-neighbor.md)
- [eBGP Multihop](docs/sources/ebgp-multihop.md)
- [TTL Security](docs/sources/ttl-security.md)
- [Confederation](docs/sources/bgp-confederation.md)
- Data Center Networking
  - [Unnumbered BGP](docs/sources/unnumbered-bgp.md)

### Externals

- [Tutorial: Using GoBGP as an IXP connecting router](http://www.slideshare.net/shusugimoto1986/tutorial-using-gobgp-as-an-ixp-connecting-router)

## Community, discussion and support

We have the [Slack](https://join.slack.com/t/gobgp/shared_invite/zt-g9il5j8i-3gZwnXArK0O9Mnn4Yu~IrQ) for questions, discussion, suggestions, etc.

You have code or documentation for GoBGP? Awesome! Send a pull
request. No CLA, board members, governance, or other mess. See [`BUILD.md`](BUILD.md) for info on
code contributing.

## Licensing

GoBGP is licensed under the Apache License, Version 2.0. See
[LICENSE](https://github.com/osrg/gobgp/blob/master/LICENSE) for the full
license text.
