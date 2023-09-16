package model

import (
	"askida-tatli/db"
)

type User struct {
	db.Model
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Project   string `bson:"project"`
}
