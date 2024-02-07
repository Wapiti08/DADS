package databasehandlers

import (
	"gopkg.in/mgo.v2"
)

type mongoDataStore struct {
	*mgo.Session
}

func NewMongoStore(conn string) (*mongoDataStore, error) {
	// conn is generally an URL
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return &mongoDataStore{Session: session}, nil
}

func (ms *mongoDataStore) AddMember() error {

}

func (ms *mongoDataStore) FindMember() () {

}


