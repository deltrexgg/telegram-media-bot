package service

import (
	"context"
	"errors"
	"strings"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
	"github.com/deltrexgg/telegram-media-bot/internal/repository"
	"github.com/deltrexgg/telegram-media-bot/internal/utils"
	"github.com/google/uuid"
)

type FileService interface {
	GetFiles(ctx context.Context, folder_id string) ([]domain.SendFile, error)
	OpenFolder(ctx context.Context, user_id string, folder_id string) error
	CloseFolder(ctx context.Context, user_id string) error
}

type fileservice struct {
	repo repository.FileRepo
}

func NewFileService(repo repository.FileRepo) FileService {
	return &fileservice{repo: repo}
}

func (s *fileservice) GetFiles(ctx context.Context, folder_id string) ([]domain.SendFile, error) {

	/*
		seperate /files from id and send the id to Get method
	*/

	parts := strings.Split(folder_id, ":")
	folder_id = parts[1]

	if _, err := uuid.Parse(folder_id); err != nil {
		return nil, errors.New("not a valid id")
	}

	return s.repo.Get(ctx, folder_id)
}

func (s *fileservice) GetLatestFiles(ctx context.Context, user_id string, folder_id string) ([]domain.SendFile, error){} 

func (s *fileservice) OpenFolder(ctx context.Context, userID string, folderID string) error {

	details := &domain.Status{
		ID:       utils.IdGenerator(),
		UserID:   userID,
		FolderID: folderID,
	}

	return s.repo.AddStatus(ctx, *details)
}

func (s *fileservice) CloseFolder(ctx context.Context, user_id string) error {
	err := s.repo.StopFileShare(ctx, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *fileservice) FileUpload(ctx context.Context, user_id string, file_id string, filetype string) (string, error) {

	folder_id, err :=
	fileinfo := &domain.Files{
		ID:         utils.IdGenerator(),
		FileID:     file_id,
		Type:       filetype,
		UploadedBy: user_id,
	}
}
