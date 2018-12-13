package main

import (
  "fmt"
)

func main(){
   // canales de escucha de go routines
   channel := make(chan string)
   // go routine
   go func(channel chan string){
     for {
       var name string
       fmt.Scanln(&name)
       // enviar informacion (orden)
       // canal <- dato a enviar
       channel <- name
     }
   }(channel) // canal de escucha

   // recibimos informacion (orden inverso)
   msg := <- channel
   fmt.Println(msg)
}
