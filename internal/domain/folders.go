package domain

import "time"

type Folders struct {
	ID          string    `json:"id"`
	ChatID      string    `json:"chat_id"`
	FolderName  string    `json:"folder_name"`
	ContentSize int       `json:"content_size"`
	CreatedAt   time.Time `json:"created_at"`
}
