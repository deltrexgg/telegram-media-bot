package service

import (
	"context"
	"errors"

	"github.com/deltrexgg/telegram-media-bot/internal/domain"
	"github.com/deltrexgg/telegram-media-bot/internal/repository"
	"github.com/deltrexgg/telegram-media-bot/internal/utils"
	"github.com/go-telegram/bot/models"
	"github.com/google/uuid"
)

type FolderService interface {
	AddFolder(ctx context.Context, foldername string, userid string) error
	RemoveFolder(ctx context.Context, id string) error
	Folderlist(ctx context.Context, userId string, text string) (*models.InlineKeyboardMarkup, error)
}

type service struct {
	repo repository.FolderRepo
}

func NewFolderService(repo repository.FolderRepo) FolderService {
	return &service{repo: repo}
}

func (s *service) AddFolder(ctx context.Context, foldername string, userid string) error {

	if len(foldername) > 10 {
		return errors.New("Folder length exceeds 10 character")
	}
	folder := domain.Folders{
		ID:        utils.IdGenerator(),
		Name:      foldername,
		CreatedBy: userid,
	}

	err := s.repo.Create(ctx, folder)
	if err != nil {
		return err
	}

	access := domain.Access{
		ID:       utils.IdGenerator(),
		FolderID: folder.ID,
		UserID:   userid,
	}

	return s.repo.GrandFolderAccess(ctx, access)

}

func (s *service) RemoveFolder(ctx context.Context, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New("not a valid id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *service) Folderlist(ctx context.Context, userId string, text string) (*models.InlineKeyboardMarkup, error) {

	var tag string

	switch text {
	case "/folders":
		tag = "/getfiles"

	case "/addfiles":
		tag = "/addfiles"
	}

	folders, err := s.repo.FolderNameByUserId(ctx, userId)
	if err != nil {
		return &models.InlineKeyboardMarkup{}, err
	}

	keyboard := [][]models.InlineKeyboardButton{}

	for _, f := range folders {
		keyboard = append(keyboard, []models.InlineKeyboardButton{
			{
				Text:         f.Name,
				CallbackData: tag + " " + f.ID, //identify the callback function
			},
		})
	}

	buttonPad := &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}

	return buttonPad, nil
}
