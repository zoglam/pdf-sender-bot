package app

import "github.com/zoglam/pdf-sender-bot/internal/service"

type App struct {
	PDFService service.PDFService
}

func NewApp(pdfService service.PDFService) *App {
	return &App{
		PDFService: pdfService,
	}
}
