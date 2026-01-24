package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
)

type FolderRepo interface {
	Create(ctx context.Context, folder domain.Folders) error
	Delete(ctx context.Context, id string) error
	GrandFolderAccess(ctx context.Context, access domain.Access) error
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

func (r *repo) Delete(ctx context.Context, id string) error {

	_, err := r.db.Exec(
		"DELETE FROM folders WHERE id = ?", id,
	)

	return err
}

func (r *repo) GrandFolderAccess(ctx context.Context, access domain.Access) error {
	_, err := r.db.Exec("INSERT INTO access(id, folder_id, user_id) VALUES(?, ?, ?);", access.ID, access.FolderID, access.UserID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
