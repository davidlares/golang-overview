package main

import (
  "net/http"
  "encoding/json"
)

type Course struct {
  // si estan en minuscula los variables, son privadas, y no se van a servir en el json
  Title string  `json: "title"`
  videos int `json: "videos"`
}

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    course := Course{"Course Title", 15}
    json.NewEncoder(w).Encode(course)
  })

  http.ListenAndServe(":4000", nil)
}
