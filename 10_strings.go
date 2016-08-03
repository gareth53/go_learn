package main

import (
  "fmt"
)

func main() {
  var s1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
  for i:=0; i<len(s1); i++ {
//    println(s1[i], s1[i:i+1])
  }
  HasPrefix("North London", "North")
  HasPrefix("North London", "South")
  HasSuffix("North London", "North")
  HasSuffix("North London", "London")
  Contains("North London", "London")
  Contains("North London", "Bristol")
}

func HasPrefix(s, prefix string) bool {
    found := len(s) > len(prefix) && s[0:len(prefix)] == prefix
    var msg string
    if found {
      msg = "is a prefix of"
    } else {
      msg = "is NOT a prefix of"
    }
    fmt.Printf("%s %s %s\n", prefix, msg, s)
    return found
}

func HasSuffix(s, suffix string) bool {
    found := len(s) > len(suffix) && s[len(suffix):] == suffix
    var msg string
    if found {
      msg = "is a suffix of"
    } else {
      msg = "is NOT a suffix of"
    }
    fmt.Printf("%s %s %s\n", suffix, msg, s)
    return found
}

func Contains(s, substr string) bool {
    if len(s) > len(substr) {
        limit := len(s)-len(substr)
        for i:=0; i<=limit; i++ {
          if s[i:i+len(substr)] == substr {
            fmt.Printf("%s found in %s\n", substr, s)
            return true
          }
        }
    }
    fmt.Printf("%s NOT found in %s\n", substr, s)
    return false
}
