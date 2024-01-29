package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// not call methods from the package directly

	_ "github.com/go-sql-driver/mysql"
)

type crewMember struct {
	id int
	name string
	secClearance int
	position string
}

type Crew []crewMember

func main() {
	// user:passwd@params
	db, err := sql.Open("mysql", "")

	if err != nil {
		log.Fatal("Could not connect, error ",err.Error())
	}

	defer db.Close()

	// get value by crew position
	cw := GetCrewByPosition(db, []string{"'Mechanic'","'Biologist'"})

	fmt.Println(cw)
	
}


func GetCrewByPosition(db *sql.DB, positions []string) Crew {
	Qs := fmt.Sprintf("SELECT id, Name, SecurityClearance, Position from Personnel where Position in (%s);", strings.Join(positions, ","))

	rows, err := db.Query(Qs)
	if err != nil {
		log.Fatal("could not get data from the Personal table", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	// retrive data
	for rows.Next() {
		member := crewMember{}
		// copies the columns values with pointer type
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err != nil {
			log.Fatal("Erro scanning row", err)
		}
		retVal = append(retVal, member)

	}

	// check the rows err
	if err := rows.Err(); err!= nil {
		log.Fatal(err)
	}
	return retVal

}

func (cm crewMember) GetCrewById(db *sql.DB, id int)  {
	// question mark is the variable wildcard
	row := db.QueryRow("Select * from Personnel where id=?",id)
	
	err := row.Scan(&cm.id, &cm.name, &cm.secClearance, &cm.position)

	if err != nil {
		log.Fatal(err)
	}

	return
}

func AddCrewMember(db *sql.DB, cm crewMember) int64 {
	res, err := db.Exec("INSERT INTO Personnel (Nmae,SecurityClearance,Position) VALUES (?,?,?)", cm.name, cm.secClearance, cm.position)
	if err != nil {
		log.Fatal(err)
	}

	ra, _ := res.RowsAffected()
	re, _ := res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted id", re)
	return re
}

