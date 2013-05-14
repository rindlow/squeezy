package mp3info

import (
  "time"
)

// The only purpose of this package is to provide the Mp3Info type
type Mp3Info struct {
        FName    string
	Size     int64        // TBD: Persist the value
	ModTime  time.Time    // TBD: Persist the value
        Name     string
        Artist   string
        Album    string
        Year     string
        Track    string
        Length   string
}

