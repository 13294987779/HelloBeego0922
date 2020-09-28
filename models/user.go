package models

import "net/mail"

type User struct {
	Name string `json:"name"`
	Birthday
	Address
	password
}
