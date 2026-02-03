package handler

import (
	"context"
	"strconv"
	"strings"

	"github.com/deltrexgg/telegram-media-bot/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler struct {
	service service.FolderService
}

func NewFolderHandler(service service.FolderService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateFolder(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update == nil || update.Message == nil {
		return
	}

	text := update.Message.Text
	chatID := update.Message.Chat.ID

	parts := strings.Fields(text)

	if len(parts) == 1 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Mention folder name eg:/makefolder folderName",
		})
		return
	}

	folderName := parts[1]
	userId := strconv.Itoa(int(update.Message.From.ID))

	if err := h.service.AddFolder(ctx, folderName, userId); err != nil {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Folder Creation Error",
		})

		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Folder Created !",
	})

}

func (h *Handler) DeleteFolder(ctx context.Context, b *bot.Bot, update *models.Update) {
	/*
		List the folders
		User select the folder the corresponding id is return (from chat)
		manage the response
	*/
}

func (h *Handler) GetFolderList(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	userID := strconv.Itoa(int(update.Message.From.ID))

	intent := "view"
	if strings.HasPrefix(update.Message.Text, "/addfiles") {
		intent = "upload"
	}
	if strings.HasPrefix(update.Message.Text, "/latest") {
		intent = "newfiles"
	}
	if strings.HasPrefix(update.Message.Text, "/share") {
		intent = "link"
	}
	if strings.HasPrefix(update.Message.Text, "/delete") {
		intent = "remove"
	}

	buttonpad, err := h.service.Folderlist(ctx, userID, intent)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Failed",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        "Select a folder:",
		ReplyMarkup: buttonpad,
	})
}
