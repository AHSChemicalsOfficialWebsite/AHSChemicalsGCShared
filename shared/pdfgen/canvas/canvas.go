package canvas

import (
	"fmt"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/company_details"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/constants"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
	"github.com/phpdave11/gofpdf"
)

var (
	PrimaryBlue          = [3]int{65, 83, 145}
	PrimaryGreen         = [3]int{165, 199, 89}
	White                = [3]int{255, 255, 255}
	Black                = [3]int{0, 0, 0}
	ShippingTableHeaders = []string{"REQUISITIONER", "SHIP VIA", "F.O.B", "SHIPPING TERMS"}
	ShippingTableValues  = [][]string{{"Robert Vodka", "In House", "Factory", "N/A"}}
	ProductTableHeaders  = []string{"SKU", "DESCRIPTION", "QTY", "PRICE", "TOTAL"}
)

type Canvas struct {
	PDF              *gofpdf.Fpdf
	X, Y             float64
	BorderX, BorderY float64 // Border starting point on page
	BorderWidth      float64
	BorderHeight     float64
	MarginLeft       float64 // Margin starting from the border
	MarginTop        float64
}

// Creates a pointer to new canvas
//
// Note:
//   - BorderX and BorderY are default positions at which the border is drawn for the pdf. By default I set the values to (8, 8)
//   - BorderWidth and BorderHeight are the dimensions of the border drawn for the pdf. By default I set the values to (193, 280)
//   - MarginLeft and MarginTop are the margins from the border. By default I set the values to (15, 15)
func NewCanvas(pdf *gofpdf.Fpdf) *Canvas {
	return &Canvas{
		PDF:          pdf,
		X:            0,
		Y:            0,
		BorderX:      8,
		BorderY:      8,
		BorderWidth:  193,
		BorderHeight: 280,
		MarginLeft:   15,
		MarginTop:    15,
	}
}

/* Setters */

func (c *Canvas) SetBorderX(x float64)      { c.X = x }
func (c *Canvas) SetBorderY(y float64)      { c.Y = y }
func (c *Canvas) SetMarginLeft(x float64)   { c.MarginLeft = x }
func (c *Canvas) SetMarginTop(y float64)    { c.MarginTop = y }
func (c *Canvas) SetBorderWidth(w float64)  { c.BorderWidth = w }
func (c *Canvas) SetBorderHeight(h float64) { c.BorderHeight = h }

/* Position helpers */

func (c *Canvas) MoveTo(x, y float64) { c.X = x; c.Y = y }
func (c *Canvas) IncX(dx float64)     { c.X += dx }
func (c *Canvas) IncY(dy float64)     { c.Y += dy }
func (c *Canvas) DecX(dx float64)     { c.X -= dx }
func (c *Canvas) DecY(dy float64)     { c.Y -= dy }
func (c *Canvas) ResetX()             { c.X = c.MarginLeft }
func (c *Canvas) ResetY()             { c.Y = c.MarginTop }

/* Reusable Draw Functions */

func (c *Canvas) AddNewPageIfEnd(offest float64, borderColor [3]int, lineWidth float64) {
	if c.Y + offest > c.BorderHeight{
		c.PDF.AddPage()
		c.DrawRectangle(&Rectangle{
			X:           c.BorderX,
			Y:           c.BorderY,
			Width:       c.BorderWidth,
			Height:      c.BorderHeight,
			LineWidth:   0.8,
			BorderColor: PrimaryBlue,
		})
		c.MoveTo(c.MarginLeft, c.MarginTop)
	}
}

func (c *Canvas) DrawCompanyDetails() {
	lines := []struct {
		text  string
		style string
	}{
		{constants.CompanyName, "B"},
		{company_details.COMPANYADDRESSLINE1, ""},
		{company_details.COMPANYADDRESSLINE2, ""},
		{"Phone: " + company_details.COMPANYPHONE, ""},
		{"Email: " + company_details.COMPANYEMAIL, ""},
		{"Website: " + constants.CompanyName, ""},
	}

	for _, line := range lines {
		c.DrawSingleLineText(&Text{
			Content: line.text,
			Font:    "Arial",
			Style:   line.style,
			Size:    10,
			X:       c.X,
			Y:       c.Y,
			Color:   Black,
		})
		c.IncY(5)
	}
}

func (c *Canvas) DrawCustomerDetails(customer *models.Customer) {
	lines := []struct {
		text  string
		style string
	}{
		{customer.Name, "B"},
		{customer.Address1, ""},
		{customer.FormatAddress2(), ""},
		{"Phone: " + customer.Phone, ""},
		{"Email: " + customer.Email, ""},
	}

	for _, line := range lines {
		c.DrawSingleLineText(&Text{
			Content: line.text,
			Font:    "Arial",
			Style:   line.style,
			Size:    10,
			X:       c.X,
			Y:       c.Y,
			Color:   Black,
		})
		c.IncY(5)
	}
}

func (c *Canvas) DrawFooter(text string) {
	textElement := &Text{
		Content: text,
		Font:    "Arial",
		Style:   "",
		Size:    8,
		X:       c.BorderX,
		Y:       c.BorderY + c.BorderHeight - 3,
		Color:   Black,
	}
	c.DrawMultipleLines(textElement, c.BorderWidth, "C")
}

func (c *Canvas) DrawBillingDetails(values []string, taxRate string) {
	labels := []string{"SUBTOTAL", fmt.Sprintf("TAX (%s)", taxRate), "TOTAL"}
	initialX := c.X
	var valueStyle = ""
	for i, label := range labels {
		c.DrawSingleLineText(&Text{
			Content: label,
			Font:    "Arial",
			Style:   "",
			Size:    10,
			X:       c.X,
			Y:       c.Y,
			Color:   Black,
		})
		c.IncX(32)
		if i == 2 {
			valueStyle = "B"
		}
		c.DrawSingleLineText(&Text{
			Content: values[i],
			Font:    "Arial",
			Style:   valueStyle,
			Size:    10,
			X:       c.X,
			Y:       c.Y,
			Color:   Black,
		})
		c.MoveTo(initialX, c.Y + 5)
	}
}
