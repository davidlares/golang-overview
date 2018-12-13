package main

import "fmt"

type Human struct {
  name string
}

func (this Human) hablar() string {
  return "bla bla bla"
}

type Dummy struct {
  name string
}

type Tutor struct {
  Human
  Dummy
}

func main(){
  // replicar comportamiento de herencias, aunque Go no es POO
  tutor := Tutor{Human{"David"}, Dummy{"Dummy"}}
  fmt.Println(tutor.Human.name)
  fmt.Println(tutor.Dummy.name)
  fmt.Println(tutor.hablar())
}
