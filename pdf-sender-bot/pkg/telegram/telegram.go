package telegram

import (
	"context"
	"time"

	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"gopkg.in/telebot.v3"
)

var (
	webAppURL         string
	middlewareMetrics func(endpoint string) telebot.MiddlewareFunc
)

type Bot struct {
	Bot  *telebot.Bot
	Menu *telebot.ReplyMarkup
}

func InitBot(token string) *Bot {
	ctx := context.Background()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		logg.Fatal(ctx).Msgf("%v", err)
	}

	return &Bot{
		Bot:  bot,
		Menu: &telebot.ReplyMarkup{ResizeKeyboard: true},
	}
}

func (b *Bot) RegisterWebAppURL(url string) {
	webAppURL = url
}

func (b *Bot) RegisterMiddleWareMetrics(middleware func(endpoint string) telebot.MiddlewareFunc) {
	middlewareMetrics = middleware
}

func (b *Bot) RegisterHandler(endpoint string, handler func(telebot.Context) error) {
	b.Bot.Handle(endpoint, handler, middlewareMetrics(endpoint))
}

func (b *Bot) RegisterMenuElements(texts []string) {

	rows := make([]telebot.Row, 0, len(texts))
	for _, t := range texts {
		rows = append(rows, b.Menu.Row(b.Menu.Text(t)))
	}

	b.Menu.Reply(rows...)
}

func (b *Bot) GetWebAppButton() *telebot.ReplyMarkup {
	selector := &telebot.ReplyMarkup{}

	webApp := &telebot.WebApp{
		URL: webAppURL,
	}

	selector.Inline(
		selector.Row(
			selector.WebApp("Личный кабинет", webApp),
		),
	)

	return selector
}
