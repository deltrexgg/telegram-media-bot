package service

import (
	"context"

	"github.com/deltrexgg/telegram-media-bot/internal/repository"
)

type FolderService interface {
	AddFolder(ctx context.Context) error
}

type service struct {
	repo repository.FolderRepo
}

func FolderServiceMethod(repo repository.FolderRepo) FolderService {
	return &service{repo: repo}
}

func (s *service) AddFolder(ctx context.Context) error {
	return s.repo.Create(ctx)
}
