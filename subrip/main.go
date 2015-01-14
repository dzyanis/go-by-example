package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    // "strconv"
    // "strings"
)

type Subtitle []SubtitlePiece

func (subs *Subtitle) SaveHTML(filename string) (err error) {
    fo, err := os.Create(filename)
    if (err != nil) {
        return err
    }
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    w := bufio.NewWriter(fo)
    for _, value := range *subs {
        fmt.Fprintf(w, "<div>%s</div>", value.content)
    }
    w.Flush()
    return nil
}

type SubtitlePiece struct{
    index int64;
    timeStart time.Time;
    timeFinish time.Time;
    content string;
}

func (subs *Subtitle) Add(piece *SubtitlePiece) {
    *subs = append(*subs, *piece)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    filename := os.Args[1]
    file, err := os.Open(filename)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    subs := new(Subtitle)
    for scanner.Scan() {
        if scanner.Text() == "" {
            scanner.Scan()
        }
        piece := new(SubtitlePiece)
        //ind := scanner.Text()
        //fmt.Println(ind)
        //piece.index, err = strconv.Atoi(ind)
        //piece.index, err = strconv.ParseInt(ind, 0, 32)
        //check(err)
        scanner.Scan()
        interval := scanner.Text()
        piece.timeStart, err = time.Parse(time.RFC3339, interval[:12])
        piece.timeFinish, err = time.Parse(time.RFC3339, interval[17:])
        for scanner.Scan() {
            if scanner.Text() == "" {
                break;
            }
            piece.content = piece.content + " " + scanner.Text()
        }
        // fmt.Println(block.Content)
        subs.Add(piece)

        err := scanner.Err()
        check(err)
    }

    err = subs.SaveHTML("output.html")
    check(err)
}
