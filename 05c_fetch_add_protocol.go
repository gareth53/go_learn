package main

import (
  "fmt"
  "io"
  "strings"
  "net/http"
  "os"
)

const (
  prefix_check = "http"
  full_prefix = "http://"
)

func main() {
    for _, url := range os.Args[1:] {
      full_url := check_protocol(url)
      print_url(full_url, os.Stdout)
    }
}

func check_protocol(url string) (string) {
  if strings.HasPrefix(url, prefix_check) {
    return url
  }
  return full_prefix + url
}

func print_url(url string, out io.Writer) {
  resp, err_get := http.Get(url)
  if err_get != nil {
    fmt.Fprintf(os.Stderr, "Error Fetching: %v\n", err_get)
    os.Exit(1)
  }
  fmt.Println("HTTP Response: ", resp.Status)
  _, err_copy := io.Copy(out, resp.Body)
  if err_copy != nil {
    fmt.Fprintf(os.Stderr, "Error copying: %v\n", err_copy)
    os.Exit(1)
  }
//  fmt.Printf("%s", body)
}
