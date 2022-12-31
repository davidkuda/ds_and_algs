package bst

import (
	"testing"
)

func TestPut(t *testing.T) {
	newNode := treeNode{key: "root", value: 42, count: 1}
	bst := BinarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
}

func TestTreeNodeGet(t *testing.T) {
	bst := prepareAnotherBst()
	v1, ok := bst.root.get("zva")
	if !ok {
		t.Error("Value was not found")
	}

	v2, ok := bst.root.get("ac")
	if !ok {
		t.Error("Value was not found")
	}

	if v1.value != v2.value {
		t.Errorf("Expected both values to be 108, got %d and %d instead", v1.value, v2.value)
	}
}

// n has left child
func TestDeleteHasLeftChild(t *testing.T) {}

// n has right child
func TestDeleteHasRightChild(t *testing.T) {}

// n has two childs
func TestDelete(t *testing.T) {
	root := treeNode{key: "07"}
	bst := BinarySearchTree{root: &root}
	bst.Put("04", 42)
	bst.Put("05", 42)
	bst.Put("04a", 42)
	bst.Put("06", 42)
	bst.Put("03", 42)
	bst.Put("09", 42)
	bst.Put("08", 42)
	bst.Put("10", 42)
	/*
			        07
				  /     \
				04      09
			  /   \    /  \
			03    05  08   10
		        /   \
			  04a    06
	*/

	if bst.root.left.key != "04" {
		t.Fatal("targeted node should have key 04")
	}

	if bst.root.left.left == nil || bst.root.left.right == nil {
		t.Fatal("targeted node should have two children")
	}

	bst.Delete("04")
	// expected:
	/*
			        07
				  /     \
				04a     09
			  /   \    /  \
			03    05  08   10
		            \
			         06
	*/

	if bst.root.left.key != "04a" {
		t.Errorf("Deletion yielded unexpected result, expected n.key{04a}, got %s", bst.root.left.key)
	}

	if bst.root.left.right.left != nil {
		t.Errorf("Deletion yileded unexpected result, expected nil, got %s", bst.root.left.right.left.key)
	}

	if bst.Size() != 8 {
		t.Errorf("Size is incorrect, expected 8, got %d", bst.Size())
	}

	val, ok := bst.Get("04")
	if ok != false {
		t.Error("Key 04 should be deleted!")
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
	bst := BinarySearchTree{}
	bst.Put("1", 1)
	bst.Put("6", 1)
	bst.Put("5", 1)
	bst.Put("3", 1)
	bst.Put("2", 1)
	bst.Put("4", 1)
	bst.Put("7", 1)
	r := bst.Rank("3")
	expected := 2
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

func prepareBst() BinarySearchTree {
	newNode := treeNode{key: "root", value: 42, count: 1}
	bst := BinarySearchTree{&newNode}
	bst.Put("alpha", 108)
	bst.Put("bravo", 108)
	bst.Put("charlie", 108)
	bst.Put("sierra", 108)
	bst.Put("romeo", 108)
	bst.Put("zulu", 108)
	return bst
}

func prepareAnotherBst() BinarySearchTree {
	newNode := treeNode{key: "root", value: 42, count: 1}
	bst := BinarySearchTree{&newNode}
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
