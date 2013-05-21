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

  go get github.com/op/go-logging

TODO
====

The package slimserver is getting overloaded, should probably be refactored into parts:
* slimproto
* disco
* streamer
* eventhandler

Should be more or less trivial to do.
