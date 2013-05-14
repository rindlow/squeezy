package musiclibrary

import (
  //"path"
  "path/filepath"
  "os"
  "musiclibrary/database"
  "musiclibrary/id3"
  "musiclibrary/mp3info"
)

func UpdateLibrary(base string) {

  // Nuke the existing database
  database.ReCreate()

  // Create a slice for keeping the data in-memory during scan
  tracks := make([]mp3info.Mp3Info, 1)

  // Iterate all files
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    if(!f.IsDir()) {
      // Use the id3 parser to get mp3 metadata
      fileInfo:=id3.Parse(p)

      // Add the filesystem metadata
      fileInfo.ModTime=f.ModTime()
      fileInfo.Size=f.Size()

      // Add the info to the slice of files
      tracks = append(tracks, fileInfo)
    }
    return nil
  })

  // Persist data to disk
  database.AddTracks(tracks)
}

