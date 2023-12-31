package userService

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	u "github.com/vilmis04/eurovision-game-monorepo/tree/main/apps/backend-service-go/user/userModel"
)

type UserService struct{}

// Create a new userModel object
func New() *UserService {
	return &UserService{}
}

// TODO: get user from storage
func (s *UserService) User(id string) (*u.User, error) {

	fmt.Printf("getting data for user %v\n", id)
	var err error = nil

	return &u.User{}, err
}

// TODO: update user in storage
func (s *UserService) UpdateUser(id string, user u.User) error {

	fmt.Printf("updating data for user %v\n", user)
	var err error = nil

	return err
}

// TODO: create user in storage
func (s *UserService) NewUser(request *http.Request) (string, error) {

	type response struct {
		Roles []string
	}
	newUser := u.User{ID: "123"}
	var body response
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	newUser.Roles = body.Roles

	fmt.Printf("create a new user with this data: %v\n", newUser)
	userId := ""

	return userId, err
}

// TODO: delete user from storage
func (s *UserService) DeleteUser(id string) error {

	fmt.Printf("delete user with id %v\n", id)
	var err error = nil

	return err
}
