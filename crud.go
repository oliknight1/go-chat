package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateUser(db *sql.DB, username string) (int, error) {
	var rowId int
	err := db.QueryRow("INSERT INTO users (username) VALUES ($1) RETURNING id", username).Scan(&rowId)
	return rowId, err
}

func CreateChatroom(db *sql.DB, name string) (int, error) {
	var rowId int
	err := db.QueryRow("INSERT INTO chatroom (name) VALUES ($1) RETURNING id", name).Scan(&rowId)
	return rowId, err
}

func CreateMessage(db *sql.DB, msg Message) (int, error) {
	var rowId int
	err := db.QueryRow("INSERT INTO messages (chatroom_id,userId,content) VALUES ($1, $2, $3) RETURNING id",
		msg.ChatroomId, msg.UserId, msg.Content).Scan(&rowId)
	return rowId, err
}

// type Message struct {
// 	MessageId  int       `db:"message_id"`
// 	UserId     int       `db:"user_id"`
// 	Content    string    `db:"content"`
// 	CreatedAt  time.Time `db:"created_at"`
// 	ChatroomId int       `db:"chatroom_id"`
// }
func GetMessages(db *sql.DB, chatroomId int) ([]Message, error) {
	rows, err := db.Query("SELECT content, user_id, created_at, chatroom_id FROM messages WHERE chatroom_id = $1 ORDER BY created_at ASC", chatroomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var message Message
		err := rows.Scan(
			&message.Id,
			&message.Content,
			&message.UserId,
			&message.CreatedAt,
			&message.ChatroomId,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
