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
  //TODO: Separate file for error logging
  //NOTE: First parameter is FileInfo
  _, err := os.Stat(logFile)
  if(err != nil) {
    //TODO: Separate file for error logging
    //NOTE: os.Create returns type File
    fcreated, err := os.Create(logFile)
    fcreated.Close()
    fmt.Println(err)
  }

  //TODO: Separate file for error logging
  //NOTE: The file should exist here
  file, _ := os.Open(logFile)

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
