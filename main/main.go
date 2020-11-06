package main

import (
	"qna/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router.CreateUrlMappings()
	router.Router.Run(":8080")

}
