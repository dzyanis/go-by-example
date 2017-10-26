package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    var sum uint64

    file, err := os.Open("list.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        number, err := strconv.ParseUint(scanner.Text(), 10, 64)
        if err != nil {
            fmt.Println(err)
            continue
        }
        sum += number
    }

    fmt.Println("Sum:", sum)
}