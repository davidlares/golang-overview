package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  fmt.Println("Hello World")
  edad := 23
  // interpolacion
  fmt.Printf("My age is %d \n", edad)
  // %v -> toma el valor nativo de la variable, aunque hay especificaciones para cada tipo de dato
  name := "David"
  fmt.Printf("Your name is %v \n", name)
  // leer datos de la terminal
  var newage int
  fmt.Println("Input your new age is: ")
  fmt.Scanf("%d \n", &newage)
  fmt.Printf("Your new age is: %d", newage)

  // Alternativa
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Input your name: ")
  text, err := reader.ReadString('\n') // busca hasta que consiga este caracter
  if(err != nil){
    fmt.Println(err)
  } else {
    fmt.Println("Hola " + text)
  }
}
