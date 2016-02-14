// Package wildmatch solves inclusion problem for the wildcard path patterns.
package wildmatch

import (
	"strings"
)

// Wildcard respresents wildcard pattern.
// It is an alias to the string type to provide extra methods of this package.
type Wildcard string

// IsSubset verifies if current wildcard is a subset of a given one.
// Wildcard A is subset of B if any possible path that matches A also matches B.
func (w Wildcard) IsSubset(set Wildcard) bool {

	// shortcut for identical sets
	if set == w {
		return true
	}

	// only empty set is a subset of an empty set
	if len(string(set)) == 0 {
		return len(string(w)) == 0
	}

	// find nesting separators
	sep := strings.Index(string(set), "/")
	wsep := strings.Index(string(w), "/")

	// check if this is a nested path
	if sep >= 0 {

		// if set is nested then tested wildcard must be nested too
		if wsep < 0 {
			return false
		}

		// check that current level names are subsets
		// and copare rest of the path to be subset also
		return (Wildcard(w[:wsep]).IsSubset(set[:sep]) &&
			Wildcard(w[wsep+1:]).IsSubset(set[sep+1:])) ||
			// Special case for /**/ mask that matches any number of levels
			(set[:sep] == "**" &&
				Wildcard(w[wsep+1:]).IsSubset(set))
	}

	// subset can't have more levels than set
	if wsep >= 0 {
		return false
	}

	// we are comparing names on the same nesing level here
	// so let's do symbol by symbol comparison
	switch string(set)[0] {
	case '?':
		// ? matches non empty character. '*' can't be a subset of '?'
		if len(string(w)) == 0 || string(w)[0] == '*' {
			return false
		}
		// any onther symbol matches '?', so let's skip to next
		return Wildcard(string(w)[1:]).IsSubset(Wildcard(string(set)[1:]))
	case '*':
		// '*' matches 0 and any other number of symbols
		// so checking 0 and recursively subset without first letter
		return w.IsSubset(Wildcard(string(set)[1:])) ||
			(len(string(w)) > 0 && Wildcard(string(w)[1:]).IsSubset(set))
	default:
		// making sure next symbol in w exists and it's the same as in set
		if len(string(w)) == 0 || string(w)[0] != string(set)[0] {
			return false
		}
		// recursively check rest of the set and w
		return Wildcard(string(w)[1:]).IsSubset(Wildcard(string(set)[1:]))
	}

}
