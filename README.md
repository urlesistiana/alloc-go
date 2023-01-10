# alloc-go

Simple `[]byte` buffer pool for go backend by `sync.Pool`. At most 50% memory waste for small buffers (`len(b) <= 65536`) and 20% for large buffers.

[![Go Reference](https://pkg.go.dev/badge/github.com/urlesistiana/alloc-go.svg)](https://pkg.go.dev/github.com/urlesistiana/alloc-go)

```go
package main

import "github.com/urlesistiana/alloc-go"

func main() {
	b := alloc.Get(1024)
	alloc.Release(b)
}
```


