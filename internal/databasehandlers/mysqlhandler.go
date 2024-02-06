package databasehandlers

import (
	"database/sql"
	// make sure the init function is executed
	_ "github.com/go-sql-driver/mysql"
)

// define the db struct
type mysqlDataStore struct {
	*sql.DB
}

// define the connection part
func NewMySQLStore(conn string) (*mysqlDataStore, error){
	/*
	:param conn: define the parameters for connection
	*/
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &mysqlDataStore{DB: db}, nil
}


// define the addmember function
func (msql *mysqlDataStore) AddMember(cm *CrewMember) error {
	_, err := msql.Exec("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES (?,?,?)", cm.Name, cm.SecClearance, cm.Position)
	return err
}


// define the findmember function
func (msql *mysqlDataStore) FindMember(id int) (CrewMember, error){
	row := msql.QueryRow("Select * from Personnel where id=?", id)
	cm := CrewMember{}
	err := row.Scan(&cm.ID, &cm.Name, &cm.SecClearance, &cm.Position)
	return cm, err
}


// define the function to return all members
func (msql *mysqlDataStore) AllMembers() (crew, error) {
	rows, err := msql.Query("Select * from Personnel")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members := crew{}
	for rows.Next() {
		member := CrewMember{}
		err = rows.Scan(&member.ID, &member.Name, &member.SecClearance, &member.Position)
		if err != nil {
			members = append(members, member)
		}

	}

	err = rows.Err()
	return members, err

}
