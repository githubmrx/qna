package core

import (
	"qna/model"
	"reflect"
	"testing"
)

func TestGetQuestionById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name       string
		args       args
		wantResult model.Question
		wantErr    bool
	}{
		{"not found", args{0}, model.Question{}, true},
		{"user1", args{1}, model.Question{Id: 1, Statement: "q1", Answer: "", CreatedBy: 1, AnsweredBy: 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetQuestionById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuestionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetQuestionById() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetQuestionsByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []model.Question
		wantErr    bool
	}{
		{"no matches", args{"userxx"}, []model.Question{}, false},
		{"user1", args{"user1"}, []model.Question{{Id: 1, Statement: "q1", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetQuestionsByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuestionsByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetQuestionsByUsername() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestValidateQuestionDataInput(t *testing.T) {
	type args struct {
		data *model.Question
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid statement", args{&model.Question{Id: 1, Statement: "", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, true},
		{"invalid author", args{&model.Question{Id: 1, Statement: "q1", Answer: "", CreatedBy: 0, AnsweredBy: 0}}, true},
		{"valid", args{&model.Question{Id: 1, Statement: "q1", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateQuestionDataInput(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateQuestionDataInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateQuestion(t *testing.T) {
	type args struct {
		id   int
		data *model.Question
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not found", args{11, &model.Question{Statement: "q1", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, true},
		{"successful", args{1, &model.Question{Statement: "q1", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateQuestion(tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddQuestion(t *testing.T) {
	type args struct {
		data *model.Question
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"successful", args{&model.Question{Statement: "q3", Answer: "", CreatedBy: 1, AnsweredBy: 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddQuestion(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AddQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteQuestion(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"not found", args{44}, true},
		{"successful", args{5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteQuestion(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
