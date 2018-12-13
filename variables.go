package main

import "fmt"

func main(){
  // con el var
  var x int
  x = 12
  // sin el var -> toma el tipo de datos con su valor tal
  a := 23
  // cadenas
  nombre := "David"
  // no puede haber una reasignacion de valores para variables asignadas :=, ni para distintos tipos de datos
  fmt.Println(x)
  fmt.Println(a)
  fmt.Println(nombre)
}
