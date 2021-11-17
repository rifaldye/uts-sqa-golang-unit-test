package config

import (
	"time"

	"gopkg.in/mgo.v2"
)


func GetMongoDB() (*mgo.Database, error) {
	host := "localhost"
	dbName := "uts"
	session, err := mgo.DialWithTimeout(host,10*time.Second)
	session.SetPoolLimit(10)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}