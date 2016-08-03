package main

import (
  "strings"
)

func main() {
  tests := [10][2]string{
    // truthy test
    {"1234", "1234"},
    {"1234", "4321"},
    {"1234", "2314"},
    {"anagram", "manraag"},
    {"anagram", "Anagram"},
    {"anagram", "MARGANA"},
    // falsy tests
    {"anagram", "ana gram"},
    {"1112", "4321"},
    {"12345", "1234"},
    {"1234", "1523"},
  }
  for _, v := range tests {
    println(isAnagram(v[0], v[1]))
  }
}

func isAnagram(s1, s2 string) bool {
  /*
  test length is equal
  search for each rune in each other
  */
  lngth := len(s1)
  if lngth != len(s2) {
    return false
  }
  // let's be case-insensitive
  s1 = strings.ToUpper(s1)
  s2 = strings.ToUpper(s2)
  for i:=0; i<lngth ; i++ {
    if strings.Contains(s1, s2[i:i+1]) != true {
      return false
    }
    if strings.Contains(s2, s1[i:i+1]) != true {
      return false
    }
  }
  return true
}
