package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", home)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func home(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "Outlander App")
}