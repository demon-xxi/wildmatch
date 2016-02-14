package wildmatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples

func ExampleIsSubset_positive() {
	fmt.Println(Wildcard("a/x?/cd").IsSubset("a/*/c?"))
	// Output: a/x?/cd is a subset of the wildcard a/*/c?
}

func ExampleIsSubset_negative() {
	fmt.Println(Wildcard("a/*/c").IsSubset("a/?/c"))
	// Output: a/*/c is not a subset of a/?/c
}

// Tests

func TestAlwaysPass(t *testing.T) {
	assert.True(t, true)
}

func TestIdenticWildcards(t *testing.T) {
	assert.True(t, Wildcard("a").IsSubset("a"))
	assert.True(t, Wildcard("a/b").IsSubset("a/b"))
	assert.True(t, Wildcard("").IsSubset(""))
	assert.True(t, Wildcard("*").IsSubset("*"))
	assert.True(t, Wildcard("?").IsSubset("?"))
}

func TestSimpleCases(t *testing.T) {
	assert.True(t, Wildcard("a.txt").IsSubset("a.txt"))
	assert.False(t, Wildcard("a.txt").IsSubset("a.doc"))
	assert.False(t, Wildcard("a").IsSubset("b"))
	assert.True(t, Wildcard("").IsSubset("*"))
	assert.True(t, Wildcard("?").IsSubset("*"))
	assert.True(t, Wildcard("a").IsSubset("?"))
	assert.True(t, Wildcard("??").IsSubset("*"))
	assert.True(t, Wildcard("?a?").IsSubset("*"))
	assert.True(t, Wildcard("?*?").IsSubset("*"))
	assert.True(t, Wildcard("*?").IsSubset("*"))
	assert.True(t, Wildcard("?*").IsSubset("*"))
	assert.True(t, Wildcard("*a").IsSubset("*"))
	assert.True(t, Wildcard("*").IsSubset("**"))
	assert.False(t, Wildcard("*").IsSubset("?"))
	assert.False(t, Wildcard("*").IsSubset("a"))
	assert.False(t, Wildcard("?").IsSubset("a"))
	assert.False(t, Wildcard("?").IsSubset("??"))
}

func TestCaseSensitive(t *testing.T) {
	assert.True(t, Wildcard("Ab").IsSubset("Ab"))
	assert.True(t, Wildcard("Ab/c").IsSubset("Ab/c"))
	assert.False(t, Wildcard("Ab").IsSubset("ab"))
	assert.False(t, Wildcard("A/b").IsSubset("a/b"))
}

func TestNested(t *testing.T) {
	assert.True(t, Wildcard("a/b").IsSubset("a/?"))
	assert.True(t, Wildcard("a/b").IsSubset("a/*"))
	assert.True(t, Wildcard("a/b").IsSubset("*/?"))
	assert.True(t, Wildcard("aa/b").IsSubset("*/?"))
	assert.False(t, Wildcard("a/bb").IsSubset("*/?"))
	assert.True(t, Wildcard("aa/b").IsSubset("**/*"))
	assert.True(t, Wildcard("a/b/c").IsSubset("*/**b**/*"))
	assert.True(t, Wildcard("aa/b/c/d").IsSubset("**/*"))
	assert.False(t, Wildcard("a/bb").IsSubset("*"))
	assert.False(t, Wildcard("a/b").IsSubset("**"))
	assert.True(t, Wildcard("a/b/").IsSubset("**/"))
	assert.True(t, Wildcard("a/b/c/d").IsSubset("a/**/d"))
}
