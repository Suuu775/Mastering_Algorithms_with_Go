package ch9

import (
	"cmp"
	"fmt"
	"iter"
	"strings"
)

const (
	AVL_RGT_HEAVY  = -1
	AVL_BALANCED   = 0
	AVL_LEFT_HEAVY = 1
)

type AVLNode[T cmp.Ordered] struct {
	key    T
	height int
	left   *AVLNode[T]
	right  *AVLNode[T]
}

type AVLTree[T cmp.Ordered] struct {
	size int
	root *AVLNode[T]
}

// Initialize an AVL Tree
func NewAVLTree[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		size: 0,
		root: nil,
	}
}

// Get the node height
func (node *AVLNode[T]) getHeight() int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

// Update the node height
func (node *AVLNode[T]) updateHeight() {
	if node == nil {
		return
	}
	node.height = max(node.left.getHeight(), node.right.getHeight()) + 1
}

// Compute the node balance factor
func (node *AVLNode[T]) computeBalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.left.getHeight() - node.right.getHeight()
}

// When the left child is inserted into the left tree and causes imbalance, perform a right rotation
func (node *AVLNode[T]) rotateRight() *AVLNode[T] {
	if node == nil || node.left == nil {
		return node
	}
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node

	node.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

// When the right child is inserted into the right tree, causing imbalance, perform left rotation
func (node *AVLNode[T]) rotateLeft() *AVLNode[T] {
	if node == nil || node.right == nil {
		return node
	}

	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node

	node.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// Inserting the right child into the left child tree causes imbalance, first rotating left and then right
func (node *AVLNode[T]) rotateLeftRight() *AVLNode[T] {
	if node == nil || node.left == nil {
		return node
	}

	node.left = node.left.rotateLeft()
	return node.rotateRight()
}

// Inserting the left child into the right tree causes imbalance, with the right one rotating first and then the left one
func (node *AVLNode[T]) rotateRightLeft() *AVLNode[T] {
	if node == nil || node.right == nil {
		return node
	}

	node.right = node.right.rotateRight()
	return node.rotateLeft()
}

// insert key into avl tree
func (tree *AVLTree[T]) Insert(key T) {
	tree.root = insertNode(tree.root, key)
	tree.size++
}

// insert key into node
func insertNode[T cmp.Ordered](node *AVLNode[T], key T) *AVLNode[T] {
	// Initialize node
	if node == nil {
		return &AVLNode[T]{key: key, height: 1}
	}

	// insert key according to Ord
	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}

	node.updateHeight()

	balance := node.computeBalanceFactor()

	// left left imbalance
	if balance > AVL_LEFT_HEAVY && key < node.left.key {
		return node.rotateRight()
	}

	// right right imbalance
	if balance < AVL_RGT_HEAVY && key > node.right.key {
		return node.rotateLeft()
	}

	// left right imbalance
	if balance > AVL_LEFT_HEAVY && key > node.left.key {
		return node.rotateLeftRight()
	}

	// right left imbalance
	if balance < AVL_RGT_HEAVY && key < node.right.key {
		return node.rotateRightLeft()
	}

	// don't need rotate
	return node
}

// find the minimum node
func findMin[T cmp.Ordered](node *AVLNode[T]) *AVLNode[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// delete key if the key in avl tree
func (tree *AVLTree[T]) Delete(key T) bool {
	oldSize := tree.size
	tree.root = deleteNode(tree.root, key)
	return oldSize > tree.size
}

// delete key if the key in avl node
func deleteNode[T cmp.Ordered](node *AVLNode[T], key T) *AVLNode[T] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = deleteNode(node.left, key)
	} else if key > node.key {
		node.right = deleteNode(node.right, key)
	} else {
		if node.left == nil || node.right == nil {
			var child *AVLNode[T]
			if node.left != nil {
				child = node.left
			} else {
				child = node.right
			}

			if child == nil {
				return nil
			} else {
				*node = *child
			}
		} else {
			successor := findMin(node.right)
			node.key = successor.key
			node.right = deleteNode(node.right, successor.key)
		}
	}

	node.updateHeight()

	balance := node.computeBalanceFactor()

	if balance > AVL_LEFT_HEAVY {
		leftBalance := node.left.computeBalanceFactor()

		if leftBalance >= AVL_BALANCED {
			return node.rotateRight()
		}
		return node.rotateLeftRight()
	}

	if balance < AVL_RGT_HEAVY {
		rightBalance := node.right.computeBalanceFactor()
		if rightBalance <= AVL_BALANCED {
			return node.rotateLeft()
		}
		return node.rotateRightLeft()
	}

	return node
}

