package rest

import (
	"github.com/zoglam/pdf-sender-bot/internal/service"
)

type Rest struct {
	PDFService  service.PDFService
	userService service.UserService
}

func NewRest(
	pdfService service.PDFService,
	userService service.UserService,
) *Rest {
	return &Rest{
		PDFService:  pdfService,
		userService: userService,
	}
}
