package musiclibrary

import (
  //"path"
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

//type Mp3File struct {
  //Path string
  //Size int
//}

type MusicLibrary struct {
 //Files []Mp3File
}

func PrintLibrary() {
  db, err := sql.Open("sqlite3", "/tmp/slim.db")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer db.Close()

rows, err := db.Query("select id, fname from track")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var fname string
		rows.Scan(&id, &fname)
		fmt.Println(id, fname)
	}
	rows.Close()

}

