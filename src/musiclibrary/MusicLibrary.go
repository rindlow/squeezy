package musiclibrary

import (
  "musiclibrary/database"
  "fmt"
)

//type Mp3File struct {
  //Path string
  //Size int
//}

type MusicLibrary struct {
 //Files []Mp3File
}

func PrintLibrary() {
  tracks:=database.GetAllTracks()

  for _, t := range tracks {
    fmt.Println(t)
  }

}
