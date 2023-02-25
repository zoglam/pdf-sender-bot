package pdf

import (
	"bytes"
	"fmt"
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
		fmt.Println(1, err)
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println(2, err)
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

	err = pdfg.Create()
	if err != nil {
		fmt.Println(3, err)
		return nil, err
	}

	return pdfg.Bytes(), nil
}
