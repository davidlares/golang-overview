package main

// importamos el paquete para IO
import (
  "fmt"
  "strconv"
)

func main(){
  age := 22
  // integer to string
  age_str := strconv.Itoa(age)
  // string age
  str_age := "22"
  // prnting message
  fmt.Println("My name is: "  + age_str)
  // usning strAge variable -> functions returns multiple values, you have to place _ for receiving something that you won't use
  strAge,_ := strconv.Atoi(str_age)
  fmt.Println(10 + strAge)

}
