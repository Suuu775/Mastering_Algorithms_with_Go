package ch5

import (
	"errors"
)

type ListElmt struct {
	data any
	next *ListElmt
}

type List struct {
	size int
	head *ListElmt
	tail *ListElmt
}

func List_init() *List {
	return &List{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// 在指定节点 element 后插入新节点
func (lst *List) Ins_next(element *ListElmt, data any) {
	new_node := &ListElmt{data: data}

	if element == nil {
		new_node.next = lst.head
		lst.head = new_node

		// 更新尾节点
		if lst.size == 0 {
			lst.tail = new_node
		}
	} else {
		new_node.next = element.next
		element.next = new_node

		// 更新尾节点
		if new_node.next == nil {
			lst.tail = new_node
		}
	}
	lst.size++
}

func (l *List) Destroy() {
	for l.size > 0 {
		l.Rem_next(nil)
	}
}

// 在指定节点 element 后删除新节点
func (lst *List) Rem_next(element *ListElmt) (any, error) {
	if lst.size == 0 {
		return nil, errors.New("empty list")
	}

	var target *ListElmt

	if element == nil {
		target = lst.head
		lst.head = target.next
		if lst.size == 1 {
			lst.tail = nil
		}
	} else {
		if element.next == nil {
			return nil, errors.New("no next element")
		}
		target = element.next
		element.next = target.next
		if element.next == nil {
			lst.tail = element
		}
	}
	lst.size--
	return target.data, nil
}

func (lst *List) Size() int {
	return lst.size
}

func (lst *List) Head() *ListElmt {
	return lst.head
}

func (lst *List) Is_head(elmt *ListElmt) bool {
	return lst.head == elmt
}

func (lst *List) Tail() *ListElmt {
	return lst.tail
}

func (lst *List) Is_Tail(elmt *ListElmt) bool {
	return lst.Tail() == elmt
}

func (elmt *ListElmt) Data() any {
	return elmt.data
}

func (elmt *ListElmt) Next() *ListElmt {
	return elmt.next
}
