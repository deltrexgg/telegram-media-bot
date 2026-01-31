package migrations

import (
	"database/sql"
	"log"
)

/*
	Init the basic schema's during the bot startup
	Check if the schema exists then implement
*/

func AutoMigrate(db *sql.DB) {

	_, err := db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		log.Fatalf("Failed to enable foreign keys: %v", err)
	}

	folderSmt := `
	CREATE TABLE IF NOT EXISTS folders (
		id TEXT PRIMARY KEY,
		created_by TEXT NOT NULL,
		name TEXT NOT NULL,
		content_size INTEGER,
		is_public INTEGER DEFAULT 1,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	fileSmt := `
	CREATE TABLE IF NOT EXISTS files (
		id TEXT PRIMARY KEY,
		file_id TEXT NOT NULL,
		type TEXT,
		folder_id TEXT NOT NULL,
		uploaded_by TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (folder_id) REFERENCES folders(id)
		ON DELETE CASCADE
	);`

	accessSmt := `
	CREATE TABLE IF NOT EXISTS access (
		id TEXT PRIMARY KEY,
		folder_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		joined_at TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (folder_id) REFERENCES folders(id)
	);`

	historySmt := `
	CREATE TABLE IF NOT EXISTS fetch_history (
	id TEXT PRIMARY KEY,
	folder_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	last_delivered_at TEXT DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (folder_id) REFERENCES folders(id) ON DELETE CASCADE,
	UNIQUE (user_id, folder_id)
	);`

	statusSmt := `
	CREATE TABLE IF NOT EXISTS upload_status (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		folder_id TEXT NOT NULL,
		FOREIGN KEY (folder_id) REFERENCES folders(id)
	);
	`

	stmts := []string{
		folderSmt,
		fileSmt,
		accessSmt,
		historySmt,
		statusSmt,
	}

	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	log.Println("Migration completed")
}
