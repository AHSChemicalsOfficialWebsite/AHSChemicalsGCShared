package sendgrid

import (
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/company_details"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/quickbooks/qbmodels"
)

func CreateQuickBooksSessionExpiredEmailMetaData() *EmailMetaData{
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		TemplateID:  QUICKBOOKS_QUICKBOOKS_EXPIRED_TEMPLATE_ID,
	}
	return emailData
}

func CreateQuickBooksInvoiceAdminEmailMetaData(order *models.Order, invoice *qbmodels.Invoice) *EmailMetaData{
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		Data: map[string]any{
			"qb_invoice_no": invoice.DocNumber,
			"order_number": order.ID,
			"customer_name": order.Customer.Name,
			"invoice_date": invoice.TxnDate,
			"items": order.CreateItemsDetailedItemsDataForAdminEmail(),
			"subtotal": order.GetFormattedSubTotal(),
			"tax_percent": order.GetFormattedTaxRate(),
			"tax": order.GetFormattedTaxAmount(),
			"total_amount": order.GetFormattedTotal(),
		},
		TemplateID: FINAL_INVOICE_INTERNAL_TEMPLATE_ID,
	}
	return emailData
}