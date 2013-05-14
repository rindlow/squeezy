package id3

import (
  "os"
  "fmt"
  "github.com/ascherkus/go-id3/src/id3"
)

func Parse(fname string) string {
  fd, _ := os.Open(fname)
  defer fd.Close()
  file := id3.Read(fd)

  ret:="<unknown>"
  if file != nil {
    ret=fmt.Sprintf("[%s] %s - %s (%s)", file.Album, file.Artist, file.Name, file.Year)
  }
  return ret
}


