package controllers

import (
	"fmt"
	"log"
	"typathon/config"
	"typathon/models/entity"
	"typathon/models/service"
	"typathon/utils"

	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo/bson"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

//AuthController is for auth logic
type AuthController struct{}

var conf *oauth2.Config

//Login is to process login request
func (auth *AuthController) Login(c *gin.Context) {
	fmt.Println(c)

	var loginInfo entity.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	//TODO
	userservice := service.Userservice{}
	user, errf := userservice.FindByEmail(loginInfo.Email)
	if errf != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
		return
	}

	token, err := user.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	//-------
	c.JSON(200, gin.H{
		"token": token,
	})
}

//Profile is to provide current user info
func (auth *AuthController) Profile(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	c.JSON(200, gin.H{
		"user_name": user.Name,
		"email":     user.Email,
	})
}

//Signup is for user signup
func (auth *AuthController) Signup(c *gin.Context) {

	type signupInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
		UserName string `json:"userName" binding:"required"`
	}
	var info signupInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	user := entity.User{}
	user.Email = info.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = info.Name
	user.UserName = info.UserName
	userservice := service.Userservice{}
	err = userservice.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
	return
}

func (auth *AuthController) ChangePassword(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))
	type passwordUpdateObject struct {
		Password string `json:"password" bson:"password" binding:"required"`
	}
	var newPassword passwordUpdateObject
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	userservice := service.Userservice{}
	err = userservice.UpdateOne("email", user.Email, bson.M{"password": string(hash)})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
	return
}

func (auth *AuthController) ResetPassword(c *gin.Context) {
	type User struct {
		Email    string `json:"email" binding:"required"`
		ResetUrl string `json:"resetUrl" binding:"required"`
	}
	var requestUser User
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	user := entity.User{}
	user.Email = requestUser.Email

	token, err := user.GetJwtToken()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	//-------
	body := "You can reset your password by following this link: <a href=" + utils.RemoveTrailingSlash(requestUser.ResetUrl) + "/" + token + ">password reset link</a>."
	utils.SendEmail(config.ConfigValues, requestUser.Email, body)
	c.JSON(200, gin.H{})
}
