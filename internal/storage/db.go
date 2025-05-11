package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/buffalo-big-tech-system/bookcenter-grpc-go/internal/config"
)

func DBOpen() *sql.DB {
	log.Println("Start db init...")

	dbCfg := config.LoadDBConfig()

	connStr := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.Name)

	sqlDB, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Cannot connect to DB: %s", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Cannot ping to DB: %s", err)
	}

	log.Println("Finish db init...")
	return sqlDB
}

func DBClose(db *sql.DB) {
	log.Println("Start DB close...")
	if err := db.Close(); err != nil {
		log.Fatalf("Cannot close database connection: %s", err)
	}
	log.Println("Finish DB close...")
}
