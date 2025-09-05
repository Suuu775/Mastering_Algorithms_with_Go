package main

import (
	"fmt"

	"github.com/Suuu775/Mastering_Algorithms_with_C/src/ch9"
)

func main() {
	bistree := ch9.NewAVLTree[int]()
	bistree.Append(1, 2, 3, 5, 6, 7)
	fmt.Println(bistree)
}
