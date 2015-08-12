package logging

import (
  "errors"
  "fmt"
  "os"
  "io"
  //"syscall"
)

//TODO: Decide if should use a logger struct

func check(e error) {
  if(e != nil) {
    panic(e)
  }
}

//TODO: Make functions for checking and opening files

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
  file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0666)

  //TODO: Include timestamp
  //TODO: Include ip?
  //NOTE: First return value returns number of bytes writen
  //TODO: Additional info for 404
  //TODO: Rework to give better info
  _, err = io.WriteString(file, requestedURL + " was requested | " + accessedURL + " was accessed\n")
  //TODO: Separate file for error logging
  if(err != nil) {
    fmt.Println(err)
    return err
  }
  file.Close()
  return errors.New("nil")
}

func logErrors(toWrite string, logFile string) error {
  _, err := os.Stat(logFile)
  if(err != nil) {
    fcreated, err :=os.Create(logFile)
    fcreated.Close()
    fmt.Println(err)
  }
  //TODO: Check syscall.O_RDWR
  file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0666)

  //TODO: Include timestamp
  //NOTE: First return value returns number of bytes writen
  _, err = io.WriteString(file, toWrite + "\n")

  if(err != nil) {
    fmt.Println("Error writing to log file!")
    fmt.Println(err)
    return err
  }
  return err
}
