package common

import (
	"log"
	"qna/constant"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func MakeResponse(c *gin.Context, err error, rsData interface{}, okCode int) {
	if err == nil {
		c.JSON(okCode, rsData)
	} else {
		// common.CheckErr(err, "GetQuestionsByUsername error")
		errStr := err.Error()
		var code int
		switch errStr {
		case constant.STR_INVALID_DATA:
			code = 400
		case constant.STR_NOT_FOUND:
			code = 404
		default:
			code = 500
			log.Println("Error:", errStr)
		}
		c.JSON(code, gin.H{"error": errStr})
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
