package main

import (
	"fmt"
	"os"
)

func GetFileContent(filename string) (content string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	content = string(bs)
	file.Close()

	return
}

func main() {
	content, err := GetFileContent("bio.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
