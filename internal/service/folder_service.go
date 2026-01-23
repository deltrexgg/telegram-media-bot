package service

import (
	"context"

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
	id := utils.IdGenerator()
	return s.repo.Create(ctx, foldername, userid, id)
}
