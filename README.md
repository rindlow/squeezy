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

  go get cgl.tideland.biz/state


TODO
====

The package slimserver is getting overloaded, should probably be refactored into parts:
* slimproto
* disco
* streamer
* eventhandler

Should be more or less trivial to do.

FSM TERMINOLOGY
===============

External modules can create EVENTS (user pressed button, track ended, connection lost), these are sent to the FSM. The FSM processes the EVENT and as a result it might change its STATE and also emit an ACTION. The actions are sent to the different modules which might react (sending a message to a player, stop streaming, show text on device display).

The channels should be configured one way:
* Event channels are always module ===> FSM
* Action channels are always FSM ===> module



