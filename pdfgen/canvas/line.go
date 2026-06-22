package canvas

// Line represents a straight line with given width and color drawn between (x1,y1) and (x2,y2)
type Line struct {
	X1    float64
	Y1    float64
	X2    float64
	Y2    float64
	Width float64
	Color [3]int
}

func (c *Canvas) DrawLine(l *Line) {
	c.PDF.SetLineWidth(l.Width)
	c.PDF.SetDrawColor(l.Color[0], l.Color[1], l.Color[2])
	c.PDF.Line(l.X1, l.Y1, l.X2, l.Y2)
}