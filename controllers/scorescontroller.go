package controllers

import (
	"fmt"
	"typathon/models/entity"
	"typathon/models/service"

	"github.com/gin-gonic/gin"
)

type ScoresController struct{}

func (route *ScoresController) SaveScore(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	var scoreInfo entity.Score
	if err := c.ShouldBindJSON(&scoreInfo); err != nil {
		fmt.Println("Could not bind json for creation of new score creation.")
		c.AbortWithStatusJSON(401, gin.H{"error": "Please input all fields"})
		return
	}
	scoreInfo.User = user.UserName

	scoreservice := service.ScoreService{}
	err := scoreservice.Create(&scoreInfo)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
	return
}

func (route *ScoresController) GetUserScores(c *gin.Context) {
	user := c.MustGet("user").(*(entity.User))

	scoreservice := service.ScoreService{}
	scores, err := scoreservice.FindUserHighestScores("user", user.UserName)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"scores": scores})
	}
	return
}

func (router *ScoresController) GetHighScores(c *gin.Context) {
	scoreservice := service.ScoreService{}
	scores, err := scoreservice.FindHighestScores()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"scores": scores})
	}
	return
}
