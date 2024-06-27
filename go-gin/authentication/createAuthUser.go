package authentication

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/constants"
	"github.com/go-gin/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	var authInput constants.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound constants.AuthUser

	filter := bson.M{"username": authInput.Username}
	database.CollectionAuth.FindOne(context.Background(), filter).Decode(&userFound)
	fmt.Println(userFound.Username, "==")
	fmt.Println(userFound)
	if userFound.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := constants.AuthUser{
		Username:  authInput.Username,
		Password:  string(passwordHash),
		CreatedAt: time.Now(),
	}

	database.CollectionAuth.InsertOne(context.Background(), user)
	c.JSON(http.StatusOK, gin.H{"data": user})

}
