package pdf

import (
	"bytes"
	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/zoglam/pdf-sender-bot/internal/dto"
)

type PDFService struct{}

func NewPDF() *PDFService {
	return &PDFService{}
}

func (p PDFService) GeneratePDF(data *dto.PDFParams) ([]byte, error) {
	var templ *template.Template
	var err error

	if templ, err = template.ParseFiles("config/template.html"); err != nil {
		return nil, err
	}

	var body bytes.Buffer
	if err = templ.Execute(&body, data); err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.MarginRight.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}
	
	return pdfg.Bytes(), nil
}
