package service

import (
	"context"
	"errors"
	"strings"

	"github.com/deltrexgg/telegram-media-bot/internal/repository"
	"github.com/google/uuid"
)

type FileService interface {
	GetFiles(ctx context.Context, folder_id string) ([]string, error)
}

type fileservice struct {
	repo repository.FileRepo
}

func NewFileService(repo repository.FileRepo) FileService {
	return &fileservice{repo: repo}
}

func (s *fileservice) GetFiles(ctx context.Context, folder_id string) ([]string, error) {

	/*
		seperate /files from id and send the id to Get method
	*/

	parts := strings.Split(folder_id, " ")
	folder_id = parts[1]

	if _, err := uuid.Parse(folder_id); err != nil {
		return nil, errors.New("not a valid id")
	}

	return s.repo.Get(ctx, folder_id)
}
