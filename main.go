package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-chat/database"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type RequestHandler struct {
	db *sql.DB
}

type Chatroom struct {
	Id             int
	Name           string
	ConnectedUsers []User
	Messages       []Message
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Id         int       `db:"id"`
	UserId     int       `db:"user_id"`
	Content    string    `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
	ChatroomId int       `db:"chatroom_id"`
}

var messages = []Message{}

var chatRoom = Chatroom{
	Id:       0,
	Name:     "main",
	Messages: []Message{},
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db := database.ConnectDB()

	defer db.Close()

	requestHandler := RequestHandler{
		db: db,
	}
	database.CreateSchema(requestHandler.db)

	// Check the connection
	err = requestHandler.db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/messages", requestHandler.PostMessages)

	// msg := Message{
	// 	Content: "Content mannnnn",
	// }

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}

func subscribeToChatroom(rw http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var user User
	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	chatRoom.ConnectedUsers = append(chatRoom.ConnectedUsers, user)
}

func getMessages(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(messages)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
