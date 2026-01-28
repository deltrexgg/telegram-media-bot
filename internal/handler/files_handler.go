package handler

import (
	"context"

	"github.com/deltrexgg/telegram-media-bot/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type FileHandler struct {
	service service.FileService
}

func NewFileHandler(service service.FileService) *FileHandler {
	return &FileHandler{service: service}
}

func (h *FileHandler) SendFiles(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery == nil {
		return
	}

	folder_id := update.CallbackQuery.Data
	chatID := update.CallbackQuery.From.ID

	filesid, err := h.service.GetFiles(ctx, folder_id)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Failed to get",
		})
		return
	}

	if filesid == nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "No Files in the folder",
		})

		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   folder_id,
	})
}

func (h *FileHandler) AddFiles(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Files added successfully",
	})
}
