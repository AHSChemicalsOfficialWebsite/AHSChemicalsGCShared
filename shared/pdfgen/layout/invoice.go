package layout

import (
	"fmt"
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/company_details"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/models"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/pdfgen/canvas"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/pdfgen/utils"
	"github.com/phpdave11/gofpdf"
)

type Invoice struct {
	Number string
	Items  []models.Product
	Customer *models.Customer
	TableValues [][]string
	Total string
	SubTotal string
	TaxAmount string
	TaxRate string
	PaymentDue string
	CreatedAt string
}

const (
	TermsAndConditions = "Payment is due within 30 days from the invoice date (Net 30). A 10% late fee will be automatically applied to the total outstanding balance if payment is not received within 14 days after the due date. Continued non-payment may result in suspension of services and additional collection actions. By receiving this invoice, you agree to these terms."
)
var (
	invoiceTableHeaders = []string{"ITEM", "QUANTITY", "PRICE PER UNIT", "AMOUNT"}
	invoiceTableColWidths = []float64{75, 25, 40, 40}
)

func NewInvoice(order *models.Order, invoiceNumber string) *Invoice{
	invoice := &Invoice{
		Number: invoiceNumber,
		Customer: &order.Customer,
		Total: order.GetFormattedTotal(),
		SubTotal: order.GetFormattedSubTotal(),
		TaxAmount: order.GetFormattedTaxAmount(),
		TaxRate: order.GetFormattedTaxRate(),
		CreatedAt: time.Now().Format("January 2, 2006"),
		PaymentDue: time.Now().AddDate(0, 0, 30).Format("January 2, 2006"),
	}
	invoice.setTableValues(order.Items)

	return invoice
}

func (i *Invoice) setTableValues(items []models.Product){
	tableValues := make([][]string, 0)
	for _, item := range items {
		tableValues = append(tableValues, []string{
			item.GetFormattedDescription(),
			item.GetFormattedQuantity(),
			item.GetFormattedUnitPrice(),
			item.GetFormattedTotalPrice(),
		})
	}
	i.TableValues = tableValues
}

func (i *Invoice) RenderToPDF() ([]byte, error) {
	
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	c := canvas.NewCanvas(pdf)
	c.MoveTo(c.BorderX, c.BorderY)

	//Draw the outer border
	c.DrawRectangle(&canvas.Rectangle{
		X:           c.BorderX,
		Y:           c.BorderY,
		Width:       c.BorderWidth,
		Height:      c.BorderHeight,
		LineWidth:   0.8,
		BorderColor: canvas.PrimaryGreen,
	})
	c.MoveTo(c.MarginLeft, c.MarginTop + 10)

	//Draw the company logo on top left
	c.DrawSingleLineText(&canvas.Text{
		Content: "INVOICE",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    26,
		Color:   canvas.PrimaryGreen,
		Style:   "B",
	})
	c.IncY(10)

	c.DrawCompanyDetails()
	companyDetailsEndYPos := c.Y
	c.MoveTo(125, c.MarginTop)

	//Draw the company logo on top left
	c.DrawImageFromURL(canvas.ImageElement{
		URL:    company_details.LOGOPATH,
		X:      c.X,
		Y:      c.Y,
		Width:  70,
		Height: 0,
	})
	c.MoveTo(c.MarginLeft, companyDetailsEndYPos + 5)

	c.DrawSingleLineText(&canvas.Text{
		Content: "Bill To",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "B",
	})
	c.IncY(5)

	c.DrawCustomerDetails(i.Customer)
	c.MoveTo(140, companyDetailsEndYPos + 5)

	//Invoice No
	c.DrawLabelWithSingleLineText(&canvas.Text{
		Content: "Invoice No:",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "B",
	}, i.Number)
	c.IncY(5)

	//Invoice Date
	c.DrawLabelWithSingleLineText(&canvas.Text{
		Content: "Invoice Date:",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "B",
	}, i.CreatedAt)
	c.IncY(5)

	//Payment Due
	c.DrawLabelWithSingleLineText(&canvas.Text{
		Content: "Payment Due:",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "B",
	}, i.PaymentDue)
	c.IncY(25)
	c.ResetX()

	tableEndYPos := (&canvas.Table{
		Header: &canvas.TableHeader{
			X:           c.X,
			Y:           c.Y,
			Headers:     invoiceTableHeaders,
			CellWidths:  invoiceTableColWidths,
			TextColor:   canvas.White,
			FillColor:   canvas.PrimaryGreen,
			BorderColor: canvas.PrimaryGreen,
		},
		Body: &canvas.TableBody{
			X:           c.X,
			Y:           c.Y,
			CellWidths:  invoiceTableColWidths,
			Rows:        i.TableValues,
			TextColor:   canvas.Black,
			BorderColor: canvas.PrimaryGreen,
		},
		Width: utils.CalculateShippingTableCellWidths(shippingTableCellWidths),
	}).Draw(c, &canvas.Text{
		Font:  "Helvetica",
		Size:  10,
		Style: "B",
		Color: canvas.White,
	})
	c.MoveTo(c.MarginLeft, tableEndYPos + 5)

	c.AddNewPageIfEnd(10,canvas.PrimaryGreen, 0.8)

	c.IncX(120)
	c.DrawBillingDetails([]string{i.SubTotal, i.TaxAmount, i.Total}, i.TaxRate)
	c.MoveTo(c.MarginLeft, c.Y + 5)

	c.DrawSingleLineText(&canvas.Text{
		Content: "Terms & Conditions",
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "B",
	})
	c.IncY(5)
	c.DrawMultipleLines(&canvas.Text{
		Content: TermsAndConditions,
		Font:    "Helvetica",
		X:       c.X,
		Y:       c.Y,
		Size:    10,
		Color:   canvas.Black,
		Style:   "",
	}, 100, "")
	
	c.DrawFooter(fmt.Sprintf("If you have any questions or concerns about this invoice please contact us at %s", company_details.COMPANYEMAIL))
	
	//Generate the PDF
	bytes, err := utils.GetGeneratedPDF(c.PDF)
	return bytes, err
}