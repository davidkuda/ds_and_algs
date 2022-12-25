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

// search for key, update if there, add new node if not
// Tree shape depends on order of insertion. Worst case: No difference to linked list.
func (bst *binarySearchTree) Put(key string, value int) {
	bst.root = bst.put(bst.root, key, value)
}

func (bst *binarySearchTree) put(n *node, key string, val int) *node {
	if n == nil {
		newNode := node{key: key, value: val}
		return &newNode
	}
	if key < n.key {
		n.left = bst.put(n.left, key, val)
	} else if key > n.key {
		n.right = bst.put(n.right, key, val)
	} else { // key == n.key
		n.value = val
		return n
	}
	return n
}

// Smallest key in symbol table: Move to the left most node
func (bst *binarySearchTree) Min() string {
	n := bst.root
	var previous *node
	for n != nil {
		previous = n
		n = n.left
	}
	return previous.key
}

// Largest key in symbol table: Move to the right most node
func (bst *binarySearchTree) Max() string {
	n := bst.root
	var previous *node
	for n != nil {
		previous = n
		n = n.right
	}
	return previous.key
}

// Largest key <= to a given key
// Three cases:
// 1) k == node -> floor is k
// 2) k < node -> floor is in the left subtree
// 3) k > node -> if k < node in right subtree, right subtree, else node
func (bst *binarySearchTree) Floor(key string) (string, bool) {
	x := bst.floor(bst.root, key)
	if x == nil {
		return "", false
	}
	return x.key, true
}

func (bst *binarySearchTree) floor(x *node, key string) *node {
	if x == nil {
		return nil
	}

	if x.key == key {
		return x
	} else if key < x.key {
		return bst.floor(x.left, key)
	}
	t := bst.floor(x.right, key)
	if t != nil {
		return t
	} else {
		return x
	}
}

// Smallest key >= to a given key
func (bst *binarySearchTree) Ceiling(key string) string {
	return ""
}