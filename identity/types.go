package identity

import (
	"time"
	"net"
)


type User struct {
	Uid		string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	Address		UserAddress
}

type UserAddress struct {
	Street		string
	HoseNr		string
	Zip		string
	City		string
	Country		string
}