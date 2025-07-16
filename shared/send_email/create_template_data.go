package send_email

import (
	"fmt"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/company_details"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/orders"
)

func createItemsDataForEmail(order *orders.Order) []map[string]any {
	orderItems := make([]map[string]any, 0)
	for _, item := range order.Items {
		mappedItem := make(map[string]any)
		mappedItem["sku"] = item.SKU
		mappedItem["description"] = fmt.Sprintf("%s-%s %.2f %s (Pack of %d)", item.Brand, item.Name, item.Size, item.SizeUnit, item.PackOf)
		mappedItem["quantity"] = item.Quantity
		mappedItem["price"] = fmt.Sprintf("$%.2f", item.UnitPrice * float64(item.Quantity))

		orderItems = append(orderItems, mappedItem)
	}
	return orderItems
}

func CreateAttachments (base64Contents, mimeTypes, filenames []string) []Attachment{
	attachments := make([]Attachment, 0)
	for i := range base64Contents {
		attachment := Attachment{
			Base64Content: base64Contents[i],
			MimeType:      mimeTypes[i],
			FileName:      filenames[i],
		}
		attachments = append(attachments, attachment)
	}
	return attachments
}

func CreateAdminOrderPlacedEmail(order *orders.Order, attachments []Attachment) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: company_details.EMAILINTERNALRECIPENTS,
		Data: map[string]any{
			"customer_name": order.Customer.DisplayName,
			"order_number": order.ID,
			"order_date": order.CreatedAt.Format("January 2, 2006 at 3:04 PM MST"),
			"items": createItemsDataForEmail(order),
			"subtotal": fmt.Sprintf("$%.2f", order.SubTotal),
			"tax_rate": fmt.Sprintf("%.2f%%", order.TaxRate* 100),
			"tax_amount": fmt.Sprintf("$%.2f", order.TaxAmount),
			"total": fmt.Sprintf("$%.2f", order.Total),
			"special_instructions": order.SpecialInstructions,
		},
		TemplateID: ORDER_PLACED_ADMIN_TEMPLATE_ID,
		Attachments: attachments,
	}
	return emailData
}

func CreateUserOrderPlacedEmail(order *orders.Order) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{order.Customer.PrimaryEmailAddr.Address: order.Customer.DisplayName},
		Data: map[string]any{
			"order_number": order.ID,
			"order_date": order.CreatedAt.Format("January 2, 2006 at 3:04 PM MST"),
			"items": createItemsDataForEmail(order),
			"subtotal": fmt.Sprintf("$%.2f", order.SubTotal),
			"tax_rate": fmt.Sprintf("%.2f%%", order.TaxRate* 100),
			"tax_amount": fmt.Sprintf("$%.2f", order.TaxAmount),
			"total": fmt.Sprintf("$%.2f", order.Total),
			"special_instructions": order.SpecialInstructions,
		},
		TemplateID: ORDER_PLACED_USER_TEMPLATE_ID,
		Attachments: []Attachment{},
	}
	return emailData
}
