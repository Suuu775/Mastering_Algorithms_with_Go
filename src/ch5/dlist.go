package ch5

import "errors"

type DListElmt struct {
	data any
	prev *DListElmt
	next *DListElmt
}

type DList struct {
	len  int
	head *DListElmt
	tail *DListElmt
}

func DList_init() *DList {
	return &DList{}
}

func (lst *DList) Ins_next(element *DListElmt,data any) error{
	new_element := &DListElmt{data: data}

	// 不允许空元素插入除非列表为空
	if  element == nil && lst.len !=0{
		return errors.New("Don't insert an void element when dlist len is bigger 0")
	}

	if lst.len == 0{
		lst.head = new_element;
		lst.head.prev = nil;
		lst.head.next = nil;
		lst.tail = new_element;
	} else {
		new_element.next = element.next;
		new_element.prev = element;

		if element.next == nil{
			lst.tail = new_element
		} else {
			element.next.prev = new_element;
		}
		element.next = new_element;
	}
	lst.len++
	return nil
}

func (lst *DList) Ins_prev(element *DListElmt,data any) error{
	new_element := &DListElmt{data: data}

	// 不允许空元素插入除非列表为空
	if  element == nil && lst.len !=0{
		return errors.New("Don't insert an void element when dlist len is bigger 0")
	}

	if lst.len == 0{
		// 当列表为空时处理插入
		lst.head = new_element
		lst.head.prev = nil
		lst.head.next = nil
		lst.tail = new_element
	} else {
		// 当列表不为空时处理插入
		new_element.next = element
		new_element.prev = element.prev

		if element.prev == nil{
			lst.head = new_element
		} else {
			element.prev.next = new_element
		}
		element.prev = new_element
	}
	lst.len++
	return nil
}

func (lst *DList) Remove(element *DListElmt) (any, error) {

    if lst.len == 0 || element == nil {
        return nil, errors.New("cannot remove from empty list or nil element")
    }

    data := element.data
    if element == lst.head {
        lst.head = element.next
        if lst.head == nil {
            lst.tail = nil
        } else {
            lst.head.prev = nil
        }
    } else {
        element.prev.next = element.next
        if element.next == nil {
            lst.tail = element.prev
        } else {
            element.next.prev = element.prev
        }
    }

    lst.len--

    return data, nil
}

func (lst *DList) Destroy(){
	for lst.len != 0{
		lst.Remove(lst.head)
	}
}

func (lst *DList) Len() int{
	return lst.len
}

func (lst *DList) Head() *DListElmt{
	return lst.head
}

func (lst *DList) Tail() *DListElmt{
	return lst.tail
}

func (lst *DList) Is_head(elmt *DListElmt) bool {
	return lst.Head() == elmt
}

func (lst *DList) Is_tail(elmt *DListElmt) bool {
	return lst.Tail() == elmt
}

func (elmt *DListElmt) Next() *DListElmt{
	return elmt.next
}

func (elmt *DListElmt) Prev() *DListElmt{
	return elmt.prev
}
