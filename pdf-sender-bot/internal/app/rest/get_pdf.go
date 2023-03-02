package rest

import (
	"net/http"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/pkg/logg"
)

func (a *Rest) GetPDF(w http.ResponseWriter, r *http.Request) {
	logg.Info().Msg("Пришел запрос")
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

	pdfBytes, err := a.PDFService.GeneratePDF(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// w.Header().Set("Content-Disposition", "attachment; filename=kittens.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}
