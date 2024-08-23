package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func CreateSchema(db *sql.DB) {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL
    );

    CREATE TABLE IF NOT EXISTS chatrooms (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        chatroom_id INT REFERENCES chatrooms(id) ON DELETE CASCADE,
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        content TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );
`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}
}

func ConnectDB() *sql.DB {

	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")
	DB_USER := os.Getenv("DB_USER")

	dbConn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	fmt.Println(dbConn)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
