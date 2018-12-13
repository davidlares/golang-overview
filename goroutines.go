package main

import (
  "fmt"
  "time"
  "strings"
)

func main(){
  // concurrencia -> dividir un problema en diferences ejecuciones de forma sincrona

  go nombre_lento("david") // go keyword -> subrutina
  fmt.Println("Second message") // esta segunda linea se ejecuta sincronamente
  // esto es una forma de mantener la ejecucion viva (bloquear information)
  var wait string
  fmt.Scanln(&wait)
}

func nombre_lento(name string){
  letters := strings.Split(name,"")
  for _,letter := range(letters){
    time.Sleep(1000 * time.Millisecond)
    fmt.Println(letter)
  }
}
