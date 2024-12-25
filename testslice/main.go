package main

import (
	"fmt"
	"sort"
)

func main() {
	res := make([]int, 0)
	res = append(res, 1)
	changeSlice(res)
	fmt.Println(res)
	sort.Search()
}

func changeSlice(res []int) {
	res = append(res, 2)
}
