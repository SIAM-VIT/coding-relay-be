package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/siam-vit/coding-relay-be/internal/utils"
)

type Dbinstance struct {
	Db *sqlx.DB
}

var DB Dbinstance

func Connect() {
	p := utils.Config("DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", utils.Config("DB_HOST"), utils.Config("DB_USER"), utils.Config("DB_PASSWORD"), utils.Config("DB_NAME"), port)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		log.Fatal("Failed to ping the database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected")

	runMigrations(db)

	DB = Dbinstance{
		Db: db,
	}
}

func runMigrations(db *sqlx.DB) {
	_, err := db.Exec(`

		CREATE TABLE IF NOT EXISTS questions (
			id SERIAL PRIMARY KEY,  
			question TEXT NOT NULL,
			test_case_id UUID[],
			set INT NOT NULL,
			difficulty VARCHAR(255) NOT NULL
		
		);
		
		CREATE TABLE IF NOT EXISTS test_cases (
			id UUID PRIMARY KEY,
			input TEXT NOT NULL,
			output TEXT NOT NULL,
			question_id INTEGER REFERENCES questions(id)
		);

		CREATE TABLE IF NOT EXISTS teams (
			team_id UUID PRIMARY KEY,
			team_name VARCHAR(255) NOT NULL,
			team_members TEXT[] NOT NULL,
			score INT NOT NULL
		);

	`)

	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
		os.Exit(2)
	}

	log.Println("Migrations completed")
}
