package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.URL, r.Method, r.Proto)
    for key, val := range r.Header {
      fmt.Fprintf(w, "[%q] = %q\n", key, val)
    }
    fmt.Fprintf(w, "Host = %s\n", r.Host)
    fmt.Fprintf(w, "Remote Address = %s\n", r.RemoteAddr)
}
