package main

import (
    "net"
    "log"
    "fmt"
    "bufio"
    "os"
)

func main() {
    log.Println("Client start")
    conn, err := net.Dial("tcp", "127.0.0.1:3333")
    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(os.Stdin)
    log.Print("enter: ")
    text, _ := reader.ReadString('\n')
    fmt.Fprintf(conn, text + "\n")
    message, _ := bufio.NewReader(conn).ReadString('\n')
    log.Printf("return: %s", message)

    if err := conn.Close(); err != nil {
        log.Fatal(err)
    }
    log.Println("Exit")
}