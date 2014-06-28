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

    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        reqDt, _ := ISO8601FullFromNow()
        fmt.Printf("%s Requesting file: %s\n", reqDt, r.URL.Path[1:])
        http.ServeFile(w, r, r.URL.Path[1:])
    })

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

    query := req.URL.Query()
    version := version_data()

    dt, _ := ISO8601FullFromNow()

    fmt.Printf("%s %s\n", dt, query.Get("version"))

    rw.Header().Set("Content-Type", "application/json")

    if version.Version == query.Get("version") {
        rw.WriteHeader(204)
        return
    }

    fmt.Fprint(rw, version_response(version))
    return
}