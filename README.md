# golang-hash

This Go package contains functions that produce non-cryptographic hashes in various bit sizes. These functions use the Fowler–Noll–Vo hash (FNV-1a) algorithm.

[![Go Reference](https://pkg.go.dev/badge/github.com/theTardigrade/golang-hash.svg)](https://pkg.go.dev/github.com/theTardigrade/golang-hash) [![Go Report Card](https://goreportcard.com/badge/github.com/theTardigrade/golang-hash)](https://goreportcard.com/report/github.com/theTardigrade/golang-hash)

## Example

```golang
package main

import (
	"fmt"

	hash "github.com/theTardigrade/golang-hash"
)

func main() {
	const input = "this is a simple test"

	fmt.Printf(
		"%x\n\n%x\n\n%x\n\n%x\n\n%x\n",
		hash.Uint8String(input),
		hash.Int8String(input),
		hash.UintString(input),
		hash.Uint256String(input),
		hash.Int256String(input),
	)
}
```

## Support

If you use this package, or find any value in it, please consider donating:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)