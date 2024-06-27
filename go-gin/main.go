package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/authentication"
	"github.com/go-gin/config"
	"github.com/go-gin/crud"
	"github.com/go-gin/database"
)

func init() {
	config.SetEnvironmentVariables()
	database.ConnectMongodb()
}
func main() {

	route := gin.Default()

	//userDetails CRUD API endpoints
	route.POST("/createUserDetails", crud.PostUserDetails)
	route.GET("/getAllUserDetails", crud.GetAllUserDetails)
	route.GET("/getUserDetailsById/:id", crud.GetUserDetailsById)
	route.PUT("/updateUserDetailsById/:id", crud.UpdateUserDetailsById)
	route.DELETE("/deleteUserDetailsById/:id", crud.DeleteUserDetailsByID)

	//Authenticaton Endpoints
	route.POST("/createAuthUser", authentication.CreateUser)
	route.POST("/authUserLogin", authentication.AuthUserLogin)
	route.Run(os.Getenv("PORT"))

}
