package canvas

import (
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/pdfgen/utils"
)

type TableHeader struct {
	X, Y            float64
	Headers         []string
	CellWidths      []float64
	Height          float64
	FillColor       [3]int
	TextColor       [3]int
	BorderColor     [3]int
	BorderThickness float64
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
			LineWidth: th.BorderThickness,
			FillColor: th.FillColor,
		})

		t.SetContent(header)
		blockHeight := t.GetMultiTextHeight(c.PDF, cellWidth)
		//Center the text with -0.5 padding to account for the font
		t.SetY(y + (th.Height-blockHeight)/2 + t.GetTextHeight(c.PDF) - 0.5)
		t.SetX(x)
		c.DrawMultipleLines(t, cellWidth, "C")
		x += cellWidth
	}
	//Draw the final outer border
	c.DrawRectangle(&Rectangle{
		X:           th.X,
		Y:           th.Y,
		Width:       utils.CalculateTableCellWidths(th.CellWidths),
		Height:      th.Height,
		LineWidth:   th.BorderThickness,
		BorderColor: th.BorderColor,
	})
}

type TableBody struct {
	X, Y            float64
	Rows            [][]string
	CellWidths      []float64
	Height          float64
	TextColor       [3]int
	BorderColor     [3]int
	BorderThickness float64
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
			Width: tb.BorderThickness,
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
		if y+rowHeight > c.BorderHeight {
			// Finish current table on this page
			c.DrawRectangle(&Rectangle{
				X:           tb.X,
				Y:           tb.Y,
				Width:       utils.CalculateTableCellWidths(tb.CellWidths),
				Height:      tb.Height,
				LineWidth:   tb.BorderThickness,
				BorderColor: tb.BorderColor,
			})
			tb.DrawTableCellsRightBorder(c)

			// Add new page
			c.PDF.AddPage()

			// Redraw the outer borders for the new page
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

		// Finally draw the row
		for j, cell := range row {
			t.SetContent(cell)
			width := tb.CellWidths[j]
			t.SetX(x)
			t.SetY(y + (rowHeight-t.GetMultiTextHeight(c.PDF, width))/2)
			c.DrawMultipleLines(t, width, "C")
			x += width
		}
		// Reset x, and increment y and height to draw the next row
		x = tb.X
		y += rowHeight
		tb.Height += rowHeight
	}
	// Draw the final table border on the last page
	c.DrawRectangle(&Rectangle{
		X:           tb.X,
		Y:           tb.Y,
		Width:       utils.CalculateTableCellWidths(tb.CellWidths),
		Height:      tb.Height,
		LineWidth:   tb.BorderThickness,
		BorderColor: tb.BorderColor,
	})
	tb.DrawTableCellsRightBorder(c)
}

type TableCell struct {
	Lines []string
	Width float64
}

type TableRow struct {
	Cells []TableCell
}

type TableBody2 struct {
	X, Y            float64
	Height          float64
	TextColor       [3]int
	BorderColor     [3]int
	CellWidths      []float64
	BorderThickness float64
	Rows            []TableRow
}

// TODO: Not complete yet. Y position gets weird when drawing the next row. Also draws the lines of the second last 
// column with 4 line gap for some reason. Will fix this later
func (tb *TableBody2) Draw(c *Canvas, t *Text) {
	initialX := tb.X
	initialY := tb.Y + 5
	yTracker := initialY
	t.SetSize(9)
	t.SetStyle("")
	t.SetColor(tb.TextColor)

	for _, row := range tb.Rows {
		x := initialX
		startY := yTracker
		var maxCellHeight float64 = 0.0

		// Draw each cell
		for _, cell := range row.Cells {
			lineHeight := t.GetMultiTextHeight(c.PDF, cell.Width)
			cellHeight := float64(len(cell.Lines)) * (lineHeight + 1)
			if cellHeight > maxCellHeight {
				maxCellHeight = cellHeight
			}

			for _, line := range cell.Lines {
				if yTracker+lineHeight > c.BorderHeight {
					// Page break
					c.PDF.AddPage()
					c.DrawRectangle(&Rectangle{
						X:           c.BorderX,
						Y:           c.BorderY,
						Width:       c.BorderWidth,
						Height:      c.BorderHeight,
						LineWidth:   0.8,
						BorderColor: tb.BorderColor,
					})
					c.MoveTo(c.MarginLeft, c.MarginTop)
					x = c.X
					yTracker = c.Y + 5
					startY = yTracker
				}
				t.SetContent(line)
				t.SetX(x)
				t.SetY(yTracker)
				c.DrawMultipleLines(t, cell.Width, "C")
				yTracker += lineHeight + 1
			}

			// Reset Y to start of row for next cell
			yTracker = startY
			x += cell.Width
		}

		yTracker += maxCellHeight + 5 
	}
}

// Represents a complete table with header and body and width.
type Table struct {
	Header *TableHeader
	Body   *TableBody
	Width  float64
}

// Draws the entire table. Returns the y position where the table ends
func (tb *Table) Draw(c *Canvas, t *Text) float64 {
	tb.Header.Draw(c, t)

	tb.Body.Y = tb.Header.Y + tb.Header.Height
	tb.Body.Height += 1
	tb.Body.Draw(c, t)

	return tb.Body.Y + tb.Body.Height
}
