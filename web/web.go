package main

import (
  "net/http"
  "io"
  "io/ioutil"
  //"html/template"
)

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
	//io.WriteString(w, "Hello world!")
  page, _ := ioutil.ReadFile("test.html")
  io.WriteString(w, string(page))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func mapMuxValues(m map[string]func(http.ResponseWriter, *http.Request)) {

}

func main() {
  mux = make(map[string]func(http.ResponseWriter, *http.Request))
  mux["/"] = hello

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
