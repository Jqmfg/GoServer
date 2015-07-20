package main

import (
  "net/http"
  "io"
  "io/ioutil"
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
  page, _ := ioutil.ReadFile("test.html")
  io.WriteString(w, string(page))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

//TODO: Make input value for file
//TODO: Error handling
func mapMuxValues() map[string]func(http.ResponseWriter, *http.Request) {
  var m = make(map[string]func(http.ResponseWriter, *http.Request))
  m["/"] = hello
  return m
}

func main() {
  mux = mapMuxValues()

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
