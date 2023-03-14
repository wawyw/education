
package controller

import "education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}

var users []User

func init() {
	admin := User{LoginName:"admin", Password:"admin", IsAdmin:"T"}
	bob := User{LoginName:"bob", Password:"bob", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, bob)
}