package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"github.com/chrisjstevenson/go-gin-mgo/conf"
)

var (
	Session *mgo.Session

	Mongo *mgo.DialInfo
)


func Connect() {
	var uri = conf.Config()["mongourl"]
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
