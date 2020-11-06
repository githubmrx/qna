package core

import (
	"fmt"
	"qna/constant"
	"qna/db"
	"qna/model"
)

func GetQuestions() (result []model.Question, err error) {
	_, err = db.Dbmap.Select(&result, "SELECT * FROM question")
	return
}

func GetQuestionById(id int) (result model.Question, err error) {
	err = db.Dbmap.SelectOne(&result, "SELECT * FROM question WHERE id=? LIMIT 1", id)
	if err != nil {
		err = constant.ERR_NOT_FOUND
	}
	return
}

func GetQuestionsByUsername(username string) (result []model.Question, err error) {
	query := `SELECT q.* FROM question q
	INNER JOIN user u ON u.id = q.created_by AND u.username = '%s'`
	_, err = db.Dbmap.Select(&result, fmt.Sprintf(query, username))
	return
}

func ValidateQuestionDataInput(data *model.Question) error {
	if data.Statement == "" || data.CreatedBy < 1 {
		return constant.ERR_INVALID_DATA
	}
	return nil
}

func UpdateQuestion(id int, data *model.Question) error {
	_, err := GetQuestionById(id)
	if err == nil {
		data.Id = id
		_, err = db.Dbmap.Update(data)
	}
	return err
}

func AddQuestion(data *model.Question) error {
	// err := db.Dbmap.Insert(data)
	result, err := db.Dbmap.Exec(`INSERT INTO question (statement, answer, created_by, answered_by) VALUES (?, ?, ?, ?)`, data.Statement, "", data.CreatedBy, 0)
	if err == nil {
		var id, err = result.LastInsertId()
		if err == nil {
			data.Id = int(id)
		}
	}
	return err
}

func DeleteQuestion(id int) error {
	item, err := GetQuestionById(id)
	if err == nil {
		_, err = db.Dbmap.Delete(&item)
	}
	return err
}
