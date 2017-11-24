package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	for {
		fmt.Fprintln(os.Stdout, rand.Int63())
	}
}
