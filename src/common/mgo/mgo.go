package mgo


import (
	"gopkg.in/mgo.v2"
	"log"
)



func NewMgo(s string) *mgo.Session {
	m, err := mgo.Dial(s)
	if err != nil {
		log.Println(err)
	}
	//defer Mgo.Close()
	// Optional. Switch the session to a monotonic behavior.
	m.SetMode(mgo.Monotonic, true)

	return m
}