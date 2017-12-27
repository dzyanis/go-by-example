package main

import (
    "net"
    "log"
    "bufio"
    "strings"
)

func main() {
    log.Println("Server start")
    listener, err := net.Listen("tcp", "127.0.0.1:3333")
    if err != nil {
        log.Fatal(err)
    }

    conn, err := listener.Accept()
    if err != nil {
        log.Fatal(err)
    }

    message, _ := bufio.NewReader(conn).ReadString('\n')
    log.Printf("Message: %s", string(message))
    message = strings.ToUpper(message)
    conn.Write([]byte(message + "\n"))

    if err := conn.Close(); err != nil {
        log.Fatal(err)
    }
    log.Println("Connection closed")

    if err := listener.Close(); err != nil {
        log.Fatal(err)
    }
    log.Println("Exit")
}