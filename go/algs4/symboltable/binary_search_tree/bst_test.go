package bst

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

func TestFloor(t *testing.T) {
	bst := prepareBst()
	k1, _ := bst.Floor("charlie")
	if k1 != "charlie" {
		t.Errorf("Floor was not succesful. Expected: charlie; Received: %s", k1)
	}

	k2, _ := bst.Floor("delta")
	if k2 != "charlie" {
		t.Errorf("Floor was not succesful. Expected: charlie; Received: %s", k2)
	}

	k3, _ := bst.Floor("tango")
	if k3 != "sierra" {
		t.Errorf("Floor was not succesful. Expected: charlie; Received: %s", k3)
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