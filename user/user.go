package user

import (
	"dudan/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       uint   `json: "id"`
	Email    string `json: "email"`
	Password string `json: "password"`
	Token    string `json: "token"`
}

func RegisterUser(c *gin.Context) {
	db := database.DbConnect()
	var newUser User
	userObj := User{Token: ""}

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{Id: newUser.Id, Email: newUser.Email, Password: newUser.Password, Token: userObj.Token}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	db := database.DbConnect()
	// check user
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func FetchUsers(c *gin.Context) {
	db := database.DbConnect()
	var users = []User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
