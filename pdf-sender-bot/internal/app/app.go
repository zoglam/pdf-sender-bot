package app

import (
	"github.com/zoglam/pdf-sender-bot/internal/service"
	"github.com/zoglam/pdf-sender-bot/pkg/telegram"
)

const (
	startMessage = `- Для получения ссылки на личный кабинет /start
- Получить путевой лист, если заполнен профиль /pdf`
	commandNotFound = "❌Команда не найдена. Вызовите повторно /start"
)

type App struct {
	PDFService  service.PDFService
	bot         *telegram.Bot
	userService service.UserService
}

func NewApp(
	pdfService service.PDFService,
	bot *telegram.Bot,
	userService service.UserService,
) *App {
	return &App{
		PDFService:  pdfService,
		bot:         bot,
		userService: userService,
	}
}
