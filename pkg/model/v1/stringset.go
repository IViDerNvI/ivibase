package v1

type StringSet map[string]interface{}

// New creates a new set.

func New() StringSet {
	return make(StringSet)
}

// Add adds a string to the set.
func (s StringSet) Add(str string) {
	s[str] = interface{}(nil)
}

// Remove removes a string from the set.
func (s StringSet) Remove(str string) {
	delete(s, str)
}

// Contains checks if a string is in the set.
func (s StringSet) Contains(str string) bool {
	_, ok := s[str]
	return ok
}

// Len returns the length of the set.
func (s StringSet) Len() int {
	return len(s)
}

// List returns the list of strings in the set.
func (s StringSet) List() []string {
	list := make([]string, 0, len(s))
	for str := range s {
		list = append(list, str)
	}
	return list
}

// Merge merges another set into the set.
func (s StringSet) Merge(other StringSet) {
	for str := range other {
		s[str] = interface{}(nil)
	}
}

// Clone clones the set.
func (s StringSet) Clone() StringSet {
	clone := make(StringSet)
	for str := range s {
		clone[str] = interface{}(nil)
	}
	return clone
}

// Equal checks if two sets are equal.
func (s StringSet) Equal(other StringSet) bool {
	if len(s) != len(other) {
		return false
	}
	for str := range s {
		if !other.Contains(str) {
			return false
		}
	}
	return true
}

// Union returns the union of two sets.
func (s StringSet) Union(other StringSet) StringSet {
	union := s.Clone()
	union.Merge(other)
	return union
}

// Intersection returns the intersection of two sets.
func (s StringSet) Intersection(other StringSet) StringSet {
	intersection := make(StringSet)
	for str := range s {
		if other.Contains(str) {
			intersection.Add(str)
		}
	}
	return intersection
}

// Difference returns the difference of two sets.
func (s StringSet) Difference(other StringSet) StringSet {
	difference := make(StringSet)
	for str := range s {
		if !other.Contains(str) {
			difference.Add(str)
		}
	}
	return difference
}

// SymmetricDifference returns the symmetric difference of two sets.
func (s StringSet) SymmetricDifference(other StringSet) StringSet {
	return s.Difference(other).Union(other.Difference(s))
}

// IsSubset checks if the set is a subset of another set.
func (s StringSet) IsSubset(other StringSet) bool {
	for str := range s {
		if !other.Contains(str) {
			return false
		}
	}
	return true
}

// IsSuperset checks if the set is a superset of another set.
func (s StringSet) IsSuperset(other StringSet) bool {
	return other.IsSubset(s)
}

// IsDisjoint checks if the set is disjoint with another set.
func (s StringSet) IsDisjoint(other StringSet) bool {
	for str := range s {
		if other.Contains(str) {
			return false
		}
	}
	return true
}

// IsEqual checks if the set is equal to another set.
func (s StringSet) IsEqual(other StringSet) bool {
	return s.Equal(other)
}

// IsIntersection checks if the set is an intersection of two sets.
func (s StringSet) IsIntersection(other StringSet) bool {
	return s.Intersection(other).IsEqual(s)
}

// IsDifference checks if the set is a difference of two sets.
func (s StringSet) IsDifference(other StringSet) bool {
	return s.Difference(other).IsEqual(s)
}

// IsSymmetricDifference checks if the set is a symmetric difference of two sets.
func (s StringSet) IsSymmetricDifference(other StringSet) bool {
	return s.SymmetricDifference(other).IsEqual(s)
}

// IsUnion checks if the set is a union of two sets.
func (s StringSet) IsUnion(other StringSet) bool {
	return s.Union(other).IsEqual(s)
}

// IsSubsetOf checks if the set is a subset of another set.
func (s StringSet) IsSubsetOf(other StringSet) bool {
	return s.IsSubset(other)
}

// IsSupersetOf checks if the set is a superset of another set.
func (s StringSet) IsSupersetOf(other StringSet) bool {
	return s.IsSuperset(other)
}

// IsDisjointWith checks if the set is disjoint with another set.
func (s StringSet) IsDisjointWith(other StringSet) bool {
	return s.IsDisjoint(other)
}
