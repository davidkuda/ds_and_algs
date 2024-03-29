package bst

// methods on the bst will be public
type BinarySearchTree struct {
	root *treeNode
}

// methods on the treeNode will be private
type treeNode struct {
	key   string
	value int
	left  *treeNode
	right *treeNode
	count int
}

func (bst *BinarySearchTree) Size() int {
	return bst.root.size()
}

func (n *treeNode) size() int {
	if n == nil {
		return 0
	} else {
		return n.count
	}
}

// return value and true if key in bst, else return 0 and false
func (bst *BinarySearchTree) Get(key string) (int, bool) {
	n, found := bst.root.get(key)
	if !found {
		return 0, false
	} else {
		return n.value, true
	}
}

// get a node; true if found, false if not
func (n *treeNode) get(key string) (*treeNode, bool) {
	for n != nil {
		if n.key == key {
			return n, true
		} else if key < n.key {
			n = n.left
		} else { // key > n.key
			n = n.right
		}
	}
	return nil, false
}

// search for key, update if there, add new node if not
// Tree shape depends on order of insertion. Worst case: No difference to linked list.
func (bst *BinarySearchTree) Put(key string, value int) {
	bst.root = bst.root.put(key, value)
}

func (n *treeNode) put(key string, val int) *treeNode {
	if n == nil {
		newNode := treeNode{key: key, value: val, count: 1}
		return &newNode
	}
	if key < n.key {
		n.left = n.left.put(key, val)
	} else if key > n.key {
		n.right = n.right.put(key, val)
	} else { // key == n.key
		n.value = val
		return n
	}
	n.count = 1 + n.left.size() + n.right.size()
	return n
}

/*
Hibbard Deletion:
Case 0:

	    Node to be deleted has no children;
		Set parent link to nil;
		update counts

Case 1:

	    Node has 1 child;
		Replace parent link with child link (like DeleteMin / DeleteMax);
		Update count.

Case 2:

	    Node has two children;
		Find min in right tree, delete it, and replace current link with min of right tree;
	    update count.

Problem of Hibbard Deletion:

	Although randomised bst.Add() will yield a fairly balanced tree,
	the Hibbard deletion will cause an asymmetry, since we replace with
	right successor, it tends to skew to the right and be unbalanced
	over time.

	It's still a problem to find a natural and simple bst.Delete(key)
*/
func (bst *BinarySearchTree) Delete(key string) {
	bst.root = bst.root.delete(key)
}

func (n *treeNode) delete(key string) *treeNode {
	if n == nil {
		return nil
	}

	// search for key:
	if n.key < key {
		n.right = n.right.delete(key)
	} else if n.key > key {
		n.left = n.left.delete(key)
	} else { // n.key == key
		// no right child
		if n.right == nil {
			return n.left
		}
		// no left child
		if n.left == nil {
			return n.right
		}

		// replace with successor
		var rightMin treeNode = *n.right.min() // copy value of ptr
		n.right.deleteMin()
		rightMin.left = n.left
		rightMin.right = n.right
		n = &rightMin
	}
	// update count
	n.count = 1 + n.left.size() + n.right.size()
	return n
}

func (bst *BinarySearchTree) DeleteMin() {
	bst.root = bst.root.deleteMin()
}

// go left until finding a node with a nil left link
// replace that node by its right link
// update subtree counts
func (n *treeNode) deleteMin() *treeNode {
	if n.left == nil {
		return n.right
	}
	n.left = n.left.deleteMin()
	n.count = 1 + n.left.size() + n.right.size()
	return n
}

func (bst *BinarySearchTree) DeleteMax() {
	bst.root = bst.root.deleteMax()
}

func (n *treeNode) deleteMax() *treeNode {
	if n.right == nil {
		return n.left
	}
	n.right = n.right.deleteMax()
	n.count = 1 + n.left.size() + n.right.size()
	return n
}

// returns the number of nodes with smaller keys than the given key
// [1,2,3,4,5,6,7] -> rank(3) -> returns 2, 2 smaller nodes than 3
func (bst *BinarySearchTree) Rank(key string) int {
	return bst.root.rank(key)
}

func (n *treeNode) rank(key string) int {
	if n == nil {
		return 0
	}
	if key < n.key {
		return n.left.rank(key)
	} else if key > n.key {
		return 1 + n.left.size() + n.right.rank(key)
	} else { // key == n.key
		return n.left.size()
	}
}

func (bst *BinarySearchTree) Keys() *Queue[string] {
	q := NewQueue[string]()
	bst.root.inorder(q)
	return q
}

func (n *treeNode) inorder(q *Queue[string]) {
	if n == nil {
		return
	}
	n.left.inorder(q)
	q.Enqueue(n.key)
	n.right.inorder(q)
}

// Smallest key in symbol table: Move to the left most node
func (bst *BinarySearchTree) Min() string {
	return bst.root.min().key
}

func (n *treeNode) min() *treeNode {
	var previous *treeNode
	for n != nil {
		previous = n
		n = n.left
	}
	return previous
}

// Largest key in symbol table: Move to the right most node
func (bst *BinarySearchTree) Max() string {
	return bst.root.max().key
}

func (n *treeNode) max() *treeNode {
	var previous *treeNode
	for n != nil {
		previous = n
		n = n.right
	}
	return previous
}

// Largest key <= to a given key
// Three cases:
// 1) k == node -> floor is k
// 2) k < node -> floor is in the left subtree
// 3) k > node -> if k < node in right subtree, right subtree, else node
func (bst *BinarySearchTree) Floor(key string) (string, bool) {
	x := floor(bst.root, key)
	if x == nil {
		return "", false
	}
	return x.key, true
}

func floor(x *treeNode, key string) *treeNode {
	if x == nil {
		return nil
	}

	if x.key == key {
		return x
	} else if key < x.key {
		return floor(x.left, key)
	}
	t := floor(x.right, key)
	if t != nil {
		return t
	} else {
		return x
	}
}

// Smallest key >= to a given key
func (bst *BinarySearchTree) Ceiling(key string) string {
	return ""
}
