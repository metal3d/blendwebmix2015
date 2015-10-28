package data

import "gopkg.in/mgo.v2"

func getMongo() *mgo.Session {
	m, _ := mgo.Dial("127.0.0.1")
	return m
}
