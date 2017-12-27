package main

import "fmt"

func main() {
    list := New([]string{"world", "world"})
    err := list.Update(0, "Hello ")
    if err != nil {
        panic(err)
    }
    cl := list.Clone()
    cl.Create("!\n")
    cl.Each(func(i int, s string) {
        fmt.Print(s)
    })
}
