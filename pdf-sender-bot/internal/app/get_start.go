package app

import "gopkg.in/telebot.v3"

func (t *App) GetStart(c telebot.Context) error {
	return c.Send(startMessage, t.bot.GetWebAppButton())
}
