package app

import "github.com/simranchawla8241/Go-Microservices-UsersAPI/controllers"

func mapUrls() {
	router.GET("/ping",controllers.Ping)

	router.GET("/users/:user_id",controllers.GetUser)
	//router.GET("/users/:user_id",controllers.SearchUser)
	router.POST("/users",controllers.CreateUser)


}
