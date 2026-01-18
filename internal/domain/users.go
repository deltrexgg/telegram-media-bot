package domain

import "time"

type Users struct {
	ID             string    `json:"id"`
	TelegramUserId string    `json:"telegram_user_id"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
}
