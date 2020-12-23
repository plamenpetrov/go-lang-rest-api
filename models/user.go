package models

import (
	"fmt"
	"errors"
)

type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
}

var (
	users []*User			//slice of user with pointers to User type. This way we have only one user in whole project
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error)  {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID")
	}
	u.ID = nextId
	nextId++
	users = append(users, &u)

	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with '%v' id not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with '%v' id not found", u.ID)
}

func RemoveUserById(id int) (error) {
	for i, u := range users {
		if u.ID == id {
			//BLACK MAGIC: get all elements form the begining of slice to needed 
			//and all elements after element that must be removed
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with '%v' id not found", id)
}