package main

import "fmt"

type User struct {
    name string
    year int
}

func (u *User) PrintInfo() {
    fmt.Printf("Name: %s: year %d\n", u.name, u.year)
}

func main() {
    u := new(User)
    u.name = "Dzyanis"
    u.year = 1987
    u.PrintInfo()
}
