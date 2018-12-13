package main

import "fmt"

func main(){
  // declaracion -> tienen valores 0, falsos o string vacios
  var array [3] int
  array2 := [3] int {1,2,3}
  fmt.Println(array)
  fmt.Println(array2)

  // metodos
  fmt.Println(len(array2)) // length

  // recorrido
  array2[2] = 20
  for i := 0; i < len(array2); i++ {
    fmt.Println(array[i])
  }

  // matrices
  var mat [2][2] int
  fmt.Println(mat)
  mat[0][1] = 10
  fmt.Println(mat[0][1])
}
