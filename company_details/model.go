package company_details

type Data struct {
	Name string
	AddressLine1 string
	AddressLine2 string
	Phone string
	Email string
	WebsiteUrl string
}

func CreateNewData() *Data {
	return &Data{
		Name, 
		AddressLine1,
		AddressLine2,
		Phone, 
		Email,
		WebsiteUrl,
	}
}