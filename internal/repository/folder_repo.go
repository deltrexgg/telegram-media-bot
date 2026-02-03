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
	FolderNameByUserId(ctx context.Context, userID string) ([]domain.FoldersList, error)
}

type folderrepo struct {
	db *sql.DB
}

func NewFolderRepo(db *sql.DB) FolderRepo {
	return &folderrepo{db: db}
}

func (r *folderrepo) Create(ctx context.Context, folder domain.Folders) error {
	_, err := r.db.Exec(
		"INSERT INTO folders(id, created_by, name) VALUES(?, ?, ?)",
		folder.ID, folder.CreatedBy, folder.Name,
	)
	return err
}

func (r *folderrepo) Delete(ctx context.Context, id string) error {

	_, err := r.db.Exec(
		"DELETE FROM folders WHERE id = ?", id,
	)

	return err
}

func (r *folderrepo) GrandFolderAccess(ctx context.Context, access domain.Access) error {
	_, err := r.db.Exec("INSERT INTO access(id, folder_id, user_id) VALUES(?, ?, ?);", access.ID, access.FolderID, access.UserID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (r *folderrepo) FolderNameByUserId(ctx context.Context, userID string) ([]domain.FoldersList, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT f.id, f.name
		 FROM access a
		 JOIN folders f ON a.folder_id = f.id
		 WHERE a.user_id = ?;`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var folders []domain.FoldersList

	for rows.Next() {
		var f domain.FoldersList
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return nil, err
		}
		folders = append(folders, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return folders, nil
}
