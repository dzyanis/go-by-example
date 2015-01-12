package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    // "strconv"
    // "strings"
)

type Subtitle []SubtitleBlock

type SubtitleBlock struct{
    Index int64;
    TimeStart time.Time;
    TimeFinish time.Time;
    Content string;
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    file, err := os.Open("example.srt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var subs Subtitle
    for scanner.Scan() {
        if scanner.Text() == "" {
            scanner.Scan()
        }
        var block SubtitleBlock;
        //ind := scanner.Text()
        //fmt.Println(ind)
        //block.Index, err = strconv.Atoi(ind)
        //block.Index, err = strconv.ParseInt(ind, 0, 32)
        //check(err)
        scanner.Scan()
        interval := scanner.Text()
        block.TimeStart, err = time.Parse(time.RFC3339, interval[:12])
        block.TimeFinish, err = time.Parse(time.RFC3339, interval[17:])
        for scanner.Scan() {
            if scanner.Text() == "" {
                break;
            }
            block.Content = block.Content + " " + scanner.Text()
        }
        // fmt.Println(block.Content)
        subs = append(subs, block)

        err := scanner.Err()
        check(err)
    }

    fo, err := os.Create("output.html")
    check(err)
    defer fo.Close();
    w := bufio.NewWriter(fo)
    for _, value := range subs {
        fmt.Fprintf(w, "<div>%s</div>", value.Content)
    }
    w.Flush()
}
