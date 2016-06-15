package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
    for _, url := range os.Args[1:] {
      print_url(url)
    }
}

func print_url(url string) {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error Fetching: %v\n", err)
    os.Exit(1)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error Reading: %v\n", err)
    os.Exit(1)
  }
  fmt.Printf("%s", body)
}
