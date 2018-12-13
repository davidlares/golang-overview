package main

import (
  "net/http"
)

func main(){
  // fileServer := http.FileServer(http.Dir("public"))
  // http.Handle("/", http.StripPrefix("/", fileServer)) - directorio publico

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w,r,"index.html")
  })

  http.ListenAndServe(":4000", nil)
}
