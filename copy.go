package main

import "fmt"

func main(){
  // copia el minimo de elementos dados en la estructura
  slice := []int {1,23,4}
  copia := make([]int, 4)
  copy(copia, slice) // slice es la fuente, copia es a donde se van a pasar los archivos
  fmt.Println(slice)
  fmt.Println(copia)
}
