package mocks

import (
	"mime/multipart"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/models"
	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/utils"
)

func CreateMockDeliveryInput(orderID string) *models.DeliveryFrontendInput{
	return &models.DeliveryFrontendInput{
		OrderID: orderID,
		ReceivedBy: "Harsh Mohan",
		DeliveredBy: "Harsh",
		Signature: createMockDeliverySignature(),
		Images: createMockDeliveryImages(),
	}
}

func createMockDeliverySignature() multipart.File{
	file, err := utils.CreateMultipartFile("extras/mock_shipping_manifest_image.png")
	if err != nil {
		panic(err)
	}
	return file
}

func createMockDeliveryImages() []multipart.File{
	files := make([]multipart.File, 0)
	for range 2 {
		file, err := utils.CreateMultipartFile("extras/mock_shipping_manifest_image.png")
		if err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	return files
}