# dontpanic 🛡️
> Don't Panic

## Installation

```bash
go get github.com/since1986/dontpanic
```

## Quick Start

```go
package main

import (
	"github.com/since1986/dontpanic"
)

func main() {
    // ... some code

	dontpanic.Go(func() {
		// might cause panic
	})

    // ...  some code
}
```

## Usage

### Basic Recovery

```go
dontpanic.Go(func() {
	// This panic will be automatically caught and logged
	panic("something went wrong")
})
```

### Custom Recovery Handler

```go
dontpanic.Go(
	func() { panic("custom error") },
	dontpanic.WithRecover(func(r any) {
		// Handle recovery your way
		fmt.Printf("Recovered value: %v\n", r)
	}),
)
```