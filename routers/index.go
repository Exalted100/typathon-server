package routers

import (
	"typathon/controllers"

	"typathon/middlewares"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	scoreController := new(controllers.ScoresController)

	router.Use(middlewares.CORSMiddleware())
	router.GET("/healthz", health)
	router.POST("/login", authController.Login)
	router.POST("/signup", authController.Signup)
	router.PUT("/password/reset", authController.ResetPassword)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/profile", authController.Profile)
	authGroup.GET("/player/scores", scoreController.GetUserScores)
	authGroup.GET("/global/scores", scoreController.GetHighScores)
	authGroup.POST("/game/score", scoreController.SaveScore)
	authGroup.PUT("/password/change", authController.ChangePassword)
}

// InitRoute ..
func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)
	return router
}

func health(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "A-O-Kay",
	})
}