package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u *User) OutputUserDetails() {	
	// ...
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	// ...
	return Admin {
		email: email,
		password: password,
		User: User {
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
func New(firstName, lastName, birthdate string) (*User, error) { //New = NewUser
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}
	createdAt := time.Now()
	return &User {
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: createdAt,
	} , nil
}