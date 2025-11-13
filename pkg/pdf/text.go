package pdf

func (p *pdf) TextBase() {
	p.pdf.SetFont("Arial", "", 12)
}

func (p *pdf) TextH1() {
	p.pdf.SetFont("Arial", "B", 16)
}

func (p *pdf) TextH2() {
	p.pdf.SetFont("Arial", "B", 14)
}

func (p *pdf) TextH3() {
	p.pdf.SetFont("Arial", "B", 12)
}
