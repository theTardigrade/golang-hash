# golang-hash

This package contains functions that produce non-cryptographic hashes.

## Example

```golang
package main

import (
	hash "github.com/theTardigrade/golang-hash"
)

func main() {
	fmt.Println(hash.UintString("this is a test"))
}
```