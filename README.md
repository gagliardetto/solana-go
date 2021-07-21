# Solana SDK library for Go

[![GoDoc](https://pkg.go.dev/badge/github.com/gagliardetto/solana-go?status.svg)](https://pkg.go.dev/github.com/gagliardetto/solana-go@v0.3.2?tab=doc)
[![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gagliardetto/solana-go?include_prereleases&label=release-tag)](https://github.com/gagliardetto/solana-go/releases)
[![Build Status](https://github.com/gagliardetto/solana-go/workflows/tests/badge.svg?branch=main)](https://github.com/gagliardetto/solana-go/actions?query=branch%3Amain)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gagliardetto/solana-go/main)](https://www.tickgit.com/browse?repo=github.com/gagliardetto/solana-go&branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/solana-go)](https://goreportcard.com/report/github.com/gagliardetto/solana-go)

Go library to interface with Solana JSON RPC and WebSocket interfaces.

Clients for Solana native programs, Solana Program Library (SPL), and [Serum DEX](https://dex.projectserum.com) are in development.

More contracts to come.

<div align="center">
    <img src="https://user-images.githubusercontent.com/15271561/126141780-cbd92d2c-e160-4385-9606-9094729b54d4.png" margin="auto" height="175"/>
</div>

## Installation

> :warning: `solana-go` works using SemVer but in 0 version, which means that the 'minor' will be changed when some broken changes are introduced into the application, and the 'patch' will be changed when a new feature with new changes is added or for bug fixing. As soon as v1.0.0 be released, `solana-go` will start to use SemVer as usual.

```bash
$ cd my-project
$ go get github.com/gagliardetto/solana-go@latest
```

## Current development status

There is currently **no stable release**. The SDK is actively developed and latest is `v0.3.2` which is an `alpha` release.

The RPC and WS client implementation is based on [this RPC spec](https://github.com/solana-labs/solana/blob/dff9c88193da142693cabebfcd3bf68fa8e8b873/docs/src/developing/clients/jsonrpc-api.md).

## Features

- [x] Full JSON RPC API
- [x] Full WebSocket JSON streaming API
- [ ] Wallet, account, and keys management
- [ ] Clients for native programs
- [ ] Clients for Solana Program Library
- [ ] Client for Serum
- [ ] More programs

## RPC usage

TODO

## Contributing

Any contributions are welcome, use your standard GitHub-fu to pitch in and improve.

## License

[Apache 2.0](LICENSE)

## Credits

- Gopher logo was originally created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the Creative Commons 3.0 Attributions license.
