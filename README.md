[![codecov](https://codecov.io/gh/gofika/weightsystem/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/weightsystem)
[![Build Status](https://github.com/gofika/weightsystem/workflows/build/badge.svg)](https://github.com/gofika/weightsystem)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/weightsystem)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/weightsystem)](https://goreportcard.com/report/github.com/gofika/weightsystem)
[![Licenses](https://img.shields.io/github/license/gofika/weightsystem)](LICENSE)

# weightsystem

A generic weight system based on Go generics


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/weightsystem
```

### Example

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/gofika/weightsystem"
)

func displayWeights(ws *weightsystem.WeightSystem[string]) {
	for item, weight := range ws.Weights() {
		fmt.Printf("Item: %#+v, Weight: %.2f\n", item, weight)
	}
}

func main() {
	ipList := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
	ipSystem := weightsystem.New[string]()
	// Add items to the system
	ipSystem.AddItems(ipList)

	// randomly adjust weights
	for i := 0; i < 5000; i++ {
		ip := ipSystem.GetItem()
		success := rand.Float32() < 0.7
		ipSystem.AdjustWeight(ip, success)
	}

	fmt.Println("Current weights:")
	displayWeights(ipSystem)

	// Add new items to the system with AVG weight
	ipSystem.AddItem("192.168.1.4")

	// Add new items to the system with custom weight
	ipSystem.AddItem("192.168.1.5", weightsystem.WithWeight(150.0))

	// Add new items to the system with custom weight. out of range, will set to minimum weight
	ipSystem.AddItem("192.168.1.6", weightsystem.WithWeight(-50.0))

	fmt.Println("\nWeights after adding new items:")
	displayWeights(ipSystem)
}
```