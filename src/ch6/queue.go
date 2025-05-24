package ch6

import "github.com/Suuu775/Mastering_Algorithms_with_C/src/ch5"

type Queue struct{
	list *ch5.List
}

// 构造 queeu
func NewQueue() *Queue{
	return &Queue{
		list:ch5.List_init(),
	}
}

// 删除首个元素
func (queue *Queue) Pop() (any,error){
	return queue.list.Rem_next(nil)
}

// 向队列尾部插入元素
func (queue *Queue) Push(data any){
	if queue.list.Size() == 0 {
        queue.list.Ins_next(nil, data)
    } else {
        queue.list.Ins_next(queue.list.Tail(), data)
    }
}

// 访问第一个元素
func (queue *Queue) Front() any{
	if head := queue.list.Head(); head != nil {
		return head.Data()
	}
	return nil
}

// 访问最后一个元素
func (queue *Queue) Back() any{
	if tail := queue.list.Tail() ; tail != nil {
		return tail.Data()
	}
	return nil
}

// 返回容纳的元素数
func (queue *Queue) Size() int{
	return queue.list.Size()
}
