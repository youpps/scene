package scenes

import (
	"errors"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getUserId(update tg.Update) (int64, error) {
	if update.SentFrom() != nil {
		return update.SentFrom().ID, nil
	}
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID, nil
	}
	if update.Message != nil {
		return update.Message.From.ID, nil
	}
	return 0, errors.New("id hasn' t been found")
}
