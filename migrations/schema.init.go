package migrations

import (
	"log"

	"github.com/deltrexgg/telegram-media-bot/internal/config"
)

/*
	Init the basic schema's during the bot startup
	Check if the schema exists then implement
*/

func AutoMigrate() {

	folderSmt := `
	CREATE TABLE IF NOT EXISTS folders (
	id TEXT PRIMARY KEY,
	created_by TEXT NOT NULL,
	name TEXT NOT NULL,
	content_size INTEGER,
	is_public INTEGER DEFAULT 1,
	created_at TEXT DEFAULT CURRENT_TIMESTAMP
	);`

	/*
		fileSmt := `
		CREATE TABLE IF NOT EXISTS files (
		id TEXT PRIMARY KEY,
		file_id TEXT NOT NULL,
		type TEXT,
		folder_id TEXT NOT NULL,
		uploaded_by TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP
		);
		` */

	_, err := config.DB.Exec(folderSmt)
	if err != nil {
		log.Fatalf("Error creating table")
	}

	log.Println("Migration completed")
}
