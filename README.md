[![ci](https://github.com/safe-waters/pbar/actions/workflows/ci.yml/badge.svg)](https://github.com/safe-waters/pbar/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/safe-waters/pbar)](https://goreportcard.com/report/github.com/safe-waters/pbar)
[![Go Reference](https://pkg.go.dev/badge/github.com/safe-waters/pbar.svg)](https://pkg.go.dev/github.com/safe-waters/pbar)

# What is this?
A simple terminal progress bar

# How to use
```go
package main

import (
	"time"

	"github.com/safe-waters/pbar"
)

func main() {
	const n = 3

	p, err := pbar.New(n)
	if err != nil {
		// n must be > 0
		panic(err)
	}

	p.Start()
	for i := 0; i < n; i++ {
		p.Increment(1)
		time.Sleep(time.Second * 1)
	}
	p.End()
}

// Output (one line, overtime):
//
// 0 / 3: |--------------------------------------------------| 0.00% Complete
// 1 / 3: |████████████████----------------------------------| 33.33% Complete
// 2 / 3: |█████████████████████████████████-----------------| 66.67% Complete
// 3 / 3: |██████████████████████████████████████████████████| 100.00% Complete
```
