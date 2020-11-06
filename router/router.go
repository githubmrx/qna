package router

import (
	"qna/common"
	"qna/endpoint"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(common.Cors())
	v1 := Router.Group("/v1")
	{
		v1.GET("/questions/", GetQuestionsHandler)
		// v1.GET("/questions/:id", endpoint.GetQuestion)
		// v1.GET("/questions/byusername/:username", endpoint.GetQuestionsByUsername)
		v1.GET("/questions/:param1", GetQuestionsHandler)         // /questions/:id
		v1.GET("/questions/:param1/:param2", GetQuestionsHandler) // /questions/byusername/:username
		v1.POST("/questions", endpoint.AddQuestion)
		v1.PUT("/questions/:id", endpoint.UpdateQuestion)
		v1.DELETE("/questions/:id", endpoint.DeleteQuestion)
	}
}

func GetQuestionsHandler(c *gin.Context) {
	param1 := c.Param("param1")
	param2 := c.Param("param2")

	if param1 == "byusername" {
		endpoint.GetQuestionsByUsername(c, param2)
	} else if param1 == "" && param2 == "" {
		endpoint.GetQuestions(c)
	} else {
		endpoint.GetQuestion(c, param1)
	}
}
