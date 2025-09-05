package ch9

import "errors"

type BiTreeNode[T comparable] struct {
	data  T
	left  *BiTreeNode[T]
	right *BiTreeNode[T]
}

type BiTree[T comparable] struct {
	size uint
	root *BiTreeNode[T]
}

// 创建BiTree
func NewBiTree[T comparable]() *BiTree[T] {
	return &BiTree[T]{
		size: 0,
		root: nil,
	}
}

// 在指定节点的左子节点位置插入一个新节点
// 若指定节点为空且树为空树,则该节点为根节点,值为data
// 若指定节点不为空且指定节点的左子树为空,插入新节点作为其左子节点
func (tree *BiTree[T]) Ins_left(node *BiTreeNode[T], data T) error {
	new_node := &BiTreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
	var position **BiTreeNode[T]

	// 决定往哪个节点插入
	if node == nil {
		if tree.Size() > 0 {
			return errors.New("only empty tree allow insert")
		}
		position = &tree.root
	} else {
		if node.Left() != nil {
			return errors.New("only allow insertion at the end of branch")
		}
		position = &node.left
	}

	*position = new_node
	tree.size++

	return nil
}

// 在指定节点的右子节点位置插入一个新节点
// 若指定节点为空且树为空树,则该节点为根节点,值为data
// 若指定节点不为空且指定节点的右子树为空,插入新节点作为其右子节点
func (tree *BiTree[T]) Ins_right(node *BiTreeNode[T], data T) error {
	new_node := &BiTreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}

	var position **BiTreeNode[T]

	if node == nil {
		if tree.size > 0 {
			return errors.New("only empty tree allow insert")
		}
		position = &tree.root
	} else {
		if node.right != nil {
			return errors.New("only allow insertion at the end of branch")
		}
		position = &node.right
	}

	*position = new_node
	tree.size++

	return nil
}

// 移除指定节点的左子树
func (t *BiTree[T]) Remove_left(node *BiTreeNode[T]) {
	if t.size == 0 {
		return
	}

	// 确定要移除的位置
	var position **BiTreeNode[T]
	if node == nil {
		position = &t.root
	} else {
		position = &node.left
	}

	// 递归移除子树
	if *position != nil {
		t.removeSubtree(*position)
		*position = nil
	}
}

// RemoveRight 移除指定节点的右子树
func (t *BiTree[T]) Remove_right(node *BiTreeNode[T]) {
	if t.size == 0 {
		return
	}

	// 确定要移除的位置
	var position **BiTreeNode[T]
	if node == nil {
		position = &t.root
	} else {
		position = &node.right
	}

	// 递归移除子树
	if *position != nil {
		t.removeSubtree(*position)
		*position = nil
	}
}

// 递归移除子树的核心方法
func (t *BiTree[T]) removeSubtree(node *BiTreeNode[T]) {
	if node == nil {
		return
	}

	// 递归移除左子树
	if node.left != nil {
		t.removeSubtree(node.left)
		node.left = nil
	}

	// 递归移除右子树
	if node.right != nil {
		t.removeSubtree(node.right)
		node.right = nil
	}

	// 更新树大小
	t.size--
}

func Bitree_merge[T comparable](merge *BiTree[T], left *BiTree[T], right *BiTree[T], data T) error {
    *merge = *NewBiTree[T]()


    if err := merge.Ins_left(nil, data); err != nil {
        return err
    }


    merge.root.left = left.root
    merge.root.right = right.root


    merge.size = 1 + left.size + right.size

    left.root = nil
    left.size = 0
    right.root = nil
    right.size = 0

    return nil
}

func (tree *BiTree[T]) Size() uint {
	return tree.size
}

func (tree *BiTree[T]) Root() *BiTreeNode[T] {
	return tree.root
}

func (node *BiTreeNode[T]) Is_eob() bool {
	return node == nil
}

func (node *BiTreeNode[T]) Is_leaf() bool {
	return node.left == nil && node.right == nil
}

func (node *BiTreeNode[T]) Data() T {
	return node.data
}

func (node *BiTreeNode[T]) Left() *BiTreeNode[T] {
	return node.left
}

func (node *BiTreeNode[T]) Right() *BiTreeNode[T] {
	return node.right
}
