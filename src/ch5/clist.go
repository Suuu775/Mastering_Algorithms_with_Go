package ch5

import "errors"

type CListElmt struct {
	data any
	next *CListElmt
}

type CList struct {
	len  int
	head *CListElmt
}

func CList_init() *CList {
	return &CList{}
}

func (lst *CList) Ins_next(element *CListElmt, data any) error {

	if element == nil && lst.len != 0 {
		return errors.New("don't insert an void element when dlist len is bigger 0")
	}
	new_element := &CListElmt{data: data}
	if lst.len == 0 {
		new_element.next = new_element
		lst.head = new_element
	} else {
		new_element.next = element.next
		element.next = new_element
	}

	lst.len++
	return nil
}

func (lst *CList) Rem_next(element *CListElmt) (any, error) {

	// 不允许空元素插入除非列表为空
	if element == nil && lst.len != 0 {
		return nil, errors.New("don't insert an void element when dlist len is bigger 0")
	}

	var old_element = &CListElmt{}
	if lst.len == 0 {
		return nil, errors.New("not allow remove element when list is empty")
	}
	data := element.next.data
	if element.next == element {
		// 处理删除最后一个元素
		// old_element = element.next
		lst.head = nil
	} else {
		// 处理删除最后一个元素以外的情况
		old_element = element.next
		element.next = element.next.next

		// 删除的为头节点
		if old_element == lst.head {
			lst.head = old_element.next
		}
	}
	lst.len--
	return data, nil
}

func (lst *CList) Len() int {
	return lst.len
}

func (lst *CList) Head() *CListElmt {
	return lst.head
}

func (lst *CListElmt) Data() any {
	return lst.data
}

func (lst *CListElmt) Next() *CListElmt {
	return lst.next
}
