package wildmatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples

func ExampleIsSubsetOf_positive() {
	fmt.Println(IsSubsetOf("a/x?/cd", "a/*/c?"))
	// Output: true
}

func ExampleIsSubsetOf_negative() {
	fmt.Println(IsSubsetOf("a/*/c", "a/?/c"))
	// Output: false
}

func ExampleIsSubsetOfAny() {
	ans := IsSubsetOfAny("a*.txt", "*", "*.txt", "*.t?t", "*.?x?")
	fmt.Println(ans)
	// Output: 1
}

// Tests

func TestAlwaysPass(t *testing.T) {
	assert.True(t, true)
}

func TestIdenticWildcards(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsSubsetOf("a", "a"))
	assert.True(IsSubsetOf("a/b", "a/b"))
	assert.True(IsSubsetOf("", ""))
	assert.True(IsSubsetOf("*", "*"))
	assert.True(IsSubsetOf("?", "?"))
}

func TestSimpleCases(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsSubsetOf("a.txt", "a.txt"))
	assert.False(IsSubsetOf("a.txt", "a.doc"))
	assert.False(IsSubsetOf("a", "b"))
	assert.True(IsSubsetOf("", "*"))
	assert.True(IsSubsetOf("?", "*"))
	assert.True(IsSubsetOf("a", "?"))
	assert.True(IsSubsetOf("??", "*"))
	assert.True(IsSubsetOf("?a?", "*"))
	assert.True(IsSubsetOf("?*?", "*"))
	assert.True(IsSubsetOf("*?", "*"))
	assert.True(IsSubsetOf("?*", "*"))
	assert.True(IsSubsetOf("*a", "*"))
	assert.True(IsSubsetOf("*", "**"))
	assert.False(IsSubsetOf("*", "?"))
	assert.False(IsSubsetOf("*", "a"))
	assert.False(IsSubsetOf("?", "a"))
	assert.False(IsSubsetOf("?", "??"))
}

func TestCaseSensitive(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsSubsetOf("Ab", "Ab"))
	assert.True(IsSubsetOf("Ab/c", "Ab/c"))
	assert.False(IsSubsetOf("Ab", "ab"))
	assert.False(IsSubsetOf("A/b", "a/b"))
}

func TestNested(t *testing.T) {
	assert := assert.New(t)

	assert.False(IsSubsetOf("a", "a/b"))
	assert.True(IsSubsetOf("a/b", "a/?"))
	assert.True(IsSubsetOf("a/b", "a/*"))
	assert.True(IsSubsetOf("a/b", "*/?"))
	assert.True(IsSubsetOf("aa/b", "*/?"))
	assert.False(IsSubsetOf("a/bb", "*/?"))
	assert.True(IsSubsetOf("aa/b", "**/*"))
	assert.True(IsSubsetOf("a/b/c", "*/**b**/*"))
	assert.True(IsSubsetOf("aa/b/c/d", "**/*"))
	assert.False(IsSubsetOf("a/bb", "*"))
	assert.False(IsSubsetOf("a/b", "**"))
	assert.True(IsSubsetOf("a/b/", "**/"))
	assert.True(IsSubsetOf("a/b", "a/**/b"))
	assert.True(IsSubsetOf("a//d", "a/**/d"))
	assert.True(IsSubsetOf("a/bc/d", "a/**/d"))
	assert.True(IsSubsetOf("a/b/c/d", "a/**/d"))
	assert.True(IsSubsetOf("a/b/c/e/f/d", "a/**/d"))
}

func TestIsSubsetOfAny(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, IsSubsetOfAny("a", "*", "*.txt", "**/*"))
	assert.Equal(-1, IsSubsetOfAny("a", "*.txt", "**/*"))
	assert.Equal(2, IsSubsetOfAny("a/b/c", "*", "**/*", "a/b/*", "a/*", "**/*"))
}
