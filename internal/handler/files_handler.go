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

func (h *FileHandler) CallBackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	parts := strings.Split(update.CallbackQuery.Data, ":")
	action := parts[0]

	switch action {
	case "view":
		// call ShowFiles
		h.ShowFiles(ctx, b, update)
	case "upload":
		//call StartShare
		h.StartShare(ctx, b, update)
	}


}


//called in callback function for viewing files
func (h *FileHandler) ShowFiles(ctx context.Context, b *bot.Bot, update *models.Update) {
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


//called in callback function
func (h *FileHandler) StartShare(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.CallbackQuery == nil {
		return
	}

	chatID := update.CallbackQuery.From.ID
	userID := strconv.Itoa(int(update.CallbackQuery.From.ID))

	// callback data format: "upload:<folder_id>"
	data := update.CallbackQuery.Data
	parts := strings.Split(data, ":")

	if len(parts) != 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Invalid folder selection",
		})
		return
	}

	folderID := parts[1]

	if err := h.service.OpenFolder(ctx, userID, folderID); err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Error opening folder",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Send files now. Type /stop when finished sending files to close the folder.",
	})

	// always answer callback
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
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
