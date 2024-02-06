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
	ID int   `bson:"id"`
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

	// define query expression
	query := bson.M{
		"security clearance": bson.M{
			"$gt": 3,
		},
		"position": bson.M{
			"$in": []string{"Mechanic", "Biologist"}
		},
	}

	var crew Crew
	// return query result to crew
	err = personnel.Find(query).All(&crew)
	log.Println("Query Results", crew)

	// names is of type []struct{Name string}
	names := []struct{
		Name string `bson:"name"`
	}{}

	// select by name
	err = personnel.Find(query).Select(bson.M{"name":1}).All(&names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(names)

	// insert
	newcr := crewMember{ID: 18, Name: "Kaya Gal", SecClearance: 4, Position: "Biologist"}
	if err := personnel.Insert(newcr); err!=nil {
		log.Fatal(err)
	}

	// update with set
	err = personnel.Update(bson.M{"id":16}, bson.M{"$set": bson.M{"position": "Engineer IV"}})
	if err != nil {
		log.Fatal(err)
	}

	// remove
	if err := personnel.Remove(bson.M{"id": 18}); err != nil {
		log.Fatal(err)
	}

	// concurrent access
	var wg sync.WaitGroup
	count, _ := personnel.Count() 
	wg.Add(count)

	for i := 1; i<= count; i++ {
		// copy will keep the original content in session
		go readId(i, session.Copy(i), &wg)
	}
	wg.Wait()

}


func readId(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup) {
	// close resources when finishing
	defer func() {
		sessionCopy.Close()
		// minus one per time
		wg.Done()
	}()
	p := sessionCopy.DB("Hydra").C("Personnel")
	cm := crewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	if err != nil {
		return
	}
	

}