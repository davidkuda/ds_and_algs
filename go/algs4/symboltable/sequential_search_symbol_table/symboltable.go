package sequential_search_symbol_table

import "fmt"

// elementary implementation:
// Simple, linear, unordered Symbol Table ("associative array" abstraction)
// API:
// - Search(key): Scan all key-value-pairs and return if key matches
// - Insert(node): search, update if key exists, place on first position if not
// time complexity: Search O(n), Insert O(n)
type SequentialSearchSymbolTable struct {
	head *node
	tail *node
}

type node struct {
	key   string
	value int
	next  *node
}

func (s *SequentialSearchSymbolTable) search(key string) *node {
	n := s.head
	var found *node
	for n != nil {
		if n.key == key {
			fmt.Println("Found Key in Symbol Table")
			found = n
			break
		} else {
			fmt.Println("Didn't find Key yet")
			n = n.next
		}
	}
	return found
}
