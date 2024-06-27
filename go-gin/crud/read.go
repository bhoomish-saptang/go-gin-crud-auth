package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/database"
)

func GetAllUserDetails(c *gin.Context) {
	UserData, err := database.GetAllUserDetails()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, UserData)
	} else {
		c.IndentedJSON(http.StatusOK, UserData)
	}
}

func GetUserDetailsById(c *gin.Context) {

	id := c.Param("id")

	userDetails, err := database.FindUserDetailsByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "requested Id of user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, userDetails)
	}
}
