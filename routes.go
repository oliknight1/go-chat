package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateChatroomRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateMessageRequest struct {
	ChatroomId int    `json:"chatroom_id" validate:"required"`
	UserID     int    `json:"user_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
}

func (rh RequestHandler) PostMessages(rw http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var response CreateMessageRequest
	err = json.Unmarshal(body, &response)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := Message{
		UserId:     response.UserID,
		Content:    response.Content,
		ChatroomId: response.ChatroomId,
	}
	CreateMessage(rh.db, msg)
	fmt.Fprint(rw, "Success")
}

func (rh RequestHandler) PostChatroom(rw http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var response CreateChatroomRequest
	err = json.Unmarshal(body, &response)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	CreateChatroom(rh.db, response.Name)

}
func (rh RequestHandler) PostUser(rw http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	var response CreateUserRequest
	err = json.Unmarshal(body, &response)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	CreateChatroom(rh.db, response.Name)

}
