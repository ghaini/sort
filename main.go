package main

import (
	"fmt"
	"sort/sort"
)

func main() {
	s := sort.NewSort([]int{4, 2, 7, 3, 8, 89, 65})
	fmt.Println(s.MergeSort())
}
