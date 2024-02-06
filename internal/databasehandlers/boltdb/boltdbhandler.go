package main

import (
	"flag"
	"log"
	"strings"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("hydra.db", 0600, nil)
	if err != nil {
		log.Fatal("open database error")
	}
	defer db.Close()

	// define the parameters parsed in console --- return pointer type
	op := flag.String("op", "", "Add or Get operation")
	user := flag.String("u", "", "Enter username")
	pwd := flag.String("p", "", "Enter password")
	flag.Parse()

	// check the condition
	switch strings.ToUpper(*op) {
	case "ADD":
		// check the length of user and pwd input
		if len(*user) != 0 && len(*pwd) != 0 {
			err = addToVault(db, *user, *pwd)
			if err != nil {
				log.Fatal(err)
			}
		}
	
	case "GET":
		if len(*user)!=0 {
			pwd, err := getPassword(db, *user)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Found password: ", pwd)
		}
		
	}
}


func addToVault(db *bolt.DB, username, password string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PasswordVault"))
		if err!=nil{
			return err
		}

		err = b.Put([]byte(username), []byte(password))
		return err
		})
}

func getPassword(db *bolt.DB, username string) (string, error) {
	password := ""
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PasswordVault"))
		// v is the byte type
		v := b.Get([]byte(username))
		password = string(v)
		return nil
	})
	return password, err
}