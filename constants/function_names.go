package constants

const (
	ContactUs                                = "contact-us"
	CreateAccount                            = "create-account"
	DeleteAccount                            = "delete-account"
	DeliverOrder                             = "deliver-order"
	GetOrderPDF                              = "get-order-pdf"
	NotificationProcessor                    = "notification-processor"
	PlaceOrder                               = "place-order"
	QuickBooksAuthBegin                      = "quickbooks-auth-begin"
	QuickBooksAuthCallback                   = "quickbooks-auth-callback"
	QuickBooksCreateInvoice                  = "quickbooks-create-invoice"
	QuickbooksRefreshTokenExpirationReminder = "quickbooks-refresh-token-expiration-reminder"
	QuickBooksSyncCustomers                  = "quickbooks-sync-customers"
	QuickBooksSyncProducts                   = "quickbooks-sync-products"
	QuickBooksWebHook                        = "quickbooks-webhook"
	QuickbooksWebHookEntityProcessor         = "quickbooks-webhook-entity-processor"
	SyncProductPricesPerCustomer             = "sync-products-prices-per-customer"
	UpdateOrder                              = "update-order"
	UpdateProductInventory                   = "update-product-inventory"
	QRScannedUnlock               			 = "qr-scanned-unlock"
)

type ServiceEndpoint struct {
	Port string
}

var Endpoints = map[string]ServiceEndpoint{
	ContactUs: {
		Port: "3002",
	},
	CreateAccount: {
		Port: "3003",
	},
	DeleteAccount: {
		Port: "3004",
	},
	DeliverOrder: {
		Port: "3005",
	},
	GetOrderPDF: {
		Port: "3006",
	},
	PlaceOrder: {
		Port: "3007",
	},
	UpdateOrder: {
		Port: "3008",
	},
	UpdateProductInventory: {
		Port: "3009",
	},
	NotificationProcessor: {
		Port: "4000",
	},
	QuickBooksAuthBegin: {
		Port: "4001",
	},
	QuickBooksAuthCallback: {
		Port: "4003",
	},
	QuickBooksCreateInvoice: {
		Port: "4004",
	},
	QuickbooksRefreshTokenExpirationReminder: {
		Port: "4005",
	},
	QuickBooksSyncCustomers: {
		Port: "4006",
	},
	QuickBooksSyncProducts: {
		Port: "4007",
	},
	QuickBooksWebHook: {
		Port: "4008",
	},
	QuickbooksWebHookEntityProcessor: {
		Port: "4009",
	},
	SyncProductPricesPerCustomer: {
		Port: "4010",
	},
	QRScannedUnlock: {
		Port: "4011",
	},
}

func GetLocalHostEndpoint(service string) string {
	e, ok := Endpoints[service]
	if !ok {
		panic("Endpoint not found: " + service)
	}
	return "http://localhost:" + e.Port + "/" + service
}
