package database

import (
  "os"
  "fmt"
  "database/sql"
  "musiclibrary/mp3info"
  _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

func getConn() (*sql.DB,error) {
  var err error
  if(db == nil) {
    // Setup the database
    db, err = sql.Open("sqlite3", "/tmp/slim.db")
    if err != nil {
      return nil, err
    }
  }
  return db, nil
}


func ReCreate() {
  // Reset the library
  fmt.Println("Nuking existing library");
  os.Remove("/tmp/slim.db")

  conn, err:=getConn()
  if err != nil {
      fmt.Println(err)
      return
  }

  // Create table
  _, err = conn.Exec("create table track (id integer not null primary key, fname text, name text, album text, artist text)")
  if err != nil {
    fmt.Printf("sql error: %s\n", err)
    return
  }
}

func AddTracks(tracks []mp3info.Mp3Info) {
  conn, err:=getConn()
  if err != nil {
      fmt.Println(err)
      return
  }

  // Begin transaction
  tx, err := conn.Begin()
  if err != nil {
    fmt.Println(err)
    return
  }

  // Prepare query
  stmt, err := tx.Prepare("insert into track(id, fname, name, album, artist) values(?, ?, ?, ?, ?)")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer stmt.Close()

  for i, t := range tracks {
    _, err = stmt.Exec(i, t.FName, t.Name, t.Album, t.Artist)
    if err != nil {
      fmt.Println(err)
      return 
    }
  }

tx.Commit()

}



func GetAllTracks() []mp3info.Mp3Info {

  conn, err:=getConn()
  if err != nil {
      fmt.Println(err)
      return nil
  }

 tracks := make([]mp3info.Mp3Info, 1)

  rows, err := conn.Query("select id, fname, name, album, artist from track")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var id int
		var info mp3info.Mp3Info
		rows.Scan(&id, &info.FName, &info.Name, &info.Album, &info.Artist)
		tracks=append(tracks, info)
	}
	rows.Close()

	return tracks

}

