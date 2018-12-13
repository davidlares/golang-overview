package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
    execution := readFile()
    fmt.Println(execution)
}

func readFile() bool{
  file,err := os.Open("./helloworld.txt")

  defer func() {
    file.Close() // agrega la instruccin de codigo un stack en particular -> condicional a cuando retorne la funcion sin importa el punto
    fmt.Println("Defer")
  }()

  if err != nil {
    fmt.Println("Error")
  }

  scanner := bufio.NewScanner(file)
  var i int
  for scanner.Scan(){
    i++
      line := scanner.Text()
      fmt.Println(i)
      fmt.Println(line)
  }

  if true {
    return true
  }
  fmt.Println("No me ejecuto")
  file.Close()
  return true
}
