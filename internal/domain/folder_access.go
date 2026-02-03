package domain

import "time"

type Access struct {
	ID       string `json:"id"`
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`

	JoinedAt time.Time `json:"joined_by"`
}
