package mocks

import (
	"mime/multipart"

	"github.com/AHSChemicalsOfficialWebsite/AHSChemicalsGCShared/utils"
)

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