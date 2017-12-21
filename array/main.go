package main

import "fmt"

func main() {
	const a1l = 2
	var a1 [a1l]int
	a1[0] = 1
	a1[1] = 2
	fmt.Println(a1)

	const a2l = a1l + 1
	a2 := [a2l]string{
		"Hello",
		" ",
		"World",
	}
	fmt.Println(a2)
}
