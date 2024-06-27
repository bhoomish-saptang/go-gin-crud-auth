package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/constants"
	"github.com/go-gin/database"
)

func PostUserDetails(c *gin.Context) {
	var newUser constants.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, newUser)
	} else {
		id := newUser.ID
		_, err := database.FindUserDetailsByID(id)
		if err == nil {
			c.IndentedJSON(http.StatusAlreadyReported, gin.H{"message": "the User details ID is already exist"})
		} else {
			database.InsertUserData(newUser)
			c.IndentedJSON(http.StatusCreated, newUser)
		}
	}
}
