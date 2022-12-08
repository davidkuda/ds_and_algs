package main

import "fmt"

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
