package database

import (
  "os"
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = nil

// TBD: Return a tuple (DB, error)
func getConn() *sql.DB {
  var err error
  if(db == nil) {
    // Setup the database
    db, err = sql.Open("sqlite3", "/tmp/slim.db")
    if err != nil {
      fmt.Println(err)
      return nil
    }
  }
  return db
}


func ReCreate() {
  // Reset the library
  fmt.Println("Nuking existing library");
  os.Remove("/tmp/slim.db")

  db=getConn()

  // Create table
  _, err := db.Exec("create table track (id integer not null primary key, fname text)")
  if err != nil {
    fmt.Printf("sql error: %s\n", err)
    return
  }
}

func AddTracks(tracks []string) {
  db:=getConn()

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

  for i, t := range tracks {
    _, err = stmt.Exec(i, t)
    if err != nil {
      fmt.Println(err)
      return 
    }
  }

tx.Commit()

}



func GetAllTracks() {

  conn := getConn()
  rows, err := conn.Query("select id, fname from track")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var id int
		var fname string
		rows.Scan(&id, &fname)
		fmt.Println(id, fname)
	}
	rows.Close()

}

