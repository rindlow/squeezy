package id3

import (
  "os"
  "musiclibrary/mp3info"
  "github.com/ascherkus/go-id3/src/id3"
)

func Parse(fname string) mp3info.Mp3Info {
  fd, _ := os.Open(fname)
  defer fd.Close()
  file:=id3.Read(fd)

  // Wrap the metadata so that we do not leak internal types
  var ret mp3info.Mp3Info

  ret.FName=fname
  if(file!=nil) {
    ret.Name=file.Name
    ret.Artist=file.Artist
    ret.Album=file.Album
    ret.Year=file.Year
    ret.Track=file.Track
    ret.Length=file.Length
  }

  return ret
}


