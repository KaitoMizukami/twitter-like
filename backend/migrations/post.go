package migrations

import "database/sql"

func CreatePostsTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			createdAt DATE
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