// search the key is exist avl tree
func (tree *AVLTree[T]) Exist(key T) bool {
	return existNode(tree.root, key) != nil
}

// search the key is exist avl node
func existNode[T cmp.Ordered](node *AVLNode[T], key T) *AVLNode[T] {
	if node == nil {
		return nil
	}
	if key == node.key {
		return node
	} else if key < node.key {
		return existNode(node.left, key)
	} else {
		return existNode(node.right, key)
	}
}

// desstroy avl tree
func (tree *AVLTree[T]) Destory() {
	tree.root = nil
	tree.size = 0
}

// insert a slice with  value of type T
func (tree *AVLTree[T]) Append(keys ...T) {
	for _, key := range keys {
		tree.Insert(key)
	}
}

// print avl tree
func (tree *AVLTree[T]) String() string {
	if tree == nil {
		return "Nil"
	} else {
		fmt.Printf("%s", "Root:")
		return stringNode(tree.root, 0)
	}
}

// print avl node
func stringNode[T cmp.Ordered](node *AVLNode[T], indent_num int) string {
	if node == nil {
		return "Nil"
	}
	indent_string := strings.Repeat(" ", indent_num)
	return fmt.Sprintf("%s\n%svalue:%v\n%sheight:%d\n%s  leftNode:%s\n%s  rightNode:%s\n", indent_string, "  "+indent_string, node.key, "  "+indent_string, node.height, indent_string+"  ", stringNode(node.left, indent_num+2), indent_string+"  ", stringNode(node.right, indent_num+2))
}

// from the avl tree get a sorted slice
func (tree *AVLTree[T]) GetSortedSlice() []T {
	if tree.root == nil {
		return nil
	}
	slice := make([]T, 0, tree.size)
	getSortedSliceHelper(tree.root, &slice)
	return slice
}

// from the avl node get a sorted slice
func getSortedSliceHelper[T cmp.Ordered](node *AVLNode[T], slice *[]T) {
	if node == nil {
		return
	}
	getSortedSliceHelper(node.left, slice)
	*slice = append(*slice, node.key)
	getSortedSliceHelper(node.right, slice)
}

// build the avl tree from sorted slice
func BuildFromSortedSlice[T cmp.Ordered](arr []T) *AVLTree[T] {
	tree := &AVLTree[T]{}
	tree.root = buildBalanced(arr, 0, len(arr)-1)
	tree.size = len(arr)
	return tree
}

// build the avl node from sorted slice
func buildBalanced[T cmp.Ordered](arr []T, start, end int) *AVLNode[T] {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &AVLNode[T]{
		key: arr[mid],
	}

	node.left = buildBalanced(arr, start, mid-1)
	node.right = buildBalanced(arr, mid+1, end)

	node.updateHeight()

	return node
}

func (tree *AVLTree[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		iterNode(tree.root, yield)
	}
}

func iterNode[T cmp.Ordered](node *AVLNode[T], yield func(T) bool) bool {
	if node == nil {
		return true
	}

	if !iterNode(node.left, yield) {
		return false
	}

	if !yield(node.key) {
		return false
	}

	return iterNode(node.right, yield)
}

func (tree *AVLTree[T]) GetSize() int{
	return tree.size
}
