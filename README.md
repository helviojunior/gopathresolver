# GO Path Resolver

This package provides support for resolve full path (absolute path) according several conditions.

## Get

```bash
go get github.com/helviojunior/gopathresolver@v0
```

## Usage


```golang

import (
    "fmt"

    resolver "github.com/helviojunior/gopathresolver"
)

file := "~/teste.md"
if resolvedPath, err := resolver.ResolveFullPath(file); err != nil {
    panic(err)
}

fmt.Println("Resolved path: ", resolvedPath)
```

## Test

```golang
go test -v .
```