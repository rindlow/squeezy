squeezy
=======

An attempt to write a squeezebox server in Go.


Dependencies
============

This app uses libraries for SQLite and ID3, first install prereqs:
 apt-get install pkg-config
 apt-get install libsqlite3-dev

Fetch the source (ignore the warnings for go-id3):
  go get -d github.com/ascherkus/go-id3/
  go get -d github.com/mattn/go-sqlite3/

