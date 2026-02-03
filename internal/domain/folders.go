package domain

import "time"

type Folders struct {
	ID        string `json:"id"`
	CreatedBy string `json:"created_by"`

	Name        string `json:"folder_name"`
	ContentSize int    `json:"content_size"`
	IsPublic    bool   `json:"is_public"`

	CreatedAt time.Time `json:"created_at"`
}

type FoldersList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
