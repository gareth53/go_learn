package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "bufio"
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

    input_filepath := "06c_urls.txt"
    input_file, err := os.Open(input_filepath)
    if err != nil {
      fmt.Printf("Error! %s\n", err)
    }
    input := bufio.NewScanner(input_file)
    ch := make(chan string)
    urls := 0
    for input.Scan() {
      fmt.Print(input.Text())
      fmt.Print("\n")
      go fetch(input.Text(), ch) // starring a goroutine
    }
    fmt.Print(urls)
    for q := 0; q < urls; q++ {
      fmt.Print(q)
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
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
      ch <- fmt.Sprintf("while reading %s: %v", url, err)
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
