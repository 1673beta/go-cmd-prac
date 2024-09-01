package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	// Loading configuration from .config/default.yml
	dbConfig, err := LoadConfig()
	if err != nil {
		fmt.Printf("error while loading config: %v", err)
		os.Exit(1)
	}

	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.DB.Host, dbConfig.DB.Port, dbConfig.DB.User, dbConfig.DB.Pass, dbConfig.DB.Db)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		fmt.Printf("error while connecting to database: %v", err)
		os.Exit(120)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error while pinging database: %v", err)
		os.Exit(1)
	}

	log.Println("Connected to database")
	return db, nil
}
