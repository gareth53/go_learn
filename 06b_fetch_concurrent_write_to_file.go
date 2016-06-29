package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "net/http"
  "os"
  "time"
)

const (
  prefix_check = "http"
  full_prefix = "http://"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
      go fetch(url, ch) // starring a goroutine
    }
    for range os.Args[1:] {
      fmt.Println(<-ch) // receive from chanel ch
    }
    fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func check_protocol(url string) (string) {
  if strings.HasPrefix(url, prefix_check) {
    return url
  }
  return full_prefix + url
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    full_url := check_protocol(url)
    resp, err := http.Get(full_url)
    if err != nil {
      ch <- fmt.Sprint(err)
      return
    }
    contents, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
      ch <- fmt.Sprintf("while reading %s: %v", url, err)
    }
    f, err := os.Create("temp.txt")
    defer f.Close()
    _, writeerr := f.Write(contents)
    if writeerr != nil {
      ch <- fmt.Sprintf("while writing %s: %v", url, writeerr)
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
