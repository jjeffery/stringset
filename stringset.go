// Copyright 2016 John Jeffery <john@jeffery.id.au>. All rights reserved.
// License: MIT. See Readme.md.

// Package stringset implements a set of strings.
package stringset

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Set is a set of strings.
type Set map[string]struct{}

// New creates a new string set populated with the values in v.
func New(v ...string) Set {
	return Add(nil, v...)
}

// Add adds the values in v to the set and returns the set.
// If the set is nil, a new set is created.
func Add(set Set, v ...string) Set {
	if set == nil {
		set = make(Set)
	}
	for _, s := range v {
		set[s] = struct{}{}
	}
	return set
}

// Add adds the values in v to the set. This function panics if set is nil.
// Use the stringset.Add function if there is a possibility of
// set being nil.
// Add returns set in order to support method chaining.
func (set Set) Add(v ...string) Set {
	if set == nil {
		panic("nil stringset")
	}
	for _, s := range v {
		set[s] = struct{}{}
	}
	return set
}

// Remove removes the values in v from the set.
// Remove returns set in order to support method chaining.
func (set Set) Remove(v ...string) Set {
	if set != nil {
		for _, s := range v {
			delete(set, s)
		}
	}
	return set
}

// Len returns the number of item in the set.
func (set Set) Len() int {
	return len(set)
}

// Contains returns true if the string set contains s.
func (set Set) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// Equal returns true if set is equal to other.
func (set Set) Equal(other Set) bool {
	if len(set) != len(other) {
		return false
	}
	for s := range set {
		if _, ok := other[s]; !ok {
			return false
		}
	}
	return true
}

// Values returns the values of the set as a slice of string.
// If the set is empty, returns nil.
func (set Set) Values() []string {
	if len(set) == 0 {
		return nil
	}
	values := make([]string, 0, len(set))
	for s := range set {
		values = append(values, s)
	}
	if len(values) > 1 {
		// Sort the strings for consistent output.
		sort.Strings(values)
	}
	return values
}

// Join concatenates the sorted elements of set to create a single string.
// The separator string sep is placed between elements in the resulting
// string.
func (set Set) Join(sep string) string {
	values := set.Values()
	return strings.Join(values, sep)
}

// MarshalJSON implements the json.Marshaler interface.
func (set Set) MarshalJSON() ([]byte, error) {
	var strs []string
	if set != nil {
		strs = set.Values()
	}
	return json.Marshal(strs)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (set *Set) UnmarshalJSON(data []byte) error {
	var strs []string
	if err := json.Unmarshal(data, &strs); err != nil {
		return err
	}
	if strs == nil {
		*set = nil
	} else {
		*set = New(strs...)
	}
	return nil
}

// GoString implements the GoStringer inteface.
func (set Set) GoString() string {
	if set == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", set.Values())
}

// Format implements the Formatter interface
func (set Set) Format(f fmt.State, c rune) {
	str := set.GoString()
	f.Write([]byte(str))
}
