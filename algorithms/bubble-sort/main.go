package main

import "fmt"
import "sort"

func BubbleSort(list sort.Interface) {
	for itemCount := list.Len() - 1; ; itemCount-- {
		hasChanged := false
		for current := 0; current < itemCount; current++ {
			next := current + 1
			if list.Less(next, current) {
				list.Swap(current, next)
				hasChanged = true
			}
		}
		if !hasChanged {
			break
		}
	}
}

func main() {
	list := []int{5, 20, 3, 11, 1, 17, 3, 12, 8, 10}
	fmt.Print(list, " -> ")
	BubbleSort(sort.IntSlice(list))
	fmt.Println(list)
}
