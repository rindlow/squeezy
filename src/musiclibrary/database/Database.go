package database

import (
  "os"
  "fmt"
  "database/sql"
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
  _, err = conn.Exec("create table track (id integer not null primary key, fname text)")
  if err != nil {
    fmt.Printf("sql error: %s\n", err)
    return
  }
}

func AddTracks(tracks []string) {
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
  stmt, err := tx.Prepare("insert into track(id, fname) values(?, ?)")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer stmt.Close()

  for i, t := range tracks {
    _, err = stmt.Exec(i, t)
    if err != nil {
      fmt.Println(err)
      return 
    }
  }

tx.Commit()

}



func GetAllTracks() []string {

  conn, err:=getConn()
  if err != nil {
      fmt.Println(err)
      return nil
  }

 tracks := make([]string, 1)

  rows, err := conn.Query("select id, fname from track")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var id int
		var fname string
		rows.Scan(&id, &fname)
		tracks=append(tracks, fmt.Sprintf("%-6d - %s", id, fname))
	}
	rows.Close()

	return tracks

}

