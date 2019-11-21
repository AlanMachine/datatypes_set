// Implementation of set
package set

import "fmt"

type Set map[string]struct{}

func (s Set) Add(item string) {
	if !s.Contains(item) {
		s[item] = struct{}{}
	}
}

func (s Set) AddRange(items []string) {
	for _, item := range items {
		if !s.Contains(item) {
			s.Add(item)
		}
	}
}

func (s Set) Remove(item string) bool {
	if !s.Contains(item) {
		return false
	}
	delete(s, item)

	return true
}

func (s Set) Contains(item string) bool {
	if _, ok := s[item]; !ok {
		return false
	}

	return true
}

// Return the number of elements in set
func (s Set) Count() int {
	return len(s)
}

// Return a new set with elements from the set and all others
func (s Set) Union(other Set) Set {
	result := Set{}
	result.AddRange(s.toSlice())
	for item := range other {
		if !s.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Return a new set with elements common to the set and all others
func (s Set) Intersection(other Set) Set {
	result := Set{}
	for item := range other {
		if s.Contains(item) {
			result.Add(item)
		}
	}

	return result
}

// Return a new set with elements in the set that are not in the others
func (s Set) Difference(other Set) Set {
	result := Set{}
	result.AddRange(s.toSlice())
	for item := range other {
		result.Remove(item)
	}

	return result
}

// Return a new set with elements in either the set or other but not both
func (s Set) SymmetricDifference(other Set) Set {
	result := Set{}
	result.AddRange(s.toSlice())

	return result.Union(other).Difference(result.Intersection(other))
}

// Test whether every element in the set is in other
func (s Set) IsSubset(other Set) bool {
	if len(s.Difference(other)) != 0 {
		return false
	}

	return true
}

func (s Set) String() string {
	return fmt.Sprintf("%v", s.toSlice())
}

func (s Set) toSlice() []string {
	var output []string
	for item := range s {
		output = append(output, item)
	}

	return output
}
