package services

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/siam-vit/coding-relay-be/internal/database"
	"github.com/siam-vit/coding-relay-be/internal/models"
)

func CreateTeam(team models.Teams) error {
	db := database.DB.Db
	// ctx := context.Background()
	// database.RedisClient.Del(ctx, "teams_by_score")
	_, err := db.Exec(`
        INSERT INTO team (team_id, team_name, team_members, score)
        VALUES ($1, $2, $3, $4)`,
		uuid.New(), team.TeamName, pq.Array(team.TeamMembers), team.Score)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTeams() ([]models.Teams, error) {
	db := database.DB.Db

	rows, err := db.Query(`SELECT team_id, team_name, team_members, score FROM team`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Teams

	for rows.Next() {
		var team models.Teams
		var members pq.StringArray

		err := rows.Scan(&team.TeamID, &team.TeamName, &members, &team.Score)
		if err != nil {
			return nil, err
		}

		team.TeamMembers = []string(members)
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}
