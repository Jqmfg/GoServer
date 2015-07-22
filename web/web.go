package main

import (
  "net/http"
  "io"
  "io/ioutil"
  "bufio"
  "os"
  "strings"
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
//TODO: Error handling
func returnMuxFunc(fileName string) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    page, _ :=ioutil.ReadFile(fileName)
    io.WriteString(w, string(page))
  }
}
//TODO: Error handling
func createMux(fileName string) map[string]func(http.ResponseWriter, *http.Request) {
  m := make(map[string]func(http.ResponseWriter, *http.Request))

  file, _ := os.Open(fileName)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    curLine := strings.Split(scanner.Text(), ":")
    m[curLine[0]] = returnMuxFunc(curLine[1])
  }
  fs := http.FileServer(http.Dir("../static"))
  m["/"] = fs
  return m
}

//TODO: Implement fileserver for css ↓
//https://groups.google.com/forum/#!topic/golang-nuts/aGMLK_2OHiM
func main() {
  fs := http.FileServer(http.Dir("../static"))

  mux = createMux("../files/map.txt")

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
