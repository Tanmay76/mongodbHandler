package main

import "time"

const (
	URLS       = ":8090"
	URLD       = "mongodb://localhost:27017"
	COLLECTION = "Bloguser"
	DATABASE   = "Blogsdb"
	SecretKey  = "secret"
)

type Bloguser struct {
	EmpId       int    `bson:"Id,omitempty" `
	Name        string `bson:"Name,omitempty"`
	Email       string `bson:"Email,omitempty"`
	Password    string `bson:"Password,omitempty"`
	Passwordh   []byte
	Gender      string    `bson:"Gender,omitempty"`
	DOB         time.Time `bson:"DOB,omitempty"`
	Address     Address   `bson:"Address,omitempty"`
	Phonenumber string    `bson:"PhoneNo,omitempty"`
	BlogId      int       `bson:"BlogId,omitempty"`
}
type Address struct {
	Street   string `bson:"Street,omitempty"`
	Locality string `bson:"locality,omitempty"`
	Pincode  int64  `bson:"Pincode,omitempty"`
}

type User struct {
	UserID   int    `json:"username"`
	Password string `json:"password"`
}
