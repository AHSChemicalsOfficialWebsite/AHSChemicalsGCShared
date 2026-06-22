package sendgrid

import (
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/company_details"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
)

// CreateOrderPlacedAdminEmailMetaData creates an email payload to notify internal admin recipients
// that a new order has been placed.
//
// Parameters:
//   - order: pointer to the Order object containing order and customer details.
//   - attachments: slice of Attachment objects representing files to be attached to the email.(Purchase order pdf)
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for admin notification.
//
// Example:
//
//	order := &orders.Order{ /* filled order details */ }
//	attachments := CreateAttachments(base64Contents, mimeTypes, filenames)
//	emailData := CreateAdminOrderPlacedEmail(order, attachments)
func CreateOrderPlacedInternalEmailMetaData(order *models.Order) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		Data: map[string]any{
			"customer_name":        order.Customer.Name,
			"order_number":         order.ID,
			"order_date":           order.CreatedAt.Format("January 2, 2006 at 3:04 PM UTC"),
			"items":                order.CreateItemsDataForAdminEmail(),
			"subtotal":             order.GetFormattedSubTotal(),
			"tax_rate":             order.GetFormattedTaxRate(),
			"tax_amount":           order.GetFormattedTaxAmount(),
			"total":                order.GetFormattedTotal(),
			"special_instructions": order.SpecialInstructions,
		},
		TemplateID: ORDER_PLACED_INTERNAL_TEMPLATE_ID,
	}
	return emailData
}

// CreateOrderPlacedUserEmailMetaData creates an email payload to notify the customer
// that their order has been successfully placed.
//
// Parameters:
//   - order: pointer to the Order object containing order and customer details.
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for customer notification.
func CreateOrderPlacedUserEmailMetaData(order *models.Order) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{order.Customer.Email: order.Customer.Name},

		Data: map[string]any{
			"order_number":         order.ID,
			"order_date":           order.CreatedAt.Format("January 2, 2006 at 3:04 PM UTC"),
			"items":                order.CreateItemsDataForUserEmail(),
			"special_instructions": order.SpecialInstructions,
		},
		TemplateID: ORDER_PLACED_USER_TEMPLATE_ID,
	}
	return emailData
}

// CreateOrderUpdatedAdminEmailMetaData creates an email payload to notify internal admin recipients
// that an order has been updated.
//
// Parameters:
//   - original: pointer to the Order object containing order and customer details.
//   - updated: pointer to the Order object containing order and customer details.
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for admin notification.
func CreateOrderUpdatedInternalEmailMetaData(order *models.Order) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		Data: map[string]any{
			"order_number":     order.ID,
			"order_updated_at": order.UpdatedAt.Format("January 2, 2006 at 3:04 PM UTC"),
			"order_status":     order.Status,
			"customer_name":    order.Customer.Name,
			"customer_email":   order.Customer.Email,
			"order_total":      order.GetFormattedTotal(),
			"items":            order.CreateItemsDataForAdminEmail(),
		},
		TemplateID: ORDER_UPDATED_INTERNAL_TEMPLATE_ID,
	}
	return emailData
}

// CreateOrderUpdatedUserEmailMetaData creates an email payload to notify the customer
// that their order has been updated.
//
// Parameters:
//   - original: pointer to the Order object containing order and customer details.
//   - updated: pointer to the Order object containing order and customer details.
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for customer notification.
func CreateOrderUpdatedUserEmailMetaData(order *models.Order) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{order.Customer.Email: order.Customer.Name},
		Data: map[string]any{
			"customer_name": order.Customer.Name,
			"order_number":  order.ID,
			"order_status":  order.Status,
			"items":         order.CreateItemsDataForUserEmail(),
		},
		TemplateID: ORDER_UPDATED_USER_TEMPLATE_ID,
	}
	return emailData
}

// CreateOrderDeliveredAdminEmailMetaData creates an email payload to notify internal admin recipients
// that an order has been delivered.
//
// Parameters:
//   - order: pointer to the Order object containing order and customer details.
//   - attachments: slice of Attachment objects representing files to be attached to the email.(Shipping manifest and temporary invoice pdf)
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for admin notification.
func CreateOrderDeliveredInternalEmailMetaData(delivery *models.Delivery) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: company_details.EmailInternalRecipients,
		Data: map[string]any{
			"order_number":  delivery.Order.ID,
			"customer_name": delivery.Order.Customer.Name,
			"items":         delivery.Order.CreateItemsDataForAdminEmail(),
			"subtotal":      delivery.Order.GetFormattedSubTotal(),
			"tax_rate":      delivery.Order.GetFormattedTaxRate(),
			"tax_amount":    delivery.Order.GetFormattedTaxAmount(),
			"total":         delivery.Order.GetFormattedTotal(),
		},
		TemplateID: ORDER_DELIVERED_INTERNAL_TEMPLATE_ID,
	}
	return emailData
}

// CreateOrderDeliveredUserEmailMetaData creates an email payload to notify the customer
// that their order has been successfully delivered.
//
// Parameters:
//   - order: pointer to the Order object containing order and customer details.
//   - delivery: pointer to the DeliveryData object containing delivery details.
//
// Returns:
//   - *EmailMetaData: a pointer to the populated EmailMetaData object configured for customer notification.
func CreateOrderDeliveredUserEmailMetaData(delivery *models.Delivery) *EmailMetaData {
	emailData := &EmailMetaData{
		Recipients: map[string]string{delivery.Order.Customer.Email: delivery.Order.Customer.Name},
		Data: map[string]any{
			"order_number": delivery.Order.ID,
			"received_by":  delivery.ReceivedBy,
			"delivered_by": delivery.DeliveredBy,
			"delivered_at": delivery.DeliveredAt.Format("January 2, 2006 at 3:04 PM UTC"),
		},
		TemplateID: ORDER_DELIVERED_USER_TEMPLATE_ID,
	}
	return emailData
}
