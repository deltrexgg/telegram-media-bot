package domain

import "time"

type Files struct {
	ID        string    `json:"id"`
	FileID    string    `json:"file_id"`
	FolderId  string    `json:"folder_id"`
	CreatedAt time.Time `json:"created_at"`
}
