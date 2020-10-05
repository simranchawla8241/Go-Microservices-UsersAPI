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

func UpdateUser(user users.User)(*users.User, *errors.RestErr){
	//get that user frm db
	current,err :=GetUser(user.Id)
	if err !=nil{
		return nil,err
	}
	current.FirstName=user.FirstName
	current.Lastname=user.Lastname
	current.Email=user.Email

	if err:=current.Update();err !=nil{
		return nil,err
	}
	return current,nil
}

func DeleteUser(userId int64) *errors.RestErr {
	//get that user frm db
	user:=&users.User{Id:userId}
	return user.Delete()
}