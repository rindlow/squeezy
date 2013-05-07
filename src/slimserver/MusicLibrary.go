package slimserver

import (
  //"path"
  "path/filepath"
  "os"
  "fmt"
)

type Mp3File struct {
  Path string
  Size int
}

type MusicLibrary struct {
 Files []Mp3File
}

func UpdateLibrary(library *MusicLibrary, base string) {

  // Reset the library
  fmt.Println("Resetting the library to nil")
  library.Files = make([]Mp3File, 0, 25000)

  // Iterate the base directory
  fmt.Println("Updating library from", base);
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    var mp3file Mp3File
    mp3file.Path = p
    mp3file.Size = 4711
    library.Files=append(library.Files, mp3file)
    return nil
  })
}

