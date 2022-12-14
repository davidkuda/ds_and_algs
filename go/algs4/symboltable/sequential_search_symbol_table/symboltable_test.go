package sequential_search_symbol_table

import (
	"testing"
)

func TestSearchFound(t *testing.T) {
	st := *prepareTestST()
	found := st.search("third")
	if found == nil {
		t.Errorf("Key third should be present")
	}
}

func TestSearchNotFound(t *testing.T) {
	st := *prepareTestST()
	found := st.search("inexistent")
	if found != nil {
		t.Errorf("An unexpected key was found")
	}
}

func prepareTestST() *SequentialSearchSymbolTable {
	n1 := node{"tail", 10, nil}
	n2 := node{"second", 20, &n1}
	n3 := node{"third", 30, &n2}
	n4 := node{"head", 40, &n3}

	st := SequentialSearchSymbolTable{&n4, &n1}
	return &st
}
