package bst

import (
	"testing"
)

func TestPut(t *testing.T) {
	newNode := node{key: "root", value: 42, count: 1}
	bst := binarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
}

func TestCount(t *testing.T) {
	bst := prepareBst()
	actual := bst.root.count
	expected := 7
	if actual != expected {
		t.Errorf("Wrong count, expected %d, got %d", expected, actual)
	}
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

func TestKeys(t *testing.T) {
	var sum int
	bst := prepareBst()
	keysQueue := bst.Keys()
	for {
		key, ok := keysQueue.Dequeue()
		if !ok {
			break
		}
		v, _ := bst.Get(key)
		sum += v
	}
	expected := 42 + 6*108
	if sum != expected {
		t.Errorf("Did you get the right keys? Expected: %d, Actual %d", expected, sum)
	}
}

func prepareBst() binarySearchTree {
	newNode := node{key: "root", value: 42, count: 1}
	bst := binarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
	bst.Put("sierra", 108)
	bst.Put("romeo", 108)
	bst.Put("zulu", 108)
	return bst
}
