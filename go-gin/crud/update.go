package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/constants"
	"github.com/go-gin/database"
)

func UpdateUserDetailsById(c *gin.Context) {

	var replaceUser constants.User

	if err := c.BindJSON(&replaceUser); err != nil {
		return
	}
	id := c.Param("id")

	userDetails, err := database.UpdateUserDetailsByID(id, replaceUser)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "update cannot be done"})
	} else {
		c.IndentedJSON(http.StatusOK, userDetails)
	}
}
