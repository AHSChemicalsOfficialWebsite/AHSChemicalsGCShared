package canvas

//Rectangle drawn in the pdf.
type Rectangle struct {
	X           float64
	Y           float64
	Width       float64
	Height      float64
	Style       string
	BorderColor [3]int
	FillColor   [3]int
	LineWidth   float64
}

func (c *Canvas) DrawRectangle(r *Rectangle){
	c.PDF.SetDrawColor(r.BorderColor[0], r.BorderColor[1], r.BorderColor[2])
	c.PDF.SetLineWidth(r.LineWidth)
	c.PDF.SetFillColor(r.FillColor[0], r.FillColor[1], r.FillColor[2])
	c.PDF.Rect(r.X, r.Y, r.Width, r.Height, r.Style)
}
