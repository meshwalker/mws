package customer

import "time"

type Customer struct {
	Id		string	`json:"id,omitempty" gorethink:"id,omitempty"`
	FirstName	string	`json:"first_name,omitempty" gorethink:"first_name,omitempty"`
	LastName	string	`json:"last_name,omitempty" gorethink:"last_name,omitempty"`
	Address		string	`json:"address,omitempty" gorethink:"address,omitempty"`
	ZIP		string	`json:"zip,omitempty" gorethink:"zip,omitempty"`
	City		string	`json:"city,omitempty" gorethink:"city,omitempty"`
	Country		string	`json:"country,omitempty" gorethink:"country,omitempty"`
	Email		string	`json:"email,omitempty" gorethink:"email,omitempty"`
	MobileNumber	string	`json:"mobile_number,omitempty" gorethink:"mobile_number,omitempty"`
	Password	[]byte	`json:"-" gorethink:"password,omitempty"`
	PasswordSalt	[]byte	`json:"-" gorethink:"password_salt,omitempty"`
	ClusterId	string	`json:"cluster_id,omitempty" gorethink:"cluster_id,omitempty"`
	CreatedAt	time.Time	`json:"created_at,omitempty" gorethink:"created_at,omitempty"`
	ModifiedAt	time.Time	`json:"modified_at,omitempty" gorethink:"modified_at,omitempty"`
}

type CustomerLogin struct {
	MobileNumber	string	`json:"mobile_number"`
	Password	string	`json:"password"`
}

type RestCustomer struct {
	FirstName	string	`json:"first_name" gorethink:"first_name"`
	LastName	string	`json:"last_name" gorethink:"last_name"`
	Address		string	`json:"address" gorethink:"address"`
	ZIP		string	`json:"zip" gorethink:"zip"`
	City		string	`json:"city" gorethink:"city"`
	Country		string	`json:"country" gorethink:"country"`
	Email		string	`json:"email" gorethink:"email"`
	MobileNumber	string	`json:"mobile_number" gorethink:"mobile_number"`
	Password	string	`json:"password" gorethink:"password"`
	ConfirmPassword	string	`json:"confirm_password" gorethink:"-"`
}