package services

import (
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/domain/users"
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/utils/errors"
)

func GetUser(userId int64)(*users.User,*errors.RestErr){
	result:=&users.User{Id:userId}
	if err:=result.Get();err !=nil{
		return nil,err
	}
	return result,nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
    //check if email address of user is valid
	if err:=user.Validate();err!=nil{
		return nil,err
	}
	if err:=user.Save();err!=nil{
		return nil,err
	}
	return &user,nil
}