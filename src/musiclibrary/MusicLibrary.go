package musiclibrary

import (
  "musiclibrary/database"
  "fmt"
)

type MusicLibrary struct {
 //Files []Mp3File
}

func PrintLibrary() {
  tracks:=database.GetAllTracks()

  for _, t := range tracks {
    fmt.Printf("[%s] %s - %s\n", t.Album, t.Artist, t.Name)
  }

}
