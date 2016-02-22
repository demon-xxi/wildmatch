package wildmatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Examples

func ExampleIsSubsetOf_positive() {
	fmt.Println(Wildcard("a/x?/cd").IsSubsetOf("a/*/c?"))
	// Output: true
}

func ExampleIsSubsetOf_negative() {
	fmt.Println(Wildcard("a/*/c").IsSubsetOf("a/?/c"))
	// Output: false
}

func ExampleIsSubsetOfAny() {
	ans, _ := Wildcard("a*.txt").IsSubsetOfAny(Wildcard("*"), Wildcard("*.txt"),
		Wildcard("*.t?t"), Wildcard("*.?x?"))
	fmt.Println(ans)
	// Output: *.txt
}

//Shim for 2 param return values
func M(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

//Shim for 3 param return values
func M3(a, b, c interface{}) []interface{} {
	return []interface{}{a, b, c}
}

//Shim for 4 param return values
func M4(a, b, c, d interface{}) []interface{} {
	return []interface{}{a, b, c, d}
}

// Tests

func TestAlwaysPass(t *testing.T) {
	assert.True(t, true)
}

func TestIdenticWildcards(t *testing.T) {
	assert := assert.New(t)

	assert.True(Wildcard("a").IsSubsetOf("a"))
	assert.True(Wildcard("a/b").IsSubsetOf("a/b"))
	assert.True(Wildcard("").IsSubsetOf(""))
	assert.True(Wildcard("*").IsSubsetOf("*"))
	assert.True(Wildcard("?").IsSubsetOf("?"))
}

func TestSimpleCases(t *testing.T) {
	assert := assert.New(t)

	assert.True(Wildcard("a.txt").IsSubsetOf("a.txt"))
	assert.False(Wildcard("a.txt").IsSubsetOf("a.doc"))
	assert.False(Wildcard("a").IsSubsetOf("b"))
	assert.True(Wildcard("").IsSubsetOf("*"))
	assert.True(Wildcard("?").IsSubsetOf("*"))
	assert.True(Wildcard("a").IsSubsetOf("?"))
	assert.True(Wildcard("??").IsSubsetOf("*"))
	assert.True(Wildcard("?a?").IsSubsetOf("*"))
	assert.True(Wildcard("?*?").IsSubsetOf("*"))
	assert.True(Wildcard("*?").IsSubsetOf("*"))
	assert.True(Wildcard("?*").IsSubsetOf("*"))
	assert.True(Wildcard("*a").IsSubsetOf("*"))
	assert.True(Wildcard("*").IsSubsetOf("**"))
	assert.False(Wildcard("*").IsSubsetOf("?"))
	assert.False(Wildcard("*").IsSubsetOf("a"))
	assert.False(Wildcard("?").IsSubsetOf("a"))
	assert.False(Wildcard("?").IsSubsetOf("??"))
}

func TestCaseSensitive(t *testing.T) {
	assert := assert.New(t)

	assert.True(Wildcard("Ab").IsSubsetOf("Ab"))
	assert.True(Wildcard("Ab/c").IsSubsetOf("Ab/c"))
	assert.False(Wildcard("Ab").IsSubsetOf("ab"))
	assert.False(Wildcard("A/b").IsSubsetOf("a/b"))
}

func TestNested(t *testing.T) {
	assert := assert.New(t)

	assert.False(Wildcard("a").IsSubsetOf("a/b"))
	assert.True(Wildcard("a/b").IsSubsetOf("a/?"))
	assert.True(Wildcard("a/b").IsSubsetOf("a/*"))
	assert.True(Wildcard("a/b").IsSubsetOf("*/?"))
	assert.True(Wildcard("aa/b").IsSubsetOf("*/?"))
	assert.False(Wildcard("a/bb").IsSubsetOf("*/?"))
	assert.True(Wildcard("aa/b").IsSubsetOf("**/*"))
	assert.True(Wildcard("a/b/c").IsSubsetOf("*/**b**/*"))
	assert.True(Wildcard("aa/b/c/d").IsSubsetOf("**/*"))
	assert.False(Wildcard("a/bb").IsSubsetOf("*"))
	assert.False(Wildcard("a/b").IsSubsetOf("**"))
	assert.True(Wildcard("a/b/").IsSubsetOf("**/"))
	assert.True(Wildcard("a/b").IsSubsetOf("a/**/b"))
	assert.True(Wildcard("a/b/c/d").IsSubsetOf("a/**/d"))
}

func TestIsSubsetOfAny(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(M(Wildcard("*"), true), M(Wildcard("a").IsSubsetOfAny("*", "*.txt", "**/*")))
	assert.Equal(M(Wildcard(""), false), M(Wildcard("a").IsSubsetOfAny("*.txt", "**/*")))
	assert.Equal(M(Wildcard("a/b/*"), true), M(Wildcard("a/b/c").IsSubsetOfAny("*", "**/*", "a/b/*", "a/*", "**/*")))
}
