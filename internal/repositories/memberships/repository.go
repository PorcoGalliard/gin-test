package memberships

import (
	"database/sql"
	"log"
)

type repository struct {
	sql *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("select id, email FROM users")
	if err != nil {
		log.Println("<<<Error Query>>>", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println("Error Scan", err)
		}
		log.Printf("id: %d, email: %s\n\n", id, email)
	}
	return &repository{
		sql: db,
	}
}