package main

import "fmt"

func reverse(s []int) {
  for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}


func rotate(s []int, i int) {
    reverse(s[:i])
    reverse(s[i:])
    reverse(s)
}

func rotation(s []int, i int) {
  /*
  write a version of roate that operaties in one pass
  */
  lim := len(s)
  // the originla implementation didn't have this safety check
  // but it's simple enough to do, so let;s do it...
  if (i > lim) {
    i = i % lim
  }
  temp := make([]int, lim)
  copy(temp, s)
  for q:=0; q<lim; q++ {
    newpos := q - i
    if newpos < 0 {
      newpos = newpos + lim
    }
    s[newpos] = temp[q]
  }
}

func main() {
      input := [...]int{0, 1, 2, 3, 4, 5}
      rotate(input[:], 2)
      fmt.Println(input)

      rotation(input[:], 13)
      fmt.Println(input)

}
