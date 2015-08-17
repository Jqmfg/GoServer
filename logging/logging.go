package logging

import (
  "os"
  "io"
  "net/http"
  "time"
  //"syscall"
)

//TODO: Make the log files come from global config file

func writeToFile(logFile string, toWrite string) error {
  //NOTE: First parameter is FileInfo
  _, err := os.Stat(logFile)
  if(err != nil) {
    //TODO: Error checking
    fcreated, _ := os.Create(logFile)
    fcreated.Close()
    //TODO: log that the file was created
  }
  //TODO: Check if this needs error checking
  //TODO: Check syscall.O_RDWR
  //NOTE: os.Create() returns type File
  file, _ := os.OpenFile(logFile, os.O_APPEND | os.O_WRONLY, 0666)

  //TODO: Include ip?
  //TODO: Additional info for 404
  //TODO: Rework to give better info
  //TODO: Check if need to append a newline
  //NOTE: First return values returns number of bytes written
  _, err = io.WriteString(file, "TIMESTAMP: " + time.Now().UTC().Format("2006-01-02 15:04:05 (UTC)") + " | " + toWrite + "\n")

  if(err != nil) {
    logErrors("log/errors.log", "Error writing file: " + err.Error())
    return err
  }
  file.Close()
  return err
}

func logErrors(logFile string, toWrite string) error {
  //NOTE: First return value returns number of bytes writen
  err := writeToFile(logFile, toWrite)
  //TODO: Check for using: return errors.New("nil")
  return err
}

//TODO: Make functions for checking and opening files

//TODO: Condense arguments
func LogWebPath(logFile string, requestedURL string, accessedURL string, r *http.Request) error {
  err := writeToFile(logFile, r.RemoteAddr + " | " + requestedURL + " was requested | " + accessedURL + " was accessed")
  //TODO: Separate file for error logging
  return err
}

func StartServer(server http.Server) {

}
