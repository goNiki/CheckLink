package pdf

func (p *pdf) ReportHeader(t string) {
	p.TextH1()
	p.pdf.CellFormat(0, 10, t, "", 1, "C", false, 0, "")
	p.pdf.Ln(5)
}

func (p *pdf) TableName(t string) {
	p.TextH2()
	p.pdf.CellFormat(0, 10, t, "", 1, "L", false, 0, "")
}

func (p *pdf) TableHeader(h []string, cw []float64) {
	p.TextH3()
	for i, str := range h {
		p.pdf.CellFormat(cw[i], 7, str, "1", 0, "C", false, 0, "")
	}
	p.pdf.Ln(-1)
}

func (p *pdf) TableLine(d []string, cw []float64) {
	p.TextBase()
	for i, v := range d {
		p.pdf.CellFormat(cw[i], 7, v, "1", 0, "L", false, 0, "")
	}
	p.pdf.Ln(-1)
}

func (p *pdf) NextPage() bool {
	if p.pdf.GetY() > 270 {
		p.pdf.AddPage()
		return true
	}
	return false
}
