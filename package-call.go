package main

import (
  "fmt"
  "./package"
)

func main(){
  // llamado del paquete custom
  fmt.Println(package.HelloPackage())
}
