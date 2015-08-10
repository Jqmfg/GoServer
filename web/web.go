package main

import (
  "net/http"
  "io"
  "io/ioutil"
  //"bufio"
  //"os"
  "strings"
  "log"
)

type myHandler struct {

}

//TODO: Make files based on file pathname rather than mux
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
  log.Println(path)

  data, err := ioutil.ReadFile(string(path))

  if err == nil {
    var contentType string

    switch {
      case strings.HasSuffix(path, ".css"):
        contentType = "text/css"
      case strings.HasSuffix(path, ".html"):
        contentType = "text/html"
      case strings.HasSuffix(path, ".png"):
        contentType = "image/png"
      case strings.HasSuffix(path, ".js"):
        contentType = "application/javascript"
      case strings.HasSuffix(path, ".svg"):
        contentType = "image/svg+xml"
      default:
        contentType = "text/plain"
    }

    w.Header().Add("Content Type", contentType)
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
//TODO: Add defaults through the mux

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
func createMux(fileName string) map[string]func(http.ResponseWriter, *http.Request) {
  m := make(map[string]func(http.ResponseWriter, *http.Request))

  file, _ := os.Open(fileName)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    curLine := strings.Split(scanner.Text(), ":")
    m[curLine[0]] = returnMuxFunc(curLine[1])
  }
  tempHandler := http.StripPrefix("/templates/", http.FileServer(http.Dir("/templates")))
  m["templates"] = tempHandler.ServeHTTP
  return m
}
*/

//TODO: Implement fileserver for css ↓
//https://groups.google.com/forum/#!topic/golang-nuts/aGMLK_2OHiM
func main() {
  //TODO: Make global variables for all of these
  //mux = createMux("templates/map.txt")

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{},
  }

  s.ListenAndServe()
}
