package musiclibrary

import (
  "musiclibrary/database"
)

//type Mp3File struct {
  //Path string
  //Size int
//}

type MusicLibrary struct {
 //Files []Mp3File
}

func PrintLibrary() {
  database.GetAllTracks()
}
