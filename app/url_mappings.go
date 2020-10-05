package app

import "github.com/simranchawla8241/Go-Microservices-UsersAPI/controllers"

func mapUrls() {
	router.GET("/ping",controllers.Ping)

	router.GET("/users/:user_id",controllers.GetUser)
	router.POST("/users",controllers.CreateUser)
	router.PUT("/users/:user_id",controllers.UpdateUser)
	router.DELETE("/users/:user_id",controllers.DeleteUser)

}
