package ch6

import "github.com/Suuu775/Mastering_Algorithms_with_C/src/ch5"

type Stack struct {
	list *ch5.List
}

func NewStack() *Stack {
	return &Stack{
		list: ch5.List_init(),
	}
}

func (s *Stack) Push(data any) {
	s.list.Ins_next(nil, data)
}

func (s *Stack) Pop() (any, error) {
	return s.list.Rem_next(nil)
}

func (s *Stack) Top() any {
	if head := s.list.Head(); head != nil {
		return head.Data()
	}
	return nil
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) Destroy() {
	s.list.Destroy()
}
