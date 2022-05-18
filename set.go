package go_common

// Set /**
type Set struct {
	list map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		list: map[interface{}]bool{},
	}
}

func (s *Set) IsEmpty() bool {
	return len(s.list) <= 0
}

// Size return the number of the items in the set
func (s *Set) Size() int {
	return len(s.list)
}

// Clear Removes all items from the set
func (s *Set) Clear() {
	s.list = map[interface{}]bool{}
}

// Add : add the specific item to the set
// Return true if the set contains the elements to be removed
func (s *Set) Add(item interface{}) bool {
	if _, ok := s.list[item]; !ok {
		s.list[item] = true
		return true
	} else {
		return false
	}
}

// Remove : remove the specific item from the set if contains it.
// Returns true if the set contains the elements to be removed
func (s *Set) Remove(item interface{}) bool {
	if _, ok := s.list[item]; ok {
		delete(s.list, item)
		return true
	} else {
		return false
	}
}

func (s *Set) Contains(item interface{}) bool {
	if _, ok := s.list[item]; ok {
		return true
	}

	return false
}
