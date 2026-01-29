package handler

import (
	"context"
	"strconv"
	"strings"

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
	UserId := strconv.Itoa(int(update.Message.From.ID))
	Text := update.Message.Text

	parts := strings.Fields(Text)

	err := h.service.OpenFolder(ctx, UserId, parts[1])
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text: "Error in opening the folder",
		})

		return
	}


	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Send files now in the chat after finishing type /stop to close the folder",
	})
}

func (h *FileHandler) StopShare(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	UserId := strconv.Itoa(int(update.Message.From.ID))

	err := h.service.CloseFolder(ctx, UserId)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text: "Cant close the folder , try again later.",
		})
		return
	} 

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text: "Folder closed.",
	})
}
