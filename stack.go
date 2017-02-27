package stack

import "sync"

// Stack - A stack implementation made up of Items
type Stack struct {
	head   *item
	length int
	lock   *sync.Mutex // Mutex so it's thread-safe
}

// item - a Stack is made up of Items
type item struct {
	value interface{}
	prev  *item // pointer to the previous item in the Stack
}

// New creates a new Stack (a pointer to one)
// head, length, and lock will be initialized here
func New() *Stack {
	stack := new(Stack)
	stack.head = nil
	stack.length = 0
	stack.lock = &sync.Mutex{}

	return stack
}

// Peek returns the value at the head of the Stack
// returns type interface{} since the Stack's item type may vary
func (s *Stack) Peek() interface{} {
	// Lock mutex
	s.lock.Lock()

	// check if Stack is empty
	if s.isEmpty() {
		return nil
	}

	s.lock.Unlock()

	// Otherwise return the head element (not an item, but interface{})
	return s.head.value
}

// Push pushes a new item to the top of the Stack
func (s *Stack) Push(newVal interface{}) {
	// First lock the mutex
	s.lock.Lock()

	// Create the new item
	newItem := &item{newVal, s.head}

	// Assign newItem as the new head of our Stack
	// and increment the length by 1
	s.head = newItem
	s.length++

	s.lock.Unlock()
}

// Pop removes the head item from the stack and returns it
func (s *Stack) Pop() interface{} {
	// First lock the Mutex
	s.lock.Lock()

	// check length of the Stack
	if s.isEmpty() {
		return nil
	}

	// Otherwise we remove the head item and return it
	// as well as setting the new head item properly (if applicable)
	pop := s.head.value
	s.head = s.head.prev
	s.length--

	s.lock.Unlock()

	return pop
}

// isEmpty returns whether or not the Stack is empty
// Unexported because we use this after locking the mutex
func (s *Stack) isEmpty() bool {
	return s.length == 0
}
