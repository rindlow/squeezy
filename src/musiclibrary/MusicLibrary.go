package musiclibrary

import (
  "musiclibrary/database"
  "musiclibrary/mp3info"
  "fmt"
)

type MusicLibrary struct {
 //Files []Mp3File
}

func GetTrackById(id int) mp3info.Mp3Info {
	track:=database.GetTrackById(id)
	return track
}

func PrintLibrary() {
  tracks:=database.GetAllTracks()

  for _, t := range tracks {
    fmt.Printf("[%s] %s - %s\n", t.Album, t.Artist, t.Name)
  }

}
