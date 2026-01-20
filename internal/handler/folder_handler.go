package handler

import (
	"context"

	"github.com/deltrexgg/telegram-media-bot/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler struct {
	service service.FolderService
}

func FolderHandler(service service.FolderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateFolder(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update == nil || update.Message == nil {
		return
	}

	msg := update.Message
	if err := h.service.AddFolder(ctx); err != nil {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   "Folder Creation Error",
		})

		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: msg.Chat.ID,
		Text:   "Folder Created !",
	})

}
