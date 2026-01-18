package domain

import "time"

type FetchHistory struct {
	ID       string `json:"id"`
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`

	LastDeliveredAt time.Time `json:"last_delivered_at"`
}
