package domain

import "time"

type Files struct {
	ID       string `json:"id"`
	FileID   string `json:"file_id"`
	Type     string `json:"type"`
	FolderId string `json:"folder_id"`

	UploadedBy string    `json:"uploaded_by"`
	CreatedAt  time.Time `json:"created_at"`
}
