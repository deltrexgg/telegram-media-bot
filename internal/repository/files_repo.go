package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
	"github.com/deltrexgg/telegram-media-bot/internal/utils"
)

type FileRepo interface {
	Get(ctx context.Context, folder_id string, lastVisit string) ([]domain.SendFile, error)
	AddStatus(ctx context.Context, details domain.Status) error
	StopFileShare(ctx context.Context, user_id string) error
	Upload(ctx context.Context, fileinfo domain.Files) error
	OpenedFolder(ctx context.Context, user_id string) (string, error)
	GetOrCreateHistory(ctx context.Context, userID string, folderID string) (string, error)
	UpdateHistory(ctx context.Context, user_id string, folder_id string) error
}

type filerepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) FileRepo {
	return &filerepo{db: db}
}

func (r *filerepo) GetOrCreateHistory(
	ctx context.Context,
	userID string,
	folderID string,
) (string, error) {

	// ensure row exists (insert if missing)
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO fetch_history (id, user_id, folder_id, last_delivered_at)
		 VALUES (?, ?, ?, '1970-01-01 00:00:00')
		 ON CONFLICT(user_id, folder_id) DO NOTHING;`,
		utils.IdGenerator(),
		userID,
		folderID,
	)
	if err != nil {
		log.Println("Error at insert :", err)
		return "", err
	}

	// now safely read
	var lastVisited string
	err = r.db.QueryRowContext(
		ctx,
		`SELECT last_delivered_at
		 FROM fetch_history
		 WHERE user_id = ? AND folder_id = ?;`,
		userID,
		folderID,
	).Scan(&lastVisited)

	if err != nil {
		log.Println("Error at getting lastvisit :", err)
		return "", err
	}

	return lastVisited, nil
}

func (r *filerepo) UpdateHistory(ctx context.Context, user_id string, folder_id string) error {
	_, err := r.db.Exec("UPDATE fetch_history SET last_delivered_at = CURRENT_TIMESTAMP WHERE user_id = ? AND folder_id = ?;", user_id, folder_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *filerepo) Get(
	ctx context.Context,
	folderID string,
	lastVisit string,
) ([]domain.SendFile, error) {

	var (
		rows *sql.Rows
		err  error
	)

	if lastVisit == "" {
		// first time: fetch all files
		rows, err = r.db.QueryContext(
			ctx,
			`SELECT file_id, type
			 FROM files
			 WHERE folder_id = ?`,
			folderID,
		)
	} else {
		// subsequent fetch: only new files
		rows, err = r.db.QueryContext(
			ctx,
			`SELECT file_id, type
			 FROM files
			 WHERE folder_id = ?
			   AND created_at > ?`,
			folderID,
			lastVisit,
		)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []domain.SendFile

	for rows.Next() {
		var file domain.SendFile
		if err := rows.Scan(&file.FileID, &file.Type); err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
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

func (r *filerepo) OpenedFolder(ctx context.Context, user_id string) (string, error) {

	var folderID string

	err := r.db.QueryRow("SELECT folder_id FROM upload_status WHERE user_id = ?", user_id).Scan(&folderID)
	if err != nil {
		return "", err
	}

	return folderID, nil

}

func (r *filerepo) Upload(ctx context.Context, fileinfo domain.Files) error {
	_, err := r.db.Exec("INSERT INTO files(id, file_id, type, folder_id, uploaded_by) VALUES (?, ?, ?, ?, ?);", fileinfo.ID, fileinfo.FileID, fileinfo.Type, fileinfo.FolderId, fileinfo.UploadedBy)
	if err != nil {
		return err
	}

	return nil
}
