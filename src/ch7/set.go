package ch7

import (
	"cmp"
	"iter"

	"github.com/Suuu775/Mastering_Algorithms_with_C/src/ch9"
)

type Set[T cmp.Ordered] struct {
	tree *ch9.AVLTree[T]
}

func NewSet[T cmp.Ordered]() *Set[T] {
	return &Set[T]{
		tree: ch9.NewAVLTree[T](),
	}
}

func (set *Set[T]) Destory() {
	set.tree.Destory()
}

func (set *Set[T]) Insert(key T) {
	set.tree.Insert(key)
}

func (set *Set[T]) Remove(key T) bool {
	return set.tree.Delete(key)
}

func (lhs *Set[T]) Union(rhs *Set[T]) *Set[T] {
	lhs_slice := lhs.tree.GetSortedSlice()
	rhs_slice := rhs.tree.GetSortedSlice()

	new_set := NewSet[T]()
	new_set.tree = ch9.BuildFromSortedSlice(union_slice(lhs_slice, rhs_slice))
	return new_set
}

func (lhs *Set[T]) Intersection(rhs *Set[T]) *Set[T] {
	lhs_slice := lhs.tree.GetSortedSlice()
	rhs_slice := rhs.tree.GetSortedSlice()

	new_set := NewSet[T]()
	new_set.tree = ch9.BuildFromSortedSlice(intersection_slice(lhs_slice, rhs_slice))
	return new_set
}

func (lhs *Set[T]) Difference(rhs *Set[T]) *Set[T]{
	lhs_slice := lhs.tree.GetSortedSlice()
	rhs_slice := rhs.tree.GetSortedSlice()

	new_set := NewSet[T]()
	new_set.tree = ch9.BuildFromSortedSlice(difference_slice(lhs_slice, rhs_slice))
	return new_set
}

func (set *Set[T]) IsMember(key T) bool{
	return set.tree.Exist(key)
}

func(lhs *Set[T]) IsSubset (rhs *Set[T]) bool{
	if lhs.GetSize() > rhs.GetSize(){
		return false
	}

	for pairs := range zip(lhs.tree.Iter(),rhs.tree.Iter()) {
		if pairs[0] != pairs[1]{
			return false
		}
	}
	return true
}

func (lhs *Set[T]) IsEqual(rhs *Set[T]) bool{
	if lhs.GetSize() == rhs.GetSize() && lhs.IsSubset(rhs){
		return true
	} else {
		return false
	}
}

func (set *Set[T]) GetSize() int{
	return set.tree.GetSize()
}

func union_slice[T cmp.Ordered](s1 []T, s2 []T) []T {
	new_slice := []T{}
	i, j := 0, 0
	len1, len2 := len(s1), len(s2)

	for i < len1 && j < len2 {
		if s1[i] == s2[j] {
			i++
			j++
		} else if s1[i] < s2[j] {
			new_slice = append(new_slice, s1[i])
			i++
		} else {
			new_slice = append(new_slice, s2[j])
			j++
		}
	}

	for i < len1 {
		new_slice = append(new_slice, s1[i])
		i++
	}

	for j < len2 {
		new_slice = append(new_slice, s2[j])
		j++
	}

	return new_slice
}

func intersection_slice[T cmp.Ordered](s1 []T, s2 []T) []T {
	new_slice := []T{}
	i, j := 0, 0
	len1, len2 := len(s1), len(s2)

	for i < len1 && j < len2 {
		if s1[i] == s2[j] {
			new_slice = append(new_slice, s1[i])
			i++
			j++
		} else if s1[i] < s2[j] {
			i++
		} else {
			j++
		}
	}

	return new_slice
}

func difference_slice[T cmp.Ordered](s1 []T, s2 []T) []T {
	var new_slice []T
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			new_slice = append(new_slice, s1[i])
			i++
		} else if s1[i] == s2[j] {
			i++
			j++
		} else {
			j++
		}
	}

	for i < len(s1) {
		new_slice = append(new_slice, s1[i])
		i++
	}

	return new_slice
}

func zip[T any](seq1, seq2 iter.Seq[T]) <-chan [2]T {
	out := make(chan [2]T)

	go func() {
		defer close(out)

		ch1 := make(chan T, 1)
		ch2 := make(chan T, 1)

		go func() {
			seq1(func(v T) bool {
				select {
				case ch1 <- v:
					return true
				default:
					return false
				}
			})
			close(ch1)
		}()

		go func() {
			seq2(func(v T) bool {
				select {
				case ch2 <- v:
					return true
				default:
					return false
				}
			})
			close(ch2)
		}()

		for v1, ok1 := <-ch1; ok1; v1, ok1 = <-ch1 {
			v2, ok2 := <-ch2
			if !ok2 {
				break
			}
			out <- [2]T{v1, v2}
		}
	}()

	return out
}
