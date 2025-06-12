package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to DB %+v\n", err)
		return nil, err
	}

	return db, nil

}