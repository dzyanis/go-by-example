package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Hello" + " World"
	fmt.Printf("Value '%v' = %v\n", str, reflect.TypeOf(str)) // string

	i := 1 + 1
	fmt.Printf("Value %v = %v\n", i, reflect.TypeOf(i)) // int

	f := float32(i) / 2
	fmt.Printf("Value %v = %v\n", f, reflect.TypeOf(f)) // float64

	d := f*0.9876543219876543321
	fmt.Printf("Value %v = %v\n", d, reflect.TypeOf(d)) // float64

	ch := 'A'
	fmt.Printf("Value %v = %v\n", ch, reflect.TypeOf(ch)) // int32

	b := true
	fmt.Printf("Value %v = %v\n", b, reflect.TypeOf(b)) // bool
}
