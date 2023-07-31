package webbasics

import (
	"database/sql" // provides a generic interface around sql
	"fmt"
	"time"

	// To connect with database
	//import _ import and store in _, prevent warnings
	_ "github.com/go-sql-driver/mysql"
)

type  portalUser struct{
	id int
	username string
	password string
	createdAt time.Time
}



func createDatabase(db *sql.DB) {
	query := `
		CREATE TABLE portaluser (
			id int AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id));`
	
		res , err :=	db.Exec(query) // Execute query w/o rows
		check(err)

		rows, _ := res.RowsAffected()
		fmt.Printf("rows affected: %#v", rows)
}

func insertToTable(db *sql.DB, user *portalUser){
	
	// Note: Similar query can be used for delete
	query := `INSERT INTO portaluser 
				(username, password, created_at) 
				VALUES (?,?,?)`

	res, err := db.Exec(query, user.username, user.password, user.createdAt)
	check(err)
	createdId, err := res.LastInsertId()
	check(err)
	fmt.Printf("Inserted ID of newly added row: %v\n ", createdId)
	
}


/*
	go sql can scan results to variables of different types from standard go 
	as well as from go.sql package. 

	e.g. sql.NullFloat64 (nullable float 64)
*/
func getRowFromTable(db *sql.DB, filter string){
	
	query := `SELECT * FROM portaluser WHERE id = ?`
	user := portalUser{}

	// Query that returns atmost single result
	row := db.QueryRow(query, filter)
	// copies the result to the destination
	err := row.Scan(&user.id, &user.username, &user.password, &user.createdAt) 
	check(err)
	
	fmt.Printf("Queried user: %#v", user)
}


func getAllRowsFromTable(db *sql.DB) {
	query := `SELECT id, username, password, created_at from portaluser`

	rows, err := db.Query(query)
	check(err)
	defer rows.Close() // close the row preventing further enumeration
	

	var users []portalUser

	// Result vs ResultSet
	for rows.Next() {
		var u portalUser
		rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		users = append(users, u)
	}

	fmt.Printf("Users : %#v", users)

	err = rows.Err() // check err
	check(err)
}


func ConnectSQL(){
	db,err := sql.Open("mysql", 
						"aj:hello-world@(127.0.0.1:3306)/testdb?parseTime=true")
	
	check(err)
	err = db.Ping()
	check(err)

	// createDatabase(db)
	
	// insertToTable(db,&portalUser{
	// 						username: "Sam Doe", 
	// 						password: "password", 
	// 						createdAt: time.Now(), 
	// 					})

	// getRowFromTable(db, "1")
	// getAllRowsFromTable(db)
}


