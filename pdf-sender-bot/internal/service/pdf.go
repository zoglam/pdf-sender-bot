package service

import (
	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/pkg/pdf"
)

type PDFService interface {
	GeneratePDF(data *dto.PDFParams) ([]byte, error)
}

func NewPDFService() PDFService {
	return pdf.NewPDF()
}
