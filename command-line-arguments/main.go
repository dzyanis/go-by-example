package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for k, v := range args {
		fmt.Println(k, v)
	}
}
