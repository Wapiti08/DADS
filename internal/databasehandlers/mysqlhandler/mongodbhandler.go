package main

import {
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
}

type crewMember strcut {
	ID int   'bson:"id"'
	Name string `bson:"name"`
	SecClearance int `bson:"security clearance"`
	Position string `bson:"position"`
}

type Crew []crewMember

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()
	// database --- connection
	personnel := session.DB("Hydra").C("Personnel")

	// get number of documents in the collection
	n, _ := personnel.Count()
	log.Println("Number of personnel is ", n)

	cm := crewMember{}
	personnel.Find(bson.M{"id":3}).One(&cm)
	log.Println(cm)


}