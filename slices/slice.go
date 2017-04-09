package main


import "fmt"


func main() {
	a := make([]int, 30)
	fmt.Printf("1. A: %#v\n", a)
	b := a[1:16]
	fmt.Printf("2. B: %#v\n", b)
	a[2] = 13
	fmt.Printf("3. A: %#v\n", a)
	fmt.Printf("   B: %#v\n", b)

	a = append(a, 1)
	a[2] = 42
	fmt.Printf("4. A: %#v\n", a)
	fmt.Printf("   B: %#v\n", b)
	// You can read more about this gotcha here https://divan.github.io/posts/avoid_gotchas/
}
