package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserIn struct {
	Username *string
	Password *string
}

type CreateUserOut struct {
	Status string `json:"status"`
}

func CreateUserView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	schema := CreateUserIn{}
	err := decoder.Decode(&schema)
	if err != nil {
		// This is still a client error.
		// Data should be parseable.
		// todo: create middleware for IncorrectPayload
		data := `{"details": "malformed body"}`
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	if schema.Username == nil || schema.Password == nil {
		// todo: create middleware for IncorrectPayload
		data := `{"details": "incorrect body"}`
		w.WriteHeader(400)
		w.Write([]byte(data))
		return
	}

	userRepository := getUserRepository()
	idFactory := getIdFactory()

	err = CreateUser(userRepository, idFactory, *schema.Username, *schema.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	respData := CreateUserOut{Status: "created"}
	resp, err := json.Marshal(respData)
	if err != nil {
		// todo: create middleware for UnrecoverableError (or UnexpectedError)
		panic(err)
	}

	w.WriteHeader(201)
	w.Write([]byte(resp))
}
