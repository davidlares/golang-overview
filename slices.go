package main

import "fmt"

func main(){
  // slices -> especie de arrays redimensionables
  slice := []int {1,2,3,4}
  fmt.Println(slice)

  // no inicializado
  var mat []int
  if mat == nil {
    fmt.Println("Esta vacio")
  } else {
    // tiene los mismos metodos que los arrays
    fmt.Println(len(mat))
  }

  // estructura slice:

  // longitud array
  // puntero al array
  // capacidad

  // SLICING -> partimos un arraglo para generar un slice
  arreglo := [3]int {11,22,33}
  slice2 := arreglo[:2]
  fmt.Println(slice2)
}
