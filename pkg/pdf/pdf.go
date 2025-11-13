package pdf

import (
	"bytes"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"
	"io"

	"github.com/jung-kurt/gofpdf"
)

const ()

type pdf struct {
	pdf *gofpdf.Fpdf
}

func New() *pdf {
	r := gofpdf.New("p", "mm", "A4", "")
	r.SetFont("Arial", "", 14)
	r.AddPage()
	return &pdf{
		pdf: r,
	}
}

func (p *pdf) OutputFileAndClose() error {
	err := p.pdf.OutputFileAndClose("report.pdf")
	if err != nil {
		return err
	}
	return nil
}

func (p *pdf) OutputToReader() (io.Reader, error) {

	var buf bytes.Buffer
	err := p.pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errorsAPP.ErrOutputPdf, err)
	}
	return &buf, nil
}
