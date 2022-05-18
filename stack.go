package go_common

// Stack First In Last Pop
type Stack struct {
	list []interface{}
}

func NewStack() *Stack {
	return &Stack{
		list: []interface{}{},
	}
}

// IsEmpty Empty check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.list) <= 0
}

// Peak Get the top of the stack without remove the item
func (s *Stack) Peak() interface{} {
	if len(s.list) <= 0 {
		return nil
	}
	return s.list[len(s.list)-1]
}

// Push an item on the top of the stack
func (s *Stack) Push(item interface{}) bool {
	s.list = append(s.list, item)
	return true
}

// Pop Remove the item on the top of the stac
func (s *Stack) Pop() interface{} {
	if len(s.list) <= 0 {
		return nil
	}
	temp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return temp
}
