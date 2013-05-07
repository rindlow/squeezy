package slimserver

import (
  "path"
  "path/filepath"
  "os"
  "fmt"
)

// TBD: Extend the library with actual information
type MusicLibrary struct {
    NumDirs int
    Albums map[string] int
}

func UpdateLibrary(library *MusicLibrary, base string) {

  // Reset the library
  fmt.Println("Resetting the library to nil")
  library.NumDirs=0
  library.Albums = make(map[string]int)

  // Iterate the directory
  fmt.Println("Updating library from", base);
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    library.NumDirs++
    library.Albums[path.Dir(p)]++
    return nil
  })
}

