package users

import (
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/utils/errors"
	"strings"
)

type User struct{
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User)Validate() *errors.RestErr{
	//check if email address of user is valid
	user.Email=strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email==""{
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}