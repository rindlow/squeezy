package musiclibrary

import (
  //"path"
  "path/filepath"
  "os"
  "musiclibrary/database"
  "musiclibrary/id3"
)

func UpdateLibrary(base string) {

  // Nuke the existing database
  database.ReCreate()

  // Create a slice for keeping the data in-memory during scan
  tracks := make([]string, 1)

  // Iterate all files
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    foo:=id3.Parse(p)
    tracks = append(tracks, foo)
    return nil
  })

  // Persist data to disk
  database.AddTracks(tracks)
}

