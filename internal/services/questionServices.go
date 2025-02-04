package services

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/siam-vit/coding-relay-be/internal/database"
	"github.com/siam-vit/coding-relay-be/internal/models"
)

func CreateQuestion(question models.Question) error {
	db := database.DB.Db

	_, err := db.Exec(`
		INSERT INTO questions ( question, set, difficulty)
		VALUES ($1, $2, $3)`,
		question.Question, question.Set, question.Difficulty)
	return err
}

func GetQuestionsByDifficulty(difficulty string) ([]models.Question, error) {
	db := database.DB.Db
	var questions []models.Question

	rows, err := db.Query(`SELECT id, question, test_case_id, set, difficulty FROM questions WHERE difficulty = $1`, difficulty)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q models.Question
		var testCaseIDs pq.StringArray

		err := rows.Scan(&q.ID, &q.Question, &testCaseIDs, &q.Set, &q.Difficulty)
		if err != nil {
			return nil, err
		}

		if len(testCaseIDs) == 0 {
			q.TestCaseID = []uuid.UUID{}
		} else {
			for _, id := range testCaseIDs {
				parsedID, err := uuid.Parse(id)
				if err != nil {
					return nil, err
				}
				q.TestCaseID = append(q.TestCaseID, parsedID)
			}
		}

		questions = append(questions, q)
	}
	return questions, nil
}

func CreateTestCase(testCase models.TestCases) error {
	db := database.DB.Db

	_, err := db.Exec(`
		INSERT INTO test_cases (id, input, output, question_id)
		VALUES ($1, $2, $3, $4)`,
		uuid.New(), testCase.Input, testCase.Output, testCase.QuestionID)
	return err
}

func GetAllTestCases() ([]models.TestCases, error) {
	db := database.DB.Db
	var testCases []models.TestCases

	rows, err := db.Query(`SELECT id, input, output, question_id FROM test_cases`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tc models.TestCases
		err := rows.Scan(&tc.ID, &tc.Input, &tc.Output, &tc.QuestionID)
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, tc)
	}
	return testCases, nil
}
