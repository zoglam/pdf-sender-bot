package app

import (
	"bytes"
	"time"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"gopkg.in/telebot.v3"
)

func (t *App) GetPDF(c telebot.Context) error {
	data := &dto.PDFParams{
		QR:                        "string",
		Organization:              "string",
		Address:                   "string",
		Phone:                     "string",
		OGRN:                      "string",
		ModelT:                    "string",
		GovermentSign:             "string",
		FirstName:                 "string",
		SecondName:                "string",
		ThirdName:                 "string",
		СertificateNumber:         "string",
		LicenseRegistrationNumber: "string",
		LicenseSerial:             "string",
		LicenseNumber:             "string",
		GarageNumber:              "string",
		PersonnelNumber:           "string",
		ShortSignFIO:              "string",
	}
	pdf, err := t.PDFService.GeneratePDF(data)
	if err != nil {
		logg.Error().Msg("Ошибка при генерации pdf")
	}

	reader := bytes.NewReader(pdf)
	file := &telebot.Document{
		File:     telebot.FromReader(reader),
		FileName: time.Now().Format("2006-01-02") + "-путевой-лист.pdf",
	}
	return c.Send(file)
}
