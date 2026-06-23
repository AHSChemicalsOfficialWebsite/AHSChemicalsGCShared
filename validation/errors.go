package validation

import "errors"

var (
	//Contact us form errors
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("invalid email address entered. Please enter the email in the correct format")
	ErrNameRequired     = errors.New("name is required")
	ErrNameInvalid      = errors.New("invalid name entered. Please enter your full name (first and last)")
	ErrPhoneRequired    = errors.New("phone number is required")
	ErrLocationRequired = errors.New("location is required")
	ErrLocationInvalid  = errors.New("invalid location entered. Location can only be between 10 and 50 characters with no symbols or special characters")
	ErrMessageRequired  = errors.New("message is required")
	ErrMessageInvalid   = errors.New("invalid message entered. Message should be between 10 and 300 characters with no symbols or special characters")

	// order errors
	ErrNoUserID                           = errors.New("no user ID found when order was placed")
	ErrNoItems                            = errors.New("no items found in order")
	ErrItemZeroQty                        = errors.New("one of the items in the order has a quantity of 0")
	ErrItemNoProduct                      = errors.New("one of the items in the order does not have any product information")
	ErrSpecialInstructionsInvalidChars    = errors.New("special instructions contain invalid characters")
	ErrSpecialInstructionsTooLong         = errors.New("special instructions are too long")
	ErrNoCustomerFound                    = errors.New("No customer found for this order")
	ErrCannotApproveNotPendingOrCancelled = errors.New("order can only be approved if it is in PENDING or CANCELLED state")
	ErrCannotApproveExpiredCancellation   = errors.New("cannot approve: more than 30 days have passed since cancellation. Place a new order instead")
	ErrCannotCancelDelivered              = errors.New("cannot cancel: order has already been delivered")
	ErrCannotCancelAlreadyCancelled       = errors.New("cannot cancel: order has already been cancelled")

	// user account creation errors
	ErrNoCustomers        = errors.New("at least one customer is required for the user")
	ErrDuplicateCustomers = errors.New("customers of the user cannot have duplicates")
	ErrNoBrands           = errors.New("brands of the user cannot be empty")
	ErrDuplicateBrands    = errors.New("brands of the user cannot have duplicates")
	ErrNoRole             = errors.New("role of the user cannot be empty")
	ErrInvalidRole        = errors.New("role of the user is not valid")

	// product errors
	ErrProductIDRequired  = errors.New("product ID is required")
	ErrBrandIDRequired    = errors.New("brand ID is required")
	ErrLocationIDRequired = errors.New("location ID is required")
	ErrNotEnoughStock     = errors.New("Not enough stock. Please update inventory")
)
