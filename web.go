package main

import (
  "net/http"
  "fmt"
  //"html/template"
)

type Page struct {
  Title string
  Body []byte
}

func myHandler(w http.ResponseWriter, r *http.Request) {
  //var p = Page{Title: "testing", Body: []byte("Monkey see monkey do")}
  fmt.Fprintf(w, "Welcome to the home page!")
}

/*func (myResponseWriter) Header() Header {

}*/

/*func (myHandler) ServeHTTP(ResponseWriter, *Request) {
  template.executeTemplate(http.ResponseWriter, "Testing", *Page)
}*/

func main() {
  http.HandleFunc("/", myHandler)
  //http.HandleFunc("/", myHandler)
  http.ListenAndServe(":8080", nil)
}
