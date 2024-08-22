package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(6)
	}

	intSlice := sort.IntSlice(arr)
	intSlice.Sort()

	//找到第一个大于等于t的下标
	t := 3
	ind, ok := sort.Find(intSlice.Len(), func(i int) int {
		if t > intSlice[i] {
			return 1
		} else if t < intSlice[i] {
			return -1
		}
		return 0
	})
	fmt.Println(intSlice, ind, ok)
	sort.Sort(sort.Reverse(intSlice))
	ind = sort.Search(intSlice.Len(), func(i int) bool { //找到比t小的数中最大的哪一个
		return intSlice[i] < t
	})
	fmt.Println(intSlice, ind)

	t = 4
	temp := []int{0, 1, 2, 3, 4, 4, 5}
	//lower_bound
	ind1 := sort.Search(len(temp), func(i int) bool {
		return temp[i] >= t
	})

	//upper_bound
	ind2 := sort.Search(len(temp), func(i int) bool {
		return temp[i] > t
	})
	fmt.Println(temp, ind1, ind2)
}
