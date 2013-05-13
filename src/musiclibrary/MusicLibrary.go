package musiclibrary

import (
  //"path"
  "path/filepath"
  "os"
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

func UpdateLibrary(library *MusicLibrary, base string) {

  // Reset the library
  fmt.Println("Resetting the library to nil")
  os.Remove("/tmp/slim.db")

  // Setup the database
  db, err := sql.Open("sqlite3", "/tmp/slim.db")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer db.Close()

  // Create table
  _, err = db.Exec("create table track (id integer not null primary key, fname text)")
  if err != nil {
    fmt.Printf("sql error: %s\n", err)
    return
  }

  // Begin transaction
  tx, err := db.Begin()
  if err != nil {
    fmt.Println(err)
    return
  }

  // Prepare query
  stmt, err := tx.Prepare("insert into track(id, fname) values(?, ?)")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer stmt.Close()

  // Iterate the base directory
  fmt.Println("Updating library from", base);
  var i=0
  filepath.Walk(base, func(p string, f os.FileInfo, err error) error {
    _, err = stmt.Exec(i, p)
    if err != nil {
      fmt.Println(err)
      return nil
    }
    i++
    return nil
  })

  // commit transaction
  tx.Commit()
}

