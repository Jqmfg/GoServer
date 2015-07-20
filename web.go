package main

import (
  "net/http"
  "time"
  "io"
  //"html/template"
)

type Page struct {
  Title string
  Body []byte
}

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello World!")
}

type myHandler struct {

}

var mux map[string]func(http.ResponseWriter, *http.Request)

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if h, ok := mux[r.URL.String()]; ok {
    h(w, r)
    return
  }
  io.WriteString(w, "My server: " + r.URL.String())
}

func main() {
  mux := make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = hello

  s := &http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
    ReadTimeout: 10 * time.Second,
    WriteTimeout: 10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  s.ListenAndServe()
}
