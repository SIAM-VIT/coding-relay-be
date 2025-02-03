package models

import "github.com/google/uuid"

type TestCases struct {
	ID         uuid.UUID `json:"id"`
	Input      string    `json:"input"`
	Output     string    `json:"output"`
	QuestionID uuid.UUID `json:"questionID"`
}

type Question struct {
	ID         uuid.UUID   `json:"id"`
	Question   string      `json:"question"`
	TestCaseID  []uuid.UUID `json:"testCaseId"`
	Set        uint        `json:"set"`
	Difficulty string      `json:"difficulty"`
}
	