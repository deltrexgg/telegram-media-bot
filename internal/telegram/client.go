/*
	InitBot function creates and returns a configured Telegram bot instance.
	It does not start the bot or manage application lifecycle.
*/

package telegram

import (
	"errors"

	"github.com/go-telegram/bot"
)

func InitBot(apiKey string, opts ...bot.Option) (*bot.Bot, error) {
	if apiKey == "" {
		return nil, errors.New("telegram api key is empty")
	}

	b, err := bot.New(apiKey, opts...)
	if err != nil {
		return nil, err
	}

	return b, nil
}
