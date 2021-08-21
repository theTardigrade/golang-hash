# golang-hash

This package contains functions that produce non-cryptographic hashes.

## Example

```golang
package main

import (
	hash "github.com/theTardigrade/golang-hash"
)

func main() {
	const input = "this is a simple test"

	fmt.Printf(
		"%d\n\n%d\n\n%d\n\n%s\n\n%s\n",
		hash.Uint8String(input),
		hash.Int8String(input),
		hash.UintString(input),
		hash.Uint256String(input),
		hash.Int256String(input),
	)
)
```