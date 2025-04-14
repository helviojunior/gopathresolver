# GOPATHRESOLVER


This package provides support for resolve full path (absolute path) according several conditions.

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

