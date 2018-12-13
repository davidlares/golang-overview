package main

import (
  "net/http"
  "io"
  "fmt"
)

func main(){
  http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request){
      io.WriteString(w,"Hello!")
  })

  http.HandleFunc("/", handler)
  http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter,r *http.Request){
  fmt.Println("new Request")
  io.WriteString(w, "Hello, World")
}
