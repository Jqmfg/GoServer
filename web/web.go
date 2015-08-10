package main

import (
  "net/http"
  "io"
  "io/ioutil"
  "bufio"
  "os"
  "strings"
  "log"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct {
  requestRouter map[string]string
}

func returnMuxFunc(fileName string) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    page, _ :=ioutil.ReadFile(fileName)
    io.WriteString(w, string(page))
  }
}

func createRequestRouter(fileName string) map[string]string {
  r := make(map[string]string)

  //TODO: Error Handling
  file, _ := os.Open(fileName)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    curLine := strings.Split(scanner.Text(), ":")
    r[curLine[0]] = curLine[1]
  }
  return r
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

  //NOTE: Currently must specify the full path of additional files in html files (e.g. /templates/style.css)
  //TODO: Correct to not have to specify full path of additional files in html files
  if(h.requestRouter[path] != "") {
    path = h.requestRouter[path]
  }
  //TODO: Make a log file
  log.Println(r.URL.Path + ": accessing " + path)

  //TODO: Error handling
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

func main() {
  //TODO: Make global variables for all of these
  requestRouter := createRequestRouter("templates/map.txt")

  s := http.Server {
    Addr: ":8080",
    Handler: &myHandler{requestRouter},
  }

  s.ListenAndServe()
}
