package musiclibrary

import (
  //"path"
  "path/filepath"
  "os"
  "musiclibrary/database"
)

func UpdateLibrary(base string) {

  // Nuke the existing database
  database.ReCreate()

  // Create a slice for keeping the data in-memory during scan
  tracks := make([]string, 1)

  // Iterate all files
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    tracks = append(tracks, p)
    return nil
  })

  // Persist data to disk
  database.AddTracks(tracks)
}

