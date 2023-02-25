package app

import (
	"net/http"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
)

func (a *App) GetPDF(w http.ResponseWriter, r *http.Request) {

	data := &dto.PDFParams{
		Name:    "Anzor",
		Surname: "Baysarov",
		Age:     15,
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
