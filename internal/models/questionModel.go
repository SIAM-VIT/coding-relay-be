package models

import "github.com/google/uuid"

type TestCases struct {
	ID         uint      `json:"id"`
	Input      string    `json:"input"`
	Output     string    `json:"output"`
	QuestionID uuid.UUID `json:"questionID"`
}

type Question struct {
	ID         uint        `json:"id"`
	Question   string      `json:"question"`
	TestCaseID []uuid.UUID `json:"testCaseId"`
	Set        uint        `json:"set"`
	Difficulty string      `json:"difficulty"`
}

//Post request, difficulty based on difficutly 