// Package wildmatch solves inclusion problem for the wildcard path patterns.
package wildmatch

import (
	"strings"
)

// Wildcard respresents wildcard pattern.
// It is an alias to the string type to provide extra methods of this package.
type Wildcard string

// IsSubset verifies if current wildcard is a subset of a given one `superset`.
// Wildcard A is subset of B if any possible path that matches A also matches B.
func (w Wildcard) IsSubsetOf(superset Wildcard) bool {
	return IsSubsetOf(string(w), string(superset))
}

// Contains verifies if current wildcard is a superset of a given one `subset`.
// This operation is inversion of IsSubsetOf().
func (w Wildcard) Contains(subset Wildcard) bool {
	return IsSubsetOf(string(subset), string(w))
}

// IsSubsetOf verifies if `w` wildcard is a subset of `s`.
// I.e. checks if `s` is a superset of subset `w`.
func IsSubsetOf(w string, s string) bool {

	// shortcut for identical sets
	if s == w {
		return true
	}

	// only empty set is a subset of an empty set
	if len(s) == 0 {
		return len(w) == 0
	}

	// find nesting separators
	sp := strings.Index(s, "/")
	wp := strings.Index(w, "/")

	// check if this is a nested path
	if sp >= 0 {

		// if set is nested then tested wildcard must be nested too
		if wp < 0 {
			return false
		}

		// Special case for /**/ mask that matches any number of levels
		if s[:sp] == "**" &&
			IsSubsetOf(w[wp+1:], s) ||
			IsSubsetOf(w, s[sp+1:]) {
			return true
		}

		// check that current level names are subsets
		// and compare rest of the path to be subset also
		return (IsSubsetOf(w[:wp], s[:sp]) &&
			IsSubsetOf(w[wp+1:], s[sp+1:]))
	}

	// subset can't have more levels than set
	if wp >= 0 {
		return false
	}

	// we are comparing names on the same nesting level here
	// so let's do symbol by symbol comparison
	switch s[0] {
	case '?':
		// ? matches non empty character. '*' can't be a subset of '?'
		if len(w) == 0 || w[0] == '*' {
			return false
		}
		// any onther symbol matches '?', so let's skip to next
		return IsSubsetOf(w[1:], s[1:])
	case '*':
		// '*' matches 0 and any other number of symbols
		// so checking 0 and recursively subset without first letter
		return IsSubsetOf(w, s[1:]) ||
			(len(w) > 0 && IsSubsetOf(w[1:], s))
	default:
		// making sure next symbol in w exists and it's the same as in set
		if len(w) == 0 || w[0] != s[0] {
			return false
		}
	}

	// recursively check rest of the set and w
	return IsSubsetOf(w[1:], s[1:])
}

// IsSubsetOfAny verifies if current wildcard is a subset of any of the given sets.
// Wildcard A is subset of B if any possible path that matches A also matches B.
// If multiple subsets match then the smallest or first lexicographical set is returned
func (w Wildcard) IsSubsetOfAny(supersets ...Wildcard) (result Wildcard, found bool) {
	for _, superset := range supersets {
		if !w.IsSubsetOf(superset) {
			continue
		}
		found = true
		if result == "" || superset.IsSubsetOf(result) {
			result = superset
		}
	}
	return
}

// IsSubsetOfAny verifies if current wildcard `w` is a subset of any of the given sets.
// Wildcard A is subset of B if any possible path that matches A also matches B.
// If multiple subsets match then the smallest or first lexicographical set is returned
func IsSubsetOfAny(w string, sets ...string) (result string, found bool) {
	if index := IsSubsetOfAnyI(w, sets...); index >= 0 {
		return sets[index], true
	}

	return "", false // not found
}

// IsSubsetOfAny verifies if current wildcard `w` is a subset of any of the given sets.
// Wildcard A is subset of B if any possible path that matches A also matches B.
// If multiple subsets match then the smallest or first lexicographical set is returned
// Return -1 if not found or superset index.
func IsSubsetOfAnyI(w string, sets ...string) (found int) {
	found = -1 // not found by default
	for i, superset := range sets {
		if !IsSubsetOf(w, superset) {
			continue
		}
		if found < 0 || IsSubsetOf(superset, sets[found]) {
			found = i
		}
	}
	return
}

// NOTE it's hard to combine both IsSubsetOfAny() functions to one
// because of missing []string <-> []Wildcard conversion
