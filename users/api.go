package users

import (
	"encoding/json"
	"net/http"

	"github.com/muse/common"
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
		panic(common.ClientError{Detail: "incorrect payload"})
	}

	if schema.Username == nil || schema.Password == nil {
		panic(common.ClientError{Detail: "incorrect payload"})
	}

	userRepository := getUserRepository()
	idFactory := getIdFactory()

	err = CreateUser(userRepository, idFactory, *schema.Username, *schema.Password)
	if err != nil {
		panic(err)
	}

	respData := CreateUserOut{Status: "created"}
	resp, err := json.Marshal(respData)
	if err != nil {
		panic(common.UnrecoverableError{Detail: "internal server error"})
	}

	w.WriteHeader(201)
	w.Write(resp)
}
