package main

import "fmt"

func main(){
  // arreglo := [3]int {1,2,3}
  arreglo := make([]int, 3, 5) // tipo de slice a crear (puntero), la longitud, el capacity
  slice := arreglo[0:]

  slice = append(slice,2) // agregar elementos con append, sin necesidad de reconstruir todo el slice

  fmt.Println(slice)
  fmt.Println(len(slice)) // len -> longitud en elementos
  fmt.Println(cap(slice)) // capacity, cuanto cabe en el slice
}
