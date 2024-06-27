package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/database"
)

func DeleteUserDetailsByID(c *gin.Context) {
	id := c.Param("id")

	userDetails, err := database.DeleteUserDetailsByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "requested Id of user not found So cannot be deleted"})
	} else {
		c.IndentedJSON(http.StatusOK, userDetails)
	}

}
