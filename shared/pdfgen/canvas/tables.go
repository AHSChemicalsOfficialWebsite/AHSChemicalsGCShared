package canvas

import (
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/pdfgen/utils"
)

type TableHeader struct {
	X, Y        float64
	Headers     []string
	CellWidths  []float64
	Height      float64
	FillColor   [3]int
	TextColor   [3]int
	BorderColor [3]int
}

func (th *TableHeader) setHeight(c *Canvas, t *Text) {
	for i, cellWidth := range th.CellWidths {
		t.SetContent(th.Headers[i])
		lineHeight := t.GetMultiTextHeight(c.PDF, cellWidth)
		th.Height = max(th.Height, lineHeight) + 1 // 1 is the padding
	}
}

func (th *TableHeader) Draw(c *Canvas, t *Text) {
	th.setHeight(c, t)
	t.SetColor(th.TextColor)
	x := th.X
	y := th.Y

	for i, header := range th.Headers {
		cellWidth := th.CellWidths[i]

		c.DrawRectangle(&Rectangle{ //Header Colored Rectangle
			X:         x,
			Y:         y,
			Width:     cellWidth,
			Height:    th.Height,
			Style:     "F",
			LineWidth: 0.8,
			FillColor: th.FillColor,
		})

		t.SetContent(header)
		blockHeight := t.GetMultiTextHeight(c.PDF, cellWidth)
		//Center the text with 1 padding 
		t.SetY(y + (th.Height - blockHeight)/2 + t.GetTextHeight(c.PDF) - 0.5)
		t.SetX(x)
		c.DrawMultipleLines(t, cellWidth, "C")
		x += cellWidth
	}
	c.DrawRectangle(&Rectangle{ //Outer border for the table
		X:           th.X,
		Y:           th.Y,
		Width:       utils.CalculateShippingTableCellWidths(th.CellWidths),
		Height:      th.Height,
		LineWidth:   0.8,
		BorderColor: th.BorderColor,
	})
}

type TableBody struct {
	X, Y        float64
	Rows        [][]string
	CellWidths  []float64
	Height      float64
	TextColor   [3]int
	BorderColor [3]int
}

func (tb *TableBody) DrawTableCellsRightBorder(c *Canvas) {
	x := tb.X
	for i := range len(tb.CellWidths) - 1 {
		x += tb.CellWidths[i]
		c.DrawLine(&Line{
			X1:    x,
			Y1:    tb.Y,
			X2:    x,
			Y2:    tb.Y + tb.Height,
			Color: tb.BorderColor,
			Width: 0.8,
		})
	}
}

func (tb *TableBody) Draw(c *Canvas, t *Text) {
	x := tb.X
	y := tb.Y + 3
	t.SetSize(9)
	t.SetStyle("")
	t.SetColor(tb.TextColor)

	for _, row := range tb.Rows {
		var rowHeight float64 = 0

		// First pass to determine rowHeight
		for j, cell := range row {
			width := tb.CellWidths[j]
			t.SetContent(cell)
			rowHeight = max(rowHeight, t.GetMultiTextHeight(c.PDF, width))
		}

		rowHeight += 2.5

		// Check for page overflow BEFORE drawing the row
		if y + 10 > c.BorderHeight {
			// Finish current table on this page
			c.DrawRectangle(&Rectangle{
				X:           tb.X,
				Y:           tb.Y,
				Width:       utils.CalculateShippingTableCellWidths(tb.CellWidths),
				Height:      tb.Height,
				LineWidth:   0.8,
				BorderColor: tb.BorderColor,
			})
			tb.DrawTableCellsRightBorder(c)

			// Add new page
			c.PDF.AddPage()

			// Redraw borders for new page
			c.DrawRectangle(&Rectangle{
				X:           c.BorderX,
				Y:           c.BorderY,
				Width:       c.BorderWidth,
				Height:      c.BorderHeight,
				LineWidth:   0.8,
				BorderColor: tb.BorderColor,
			})

			// Reset positions
			c.MoveTo(c.MarginLeft, c.MarginTop)
			tb.X = c.X
			tb.Y = c.Y
			x = tb.X
			y = tb.Y + 3
			tb.Height = 0
		}

		// Draw the row
		for j, cell := range row {
			t.SetContent(cell)
			width := tb.CellWidths[j]
			t.SetX(x)
			t.SetY(y + (rowHeight - t.GetMultiTextHeight(c.PDF, width))/2)
			c.DrawMultipleLines(t, width, "C")
			x += width
		}

		// Reset x, and increment y and height
		x = tb.X
		y += rowHeight
		tb.Height += rowHeight
	}

	// Draw the final table border on the last page
	c.DrawRectangle(&Rectangle{
		X:           tb.X,
		Y:           tb.Y,
		Width:       utils.CalculateShippingTableCellWidths(tb.CellWidths),
		Height:      tb.Height,
		LineWidth:   0.8,
		BorderColor: tb.BorderColor,
	})
	tb.DrawTableCellsRightBorder(c)
}

type Table struct {
	Header *TableHeader
	Body   *TableBody
	Width  float64
}

// Returns the y position where the table ends
func (tb *Table) Draw(c *Canvas, t *Text) float64 {
	tb.Header.Draw(c, t)

	tb.Body.Y = tb.Header.Y + tb.Header.Height
	tb.Body.Height += 1
	tb.Body.Draw(c, t)

	return tb.Body.Y + tb.Body.Height
}
