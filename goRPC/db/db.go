package db

import (
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pgext"
)

// Database instance
var Database *pg.DB

// Connect to database
func Connect() {
	pgHost, pgUser, pgPass, pgDb := getDatabaseCreds()

	Database = pg.Connect(&pg.Options{
		Addr:     pgHost,
		User:     pgUser,
		Password: pgPass,
		Database: pgDb,
	})

	Database.AddQueryHook(&pgext.DebugHook{
		Verbose: true,
	})
}

func getDatabaseCreds() (string, string, string, string) {
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		pgHost = "localhost:5432"
	}
	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		pgUser = "postgres"
	}
	pgPass, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		pgPass = "postgres"
	}
	pgDb, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		pgDb = "postgres"
	}

	return pgHost, pgUser, pgPass, pgDb
}
