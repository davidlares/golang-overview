package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  executeReadFile()
  fmt.Println("Esto no se ejecuta, al menos que exista recover") // si se ejecutan todas las acciones del stack se cierra el programa
}

func executeReadFile(){
    execution := readFile()
    fmt.Println(execution)
}

func readFile() bool{
  file,err := os.Open("./helloworlda.txt")

  defer func() {
    file.Close() // agrega la instruccin de codigo un stack en particular -> condicional a cuando retorne la funcion sin importa el punto
    fmt.Println("Defer")
    recover() // esto detiene el panic
  }()

  if err != nil {
    panic(err) // forzamos el panic -> forma para imprimir el stack del error
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
