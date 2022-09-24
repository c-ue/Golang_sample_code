package main

import (
	"errors"
	"log"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func main() {
	// Choose auth method and set it up
	auth := LoginAuth("account", "password")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"john@gmail.com"}
	msg := []byte("From: \"account\" <account@gmail.com>\r\n" +
		"To: \"john\" <john@gmail.com>\r\n" +
		"Subject: Golang Test\r\n" +
		"\r\n" +
		"From Golang Example Code\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "account@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
