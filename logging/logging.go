package logging

import (
  "errors"
  "fmt"
  "os"
  "io"
)

func check(e error) {
  if(e != nil) {
    panic(e)
  }
}

func LogWebPath(requestedURL string, accessedURL string, logFile string) error {
  //TODO: Ensure file and directory exist

  file, err := os.Create(logFile)

  //TODO: Separate file for error logging
  if(err != nil) {
    fmt.Println(err)
    return err
  }

  //TODO: Include timestamp
  //TODO: Include ip?
  //NOTE: First return value returns number of bytes writen
  _, err = io.WriteString(file, requestedURL + " was requested | " + accessedURL + " was accessed")
  //TODO: Separate file for error logging
  if(err != nil) {
    fmt.Println(err)
    return err
  }
  file.Close()
  return errors.New("nil")
}
