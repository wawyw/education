
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
	root := User{LoginName:"root", Password:"root", IsAdmin:"T"}
	alice := User{LoginName:"alice", Password:"alice", IsAdmin:"F"}
	bob := User{LoginName:"bob", Password:"bob", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, root)
	users = append(users, alice)
	users = append(users, bob)

}