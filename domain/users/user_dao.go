//file where we work with database
package users

import (
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/utils/errors"
	"fmt"
)

var (
	usersDB=make(map[int64]*User)
)

func (user *User)Get() (*errors.RestErr){
	result:= usersDB[user.Id]
	if result== nil{
		return  errors.NewNotFoundError(fmt.Sprintf("User id %d not found",user.Id))
	}

	user.Id=result.Id
	user.FirstName=result.FirstName
	user.Lastname=result.Lastname
	user.Email=result.Email
	user.DateCreated=result.DateCreated
	return nil
}

func (user *User)Save() *errors.RestErr{
	current:= usersDB[user.Id]
	if current!= nil{
		if current.Email ==user.Email{
			return  errors.NewNotFoundError(fmt.Sprintf("Email id %s already registered",user.Email))
		}
		return  errors.NewNotFoundError(fmt.Sprintf("User id %d already exists",user.Id))
	}
	usersDB[user.Id]= user
   return nil
}

