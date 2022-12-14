package main

type binarySearchTree struct {
	root *node
}

type node struct {
	key   string
	value int
	left  *node
	right *node
}

// return value and true if key in bst, else return 0 and false
func (bst *binarySearchTree) get(key string) (int, bool) {
	var cur *node = bst.root
	for cur != nil {
		if cur.key == key {
			return cur.value, true
		} else if key < cur.key {
			cur = cur.left
		} else { // key > cur.key
			cur = cur.right
		}
	}
	return 0, false
}
