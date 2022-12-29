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

func TestDelete(t *testing.T) {
	root := node{key: "07"}
	bst := binarySearchTree{root: &root}
	bst.Put("04", 42)
	bst.Put("05", 42)
	bst.Put("06a", 42)
	bst.Put("06", 42)
	bst.Put("06b", 42)
	bst.Put("03", 42)
	bst.Put("09", 42)
	bst.Put("08", 42)
	bst.Put("10", 42)

	if bst.root.left.right.key != "05" {
		t.Error("Something wrong with the bst")
	}

	bst.Delete("05")

	if bst.root.left.right.key != "06" {
		t.Error("Something wrong with the bst")
	}

	if bst.root.left.right.right.left.key != "06a" {
		t.Error("Something wrong with the bst!")
	}

	if bst.Size() != 9 {
		t.Errorf("Size is incorrect, expected 7, got %d", bst.Size())
	}

	val, ok := bst.Get("05")
	if ok != false {
		t.Error("Key 05 should be deleted!")
	}
	if val != 0 {
		t.Errorf("Val should be 0, got %d instead", val)
	}
}

func TestDeleteMin(t *testing.T) {
	bst := prepareAnotherBst()
	oldSize := bst.root.size()
	min := bst.Min()
	if min != "a" {
		t.Errorf("Incorrect min! Expected: \"alpha\"; Received: \"%s\"", min)
	}
	bst.DeleteMin()
	newMin := bst.Min()
	if min == newMin {
		t.Error("min was not deleted")
	}
	if newMin != "a1" {
		t.Error("expected a1, got sth else")
	}
	if oldSize != bst.root.size()+1 {
		t.Error("Wrong Size of BST!")
	}
}

func TestDeleteMax(t *testing.T) {
	bst := prepareAnotherBst()
	oldSize := bst.root.size()
	
	max := bst.Max()
	expMax := "zvb"
	if max != expMax {
		t.Errorf("Incorrect max! Expected: \"%s\"; Received: \"%s\"", expMax, max)
	}

	bst.DeleteMax()

	newMax := bst.Max()
	if max == newMax {
		t.Error("min was not deleted")
	}
	if newMax != "zva" {
		t.Error("expected zva, got sth else")
	}
	if oldSize != bst.root.size()+1 {
		t.Error("Wrong Size of BST!")
	}
}

func TestCount(t *testing.T) {
	bst := prepareBst()
	actual := bst.root.count
	expected := 7
	if actual != expected {
		t.Errorf("Wrong count, expected %d, got %d", expected, actual)
	}
}

func TestRank(t *testing.T) {
	bst := prepareBst()
	r := bst.Rank("delta")
	expected := 3
	if r != expected {
		t.Errorf("got wrong rank: expected %d, got %d", expected, r)
	}
}

func TestMin(t *testing.T) {
	bst := prepareBst()
	min := bst.Min()
	t.Log(min)
	if min != "alpha" {
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

func prepareAnotherBst() binarySearchTree {
	newNode := node{key: "root", value: 42, count: 1}
	bst := binarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
	bst.Put("sierra", 108)
	bst.Put("ac", 108)
	bst.Put("romeo", 108)
	bst.Put("zulu", 108)
	bst.Put("b", 108)
	bst.Put("a", 108)
	bst.Put("a1", 108)
	bst.Put("a2", 108)
	bst.Put("a1a", 108)
	bst.Put("zv", 108)
	bst.Put("zva", 108)
	bst.Put("zvb", 108)
	return bst
}
