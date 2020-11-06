package model

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Fullname string `db:"fullname" json:"fullname"`
}

type Question struct {
	Id         int    `db:"id" json:"id"`
	Statement  string `db:"statement" json:"statement"`
	Answer     string `db:"answer" json:"answer"`
	CreatedBy  int    `db:"created_by" json:"created_by"`
	AnsweredBy int    `db:"answered_by" json:"answered_by"`
}
