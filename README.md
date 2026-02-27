# Wazero Atomic

A [wazero](https://pkg.go.dev/github.com/tetratelabs/wazero) host module, ABI and guest SDK providing atomic data 
structures suitable for sharing data between concurrent WASI modules.

## Host Module

[![Go Reference](https://godoc.org/github.com/pantopic/wazero-atomic/host?status.svg)](https://godoc.org/github.com/pantopic/wazero-atomic/host)
[![Go Report Card](https://goreportcard.com/badge/github.com/pantopic/wazero-atomic/host)](https://goreportcard.com/report/github.com/pantopic/wazero-atomic/host)
[![Go Coverage](https://github.com/pantopic/wazero-atomic/wiki/host/coverage.svg)](https://raw.githack.com/wiki/pantopic/wazero-atomic/host/coverage.html)

First register the host module with the runtime

```go
import (
    "github.com/tetratelabs/wazero"
    "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

    "github.com/pantopic/wazero-atomic/host"
)

func main() {
    ctx := context.Background()
    r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig())
    wasi_snapshot_preview1.MustInstantiate(ctx, r)

    module := wazero_atomic.New()
    module.Register(ctx, r)

    // ...
}
```

## Guest SDK (Go)

[![Go Reference](https://godoc.org/github.com/pantopic/wazero-atomic/sdk-go?status.svg)](https://godoc.org/github.com/pantopic/wazero-atomic/sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/pantopic/wazero-atomic/sdk-go)](https://goreportcard.com/report/github.com/pantopic/wazero-atomic/sdk-go)

Then you can import the guest SDK into your WASI module to send messages from one WASI module to another.

```go
package main

import (
    "github.com/pantopic/wazero-atomic/sdk-go"
)

var n *atomic.Uint64

func main() {
    n = atomic.NewUint64()
}

//export test
func test() {
    println(n.Add(1)) // 1
    println(n.Add(2)) // 3
}
```

## Roadmap

This project is in alpha. Breaking API changes should be expected until Beta.

- `v0.0.x` - Alpha
  - [ ] Stabilize API
- `v0.x.x` - Beta
  - [ ] Finalize API
  - [ ] Test in production
- `v1.x.x` - General Availability
  - [ ] Proven long term stability in production
