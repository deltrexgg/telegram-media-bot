package service

import (
	"context"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
	"github.com/deltrexgg/telegram-media-bot/internal/repository"
	"github.com/deltrexgg/telegram-media-bot/internal/utils"
)

type FolderService interface {
	AddFolder(ctx context.Context, foldername string, userid string) error
}

type service struct {
	repo repository.FolderRepo
}

func NewFolderService(repo repository.FolderRepo) FolderService {
	return &service{repo: repo}
}

func (s *service) AddFolder(ctx context.Context, foldername string, userid string) error {

	folder := domain.Folders{
		ID:        utils.IdGenerator(),
		Name:      foldername,
		CreatedBy: userid,
	}
	return s.repo.Create(ctx, folder)
}
