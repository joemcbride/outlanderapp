package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
)

type Page struct {
    Title string
    Version Version 
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/version", version)

    dt, _ := ISO8601FullFromNow()

    fmt.Printf("Listening at %s\n", dt)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func home(res http.ResponseWriter, req *http.Request) {
    p1 := &Page{Version: version_data()}
    t, _ := template.ParseFiles("home.html")
    t.Execute(res,p1);
}

func version(rw http.ResponseWriter, req *http.Request) {
        rw.Header().Set("Content-Type", "application/json")
        fmt.Fprint(rw, version_response(version_data()))
        return
}