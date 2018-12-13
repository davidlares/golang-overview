package main

import (
  "fmt"
  "io/ioutil"
  "bufio"
  "os"
)

func main(){
  // la desventaja es la carga en la memoria del archivo -> ideal: archivos pequenos
    data,err := ioutil.ReadFile("./helloworld.txt")
    if err != nil{
      fmt.Println("Error")
    }
    fmt.Println(string(data)) // convertimos bytes a string

    file,err := os.Open("./helloworld.txt")
    if err != nil {
      fmt.Println("Error")
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        fmt.Println(line)
    }

    file.Close()
}
