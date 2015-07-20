package main

import (
  "net/http"
  "io"
  //"html/template"
)

/*type Page struct {
  Title string
  Body []byte
}*/

type myHandler struct {

}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "My server: "+r.URL.String())
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
  mux = make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = hello

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
