package main

import (
	"testing"
)

func TestPut(t *testing.T) {
	newNode := node{"root", 42, nil, nil}
	bst := binarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
	bst.Put("sierra", 108)
	bst.Put("romeo", 108)
}

func TestMin(t *testing.T) {
	bst := prepareBst()
	min := bst.Min()
	t.Log(min)
	if min != "bravo" {
		t.Errorf("Incorrect min! Expected: \"alpha\"; Received: \"%s\"", min)
	}
}

func TestMax(t *testing.T) {
	bst := prepareBst()
	max := bst.Max()
	if max != "zulu" {
		t.Errorf("Incorrect max! Expected: \"zulu\"; Received: \"%s\"", max)
	}
}

func prepareBst() binarySearchTree {
	newNode := node{"root", 42, nil, nil}
	bst := binarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
	bst.Put("sierra", 108)
	bst.Put("romeo", 108)
	bst.Put("zulu", 108)
	return bst
}