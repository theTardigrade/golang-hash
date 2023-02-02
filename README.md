# golang-hash

This package contains functions that produce non-cryptographic hashes in various bit sizes. These functions use the Fowler–Noll–Vo hash (FNV-1a) algorithm.

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

If you use this package, or find any value in it, please consider donating at [Ko-fi](https://ko-fi.com/thetardigrade).