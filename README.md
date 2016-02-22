[![Build Status](https://travis-ci.org/demon-xxi/wildmatch.svg?branch=master)](https://travis-ci.org/demon-xxi/wildmatch) [![Go Report Card](https://goreportcard.com/badge/github.com/demon-xxi/wildmatch)](https://goreportcard.com/report/github.com/demon-xxi/wildmatch) [![Coverage](http://gocover.io/_badge/github.com/demon-xxi/wildmatch)](http://gocover.io/github.com/demon-xxi/wildmatch) [![GoDoc](https://godoc.org/github.com/demon-xxi/wildmatch?status.svg)](https://godoc.org/github.com/demon-xxi/wildmatch)
# wildmatch
Package wildmatch solves inclusion problem for the wildcard path patterns.

Supported patterns:

`*` - any characters. zero or more.

`?` - any character. exactly one.

`/` - folder separator.

`/**/` - any number of nested folders

For example `a/b/*` is a subset of `a/?/*` but not a subset of `a/*`.

This is because `*` includes any character withing folder/file name.

For multiple folders matching use `/**/` pattern which means any number of nested folders.

```golang
import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

// Examples

func ExampleIsSubset_positive() {
	fmt.Println(Wildcard("a/x?/cd").IsSubsetOf("a/*/c?"))
	// Output: true
}

func ExampleIsSubset_negative() {
	fmt.Println(Wildcard("a/*/c").IsSubsetOf("a/?/c"))
	// Output: false
}


```