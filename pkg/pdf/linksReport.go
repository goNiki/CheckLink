package pdf

import (
	"fmt"
	"goNiki/CheckLink/internal/domain"
)

func (p *pdf) CreateRepors(nameReport string, links []domain.LinkBatch) {
	cw := []float64{20, 120, 40}
	th := []string{"id", "URL", "Validate"}

	p.ReportHeader(nameReport)

	for _, link := range links {
		nameTable := fmt.Sprintf("Report with Links_id: %d", link.Number)
		p.CreateTable(&link, cw, nameTable, th)
	}

}

func (p *pdf) CreateTable(l *domain.LinkBatch, cw []float64, name string, headers []string) {
	p.TableName(name)
	p.TableHeader(headers, cw)
	id := 0
	for u, s := range l.Links {
		if p.NextPage() {
			p.TableHeader(headers, cw)
		}
		id++
		idstr := fmt.Sprintf("%d", id)
		d := []string{idstr, u, s.String()}
		p.TableLine(d, cw)
	}

	p.pdf.Ln(10)

}
