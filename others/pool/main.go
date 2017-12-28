package main

import (
    "sync"
    "bytes"
    "io"
    "time"
    "os"
)

// limit capacity of the pool
const poolCap = 100

// limit []byte cap
const maxCap = 1024

// create pool
var chanPool = make(chan bytes.Buffer, poolCap)

func getBuffer() (b bytes.Buffer) {
    select {
    case bt, ok := <-chanPool:
        if ok {
            return bt
        }
        // non-normal behaivor, need fix!
        panic("Who, the fuck, closed the cahnnel of the chanPool")
    default:
    }
    return // nil<[]byte>
}

func putBuffer(b bytes.Buffer) {
    // b = nil
    select {
    case chanPool <- b:
    default:
    }
    return
}

func LogChan(w io.Writer, key, val string) {
    b := getBuffer()
    b.Reset()
    b.WriteString(time.Now().UTC().Format(time.RFC3339))
    b.WriteByte(' ')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
    putBuffer(b)
}

var pool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func LogPool(w io.Writer, key, val string) {
    b := pool.Get().(*bytes.Buffer)
    b.Reset()
    b.WriteString(time.Now().UTC().Format(time.RFC3339))
    b.WriteByte(' ')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
    pool.Put(b)
}

func Log(w io.Writer, key, val string) {
    b := bytes.NewBuffer([]byte{})
    b.Reset()
    b.WriteString(time.Now().UTC().Format(time.RFC3339))
    b.WriteByte(' ')
    b.WriteString(key)
    b.WriteByte('=')
    b.WriteString(val)
    w.Write(b.Bytes())
}

func main() {
    Log(os.Stdout, "path", "/search?q=flowers\n")
    LogPool(os.Stdout, "path", "/search?q=flowers\n")
    LogChan(os.Stdout, "path", "/search?q=flowers\n")
}
