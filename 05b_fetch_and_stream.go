package main

import (
  "fmt"
  "io"
//  "io/ioutil"
  "net/http"
  "os"
)

func main() {
    for _, url := range os.Args[1:] {
      print_url(url, os.Stdout)
    }
}
/*
THe function call io.Copy(dst, src) reads from src and copies to dst.
Use it instaed of ioutil.ReadAll to copy the response body to os.Stdout
without requiring a buffer large enough to hold the entire stream.
Be sure to check the error statu f io.Copy.

*/
func print_url(url string, out io.Writer) {
  resp, err_get := http.Get(url)
  if err_get != nil {
    fmt.Fprintf(os.Stderr, "Error Fetching: %v\n", err_get)
    os.Exit(1)
  }
  _, err_copy := io.Copy(out, resp.Body)
  if err_copy != nil {
    fmt.Fprintf(os.Stderr, "Error copying: %v\n", err_copy)
    os.Exit(1)
  }
//  fmt.Printf("%s", body)
}
