package main

import "fmt"

func main(){
  x := 10
  y := 5

  /* ==  != < > >= <= && || */

  if x >= y {
    fmt.Println("%d es mayor que %d \n",x,y)
  } else {
    fmt.Println("%d es menor que %d \n",x,y)
  }

  // no se pueden iniciar las llaves del condicional en la linea inferior de la declaracion
  // si el if tiene una linea, se suele no colocar la llaves, en Go se tienen que poner las llaves
}
