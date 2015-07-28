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
	path := r.URL.Path[1:]
  log.Println(path)

  data, err := ioutil.ReadFile(string(path))

  if err == nil {
    w.Write(data)
  } else {
    w.WriteHeader(404)
    w.Write([]byte("404 error find another page. " + http.StatusText(404)))
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
  page, _ := ioutil.ReadFile("templates/test.html")
  io.WriteString(w, string(page))
}

var mux map[string]func(http.ResponseWriter, *http.Request)
//TODO: Error handling

//old
/*
func returnMuxFunc(fileName string) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    page, _ :=ioutil.ReadFile(fileName)
    io.WriteString(w, string(page))
  }
}
*/


//old code for using a mux
/*
//todo: Error handling
func createMux(fileName string) map[string]func(http.ResponseWriter, *http.Request) {
  m := make(map[string]func(http.ResponseWriter, *http.Request))

  file, _ := os.Open(fileName)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    curLine := strings.Split(scanner.Text(), ":")
    m[curLine[0]] = returnMuxFunc(curLine[1])
  }
  //todo: Switch to looking for .css files rather than in templates
  tempHandler := http.StripPrefix("/templates/", http.FileServer(http.Dir("/templates")))
  m["templates"] = tempHandler.ServeHTTP
  return m
}
*/

//TODO: Implement fileserver for css ↓
//https://groups.google.com/forum/#!topic/golang-nuts/aGMLK_2OHiM
func main() {
  //TODO: Make global variables for all of these
  mux = createMux("templates/map.txt")

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
