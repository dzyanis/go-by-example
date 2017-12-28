package main

import (
    "testing"
    "bytes"
)

func BenchmarkLog(b *testing.B) {
    w := bytes.NewBuffer([]byte{})
    for n := 0; n < b.N; n++ {
        Log(w,"path", "/search?q=flowers")
    }
}

func BenchmarkLogPool(b *testing.B) {
    w := bytes.NewBuffer([]byte{})
    for n := 0; n < b.N; n++ {
        LogPool(w,"path", "/search?q=flowers")
    }
}

func BenchmarkLogChan(b *testing.B) {
    w := bytes.NewBuffer([]byte{})
    for n := 0; n < b.N; n++ {
        LogChan(w,"path", "/search?q=flowers")
    }
}

