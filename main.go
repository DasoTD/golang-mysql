package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//_ "github.com/lib/pq"

func main() {
	// message := fmt.Sprintf("Hi, Welcome!")
	// return messagess
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/mongoAPI?parseTime=true")
	if err != nil {
		panic(err)
	}
	//See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	db.Ping()
	fmt.Println("DB CONNECTED")

	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	db.Exec(query)

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	result.LastInsertId()
	// userID, err := result.LastInsertId()
	// return userID

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			panic(err)
		}

		fmt.Println(id, username, password, createdAt)
	}
}
