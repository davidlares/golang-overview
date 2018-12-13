package main

import "fmt"

type User struct {
  age int
  name, lastname string
}


func main(){
  // estructuras
  var david User
  fmt.Println(david)
  fmt.Println(User{name: "david", lastname: "lares", age: 28})
  // inicializando
  daniel := User{name: "daniel", lastname: "lares", age: 28}
  fmt.Println(daniel)
  // otra forma
  lucy := User{20, "lucy","novoa"}
  fmt.Println(lucy)
  // estructura ahora con enteros
  uriel := new(User)
  uriel.name = "Uriel"
  fmt.Println(uriel.name)
}
