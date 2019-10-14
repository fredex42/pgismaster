package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	log.Printf("pgismaster, Andy Gallagher 2019. See https://github.com/fredex42/pgismaster for details.")
	connStr := "user=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to local database: %s", err)
	}

	rows, err := db.Query("select pg_is_in_recovery();")

	if err != nil {
		log.Fatalf("Could not query database: %s", err)
	}

	var isInRecovery bool

	rows.Next()
	scanErr := rows.Scan(&isInRecovery)
	if scanErr != nil {
		log.Fatalf("Could not interpret data from server: %s", err)
	}

	if isInRecovery {
		log.Printf("Server is a standby")
		os.Exit(1)
	} else {
		log.Printf("Server is a master")
		os.Exit(0)
	}
}
