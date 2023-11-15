package routes

import (
	"NewProject/controllers"
	"net/http"
)

func SetupRoutes() {
	userController := &controllers.UserController{}

	http.HandleFunc("/api/user", userController.GetUser)
	http.HandleFunc("/api/user/create", userController.CreateUser)
	http.HandleFunc("/api/user/update", userController.UpdateUser)
}
