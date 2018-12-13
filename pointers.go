package main

import "fmt"

func main(){
  // punteros -> direccion de memoria, en vez del valor, tenemos la direccion donde esta el valor.
  // Modificar el mismo dato a traves de diferntes partes del codigo
  // *T -> tipo de dato => *int, *string, *Struct
  // valor zero = nil (nulo)

  var x,y *int
  entero := 5
  x = &entero // ampersam accede al valor del dato, no la direccion
  y = &entero

  fmt.Println(x)
  fmt.Println(y)

  // cambiamos el valor

  *x = 6  // * da acceso al valor de la memoria

  fmt.Println(*x) // esto imprime el valor
  fmt.Println(*y) // esto imprime el valor

}
