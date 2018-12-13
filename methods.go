package main

import "fmt"

type User struct {
  age int
  name, lastname string
}

func (user User) fullname() string {
  return user.name + " " + user.lastname
}

func (user *User) set_name(n string) {
  user.name = n
  // al hacerlo un puntero la vaina se sobreescribe
  // al pasar una referencia es mas ligero
}

func main(){
  david := new(User)
  david.name = "David"
  david.set_name("Lares") // se genera otra estructura, si no se pasa como un puntero
  fmt.Println(david.fullname())
}
