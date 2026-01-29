package repository

import (
	"context"
	"database/sql"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
)

type FileRepo interface {
	Get(ctx context.Context, folder_id string) ([]string, error)
	AddStatus(ctx context.Context, details domain.Status) error
	StopFileShare(ctx context.Context, user_id string) error
}

type filerepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) FileRepo {
	return &filerepo{db: db}
}

func (r *filerepo) Get(ctx context.Context, folder_id string) ([]string, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id FROM files WHERE folder_id = ?;`, folder_id,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fileids []string
	for rows.Next() {
		var fileid string
		if err := rows.Scan(&fileid); err != nil {
			return nil, err
		}
		fileids = append(fileids, fileid)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fileids, nil
}

func (r *filerepo) AddStatus(ctx context.Context, details domain.Status) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"DELETE FROM upload_status WHERE user_id = ?;",
		details.UserID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO upload_status(id, user_id, folder_id) VALUES(?, ?, ?);",
		details.ID,
		details.UserID,
		details.FolderID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *filerepo) StopFileShare(ctx context.Context, user_id string) error {
	_, err := r.db.Exec("DELETE FROM upload_status WHERE user_id = ?;", user_id)
	if err != nil {
		return err
	}
	return nil
}
