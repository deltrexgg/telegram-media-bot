package repository

import (
	"context"

	"github.com/deltrexgg/telegram-media-bot/internal/config"
)

type FolderRepo interface {
	Create(ctx context.Context, foldername string, userid string, id string) error
}

type repo struct{}

func NewFolderRepo() FolderRepo {
	return &repo{}
}

func (r *repo) Create(ctx context.Context, foldername string, userid string, id string) error {

	db := config.DB
	_, err := db.Exec("INSERT INTO folders(id, created_by, name) VALUES(?, ?, ?)", id, userid, foldername)
	if err != nil {
		return err
	}

	return nil
}
