package main

import "fmt"

type Users struct {
	username string
	email    string
	password string
}

func (u *Users) changePassword(pasword string) {
	u.password = pasword

}

func (u Users) getDetails() string {
	return fmt.Sprintf("username %s  email %s  password %s", u.username, u.email, u.password)
}

func main() {
	user := Users{
		username: "john doe",
		email:    "jhondoe@gmail.com",
		password: "myseceretpassword",
	}

	fmt.Println(user.getDetails())
	user.changePassword("superseceretpassword")
	fmt.Println("New password", user.password)

}
