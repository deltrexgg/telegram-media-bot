package repository

import (
	"context"
	"database/sql"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
)

type FolderRepo interface {
	Create(ctx context.Context, folder domain.Folders) error
}

type repo struct {
	db *sql.DB
}

func NewFolderRepo(db *sql.DB) FolderRepo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, folder domain.Folders) error {
	_, err := r.db.Exec(
		"INSERT INTO folders(id, created_by, name) VALUES(?, ?, ?)",
		folder.ID, folder.CreatedBy, folder.Name,
	)
	return err
}
