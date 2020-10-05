//file where we work with database
package users

import (
	"fmt"
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/datasources/mysql/users_db"
	"github.com/simranchawla8241/Go-Microservices-UsersAPI/utils/errors"
	"time"
)

//local database
//var (
//	usersDB=make(map[int64]*User)
//)

const(
	queryInsertUser="INSERT INTO users(first_name,last_name,email,date_created) VALUES (?,?,?,?);"
	queryGetUser="SELECT * FROM users where id=?;"
	queryUpdateUser="UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser="DELETE FROM users where id=?;"
)

func (user *User)Get() (*errors.RestErr){
	stmt,err:=users_db.Client.Prepare(queryGetUser)
	if err!=nil{
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result:=stmt.QueryRow(user.Id)
	if err :=result.Scan(&user.Id,&user.FirstName,&user.Lastname,&user.Email,&user.DateCreated);err!=nil{
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get %d : %s",user.Id,err.Error()))
	}

	//local database code
	//result:= usersDB[user.Id]
	//if result== nil{
	//	return  errors.NewNotFoundError(fmt.Sprintf("User id %d not found",user.Id))
	//}
	//
	//user.Id=result.Id
	//user.FirstName=result.FirstName
	//user.Lastname=result.Lastname
	//user.Email=result.Email
	//user.DateCreated=result.DateCreated
	return nil
}

func (user *User)Save() *errors.RestErr{
	stmt,err:=users_db.Client.Prepare(queryInsertUser)
	if err!=nil{
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	now :=time.Now()
	user.DateCreated=now.Format("2006-01-02T15:04:05Z")
	insertResult,err:=stmt.Exec(user.FirstName,user.Lastname,user.Email,user.DateCreated)
	if err!=nil{
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save save user %s",err.Error()))
	}

	//if record is inserted
	userId,err:=insertResult.LastInsertId()
	if err !=nil{
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save save user %s",err.Error()))
	}
	user.Id=userId

	//this part of commented code is for temp db
	//current:= usersDB[user.Id]
	//if current!= nil{
	//	if current.Email ==user.Email{
	//		return  errors.NewNotFoundError(fmt.Sprintf("Email id %s already registered",user.Email))
	//	}
	//	return  errors.NewNotFoundError(fmt.Sprintf("User id %d already exists",user.Id))
	//}
	////time package
	//now :=time.Now()
	//user.DateCreated=now.Format("2006-01-02T15:04:05Z")
	//
	//usersDB[user.Id]= user
   return nil
}

func (user *User)Update() *errors.RestErr {
	stmt,err:=users_db.Client.Prepare(queryUpdateUser)
	if err!=nil{
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_,err =stmt.Exec(user.FirstName,user.Lastname,user.Email,user.Id)
	if err!=nil{
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save save user %s",err.Error()))
	}
	return nil
}

func (user *User)Delete() *errors.RestErr{

	stmt,err:=users_db.Client.Prepare(queryDeleteUser)
	if err!=nil{
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_,err =stmt.Exec(user.Id)
	if err!=nil{
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save save user %s",err.Error()))
	}
	return nil

}
