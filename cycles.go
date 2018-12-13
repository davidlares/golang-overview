package main

import "fmt"

func main(){
  // go solo tiene for
  for i:= 0; i < 10; i++ {
    fmt.Println("Hello, David")
  }

  // otra forma
  a := 0
  for {
    if a == 2 {
      a++
      continue
    }

    fmt.Println(a)
    a++
    
    if a > 10 {
      break
    }
  }
}
