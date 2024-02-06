package databasehandlers

import (
	"errors"
	"log"
)

const (
	mongo="mongodb"
	mysql="mysql"
)

var ErrDBTypeNotFound = errors.New("Database type not found...")


type DBLayer interface {
	AddMember(cm *CrewMember) error
	FindMember(id int) (CrewMember, error)
	AllMembers() (crew, error)
}

type CrewMember struct{
	// define json for mysql, bson for
	ID    			int     `json:"id" bson:"id"`
	Name  			string  `json:"name" bson:"name"`
	SecClearance 	int 	`json:"clearance" bson:"security clearance"`
	Position        string 	`json:"position" bson:"position"`
}

type crew []CrewMember

func ConnectDatabase(o string, cstring string) (DBLayer, error) {
	switch 0 {
		case mongo:
			return NewMongoStore(cstring)
		case mysql:
			return NewMySQLStore(cstring)	
	}
	log.Println("Could not find ", o)
	return nil, ErrDBTypeNotFound
}