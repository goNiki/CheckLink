package linkreport

import (
	"context"
	"fmt"
	"goNiki/CheckLink/pkg/errorsAPP"

	"github.com/jung-kurt/gofpdf"
)

func (s *service) CreateReport(ctx context.Context, numbers []int64) error {

	linkBatch, err := s.linksstorage.GetByIDs(ctx, numbers)
	if err != nil {
		return fmt.Errorf("%w, %v", errorsAPP.ErrInternalDB, err)
	}

	report := gofpdf.New("P", "mm", "A4", "")
	report.SetFont("Arial", "", 14)
	report.AddPage()

	addHeader := func() {
		report.SetFont("Arial", "B", 16)
		report.CellFormat(0, 10, "Report", "", 1, "C", false, 0, "")
		report.Ln(5)
		report.SetFont("Arial", "", 12)
	}

	addHeader()

	colWidths := []float64{20, 120, 40}

	// Создание таблиц
	for _, table := range linkBatch {
		report.SetFont("Arial", "B", 14)
		//в будущем добавить в заголовок дату проверки
		title := fmt.Sprintf("Report with links_id: %d", table.Number)
		report.CellFormat(0, 10, title, "", 1, "L", false, 0, "")
		report.SetFont("Arial", "B", 12)

		// Добавление заголовок таблицы
		headers := []string{"№", "URL", "Validate"}
		for i, str := range headers {
			report.CellFormat(colWidths[i], 7, str, "1", 0, "C", false, 0, "")
		}
		report.Ln(-1)
		report.SetFont("Arial", "", 12)
		var rowID int64
		for url, status := range table.Links {
			rowID++
			if report.GetY() > 270 {
				report.AddPage()
				addHeader()
				report.SetFont("Arial", "B", 14)
				report.CellFormat(0, 10, title, "", 1, "L", false, 0, "")
				report.SetFont("Arial", "", 12)
				for i, str := range headers {
					report.CellFormat(colWidths[i], 7, str, "1", 0, "C", false, 0, "")
				}
				report.Ln(-1)
			}
			report.CellFormat(colWidths[0], 7, fmt.Sprintf("%d", rowID), "1", 0, "C", false, 0, "")
			report.CellFormat(colWidths[1], 7, url, "1", 0, "L", false, 0, "")
			report.CellFormat(colWidths[2], 7, status.String(), "1", 0, "L", false, 0, "")
			report.Ln(-1)
		}
		report.Ln(10)
	}

	err = report.OutputFileAndClose("report.pdf")
	if err != nil {
		return err
	}
	return nil

}
