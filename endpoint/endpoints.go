package endpoint

import (
	"qna/common"
	"qna/core"
	"qna/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetQuestions(c *gin.Context) {
	questions, err := core.GetQuestions()

	common.MakeResponse(c, err, questions, 200)
}

func GetQuestion(c *gin.Context, id string) {
	intId, err := strconv.Atoi(id)
	var question model.Question
	if err == nil {
		question, err = core.GetQuestionById(intId)
	}

	common.MakeResponse(c, err, question, 200)
}

func GetQuestionsByUsername(c *gin.Context, username string) {
	questions, err := core.GetQuestionsByUsername(username)

	common.MakeResponse(c, err, questions, 200)
}

func AddQuestion(c *gin.Context) {
	var data model.Question
	c.Bind(&data)
	err := core.ValidateQuestionDataInput(&data)
	if err == nil {
		err = core.AddQuestion(&data)
	}

	common.MakeResponse(c, err, data, 201)
}

func UpdateQuestion(c *gin.Context) {
	var data model.Question
	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	if err == nil {
		c.Bind(&data)
		err = core.ValidateQuestionDataInput(&data)
		if err == nil {
			err = core.UpdateQuestion(intId, &data)
		}
	}

	common.MakeResponse(c, err, data, 200)
}

func DeleteQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	intId, err := strconv.Atoi(id)
	if err == nil {
		err = core.DeleteQuestion(intId)
	}

	common.MakeResponse(c, err, gin.H{"msg": "Item deleted successfully"}, 200)
}
