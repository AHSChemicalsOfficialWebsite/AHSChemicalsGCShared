// package pdfgen is used to generate pdfs from the structures defined in layout package
package pdfgen

import (
	"encoding/base64"
	"fmt"
	"os"
)

type PDFGen interface {
	RenderToPDF() ([]byte, error) //Each layout implements this function to generate the pdf
}

func GenerateBase64PDF(pdfGen PDFGen) (string, error) {
	pdfBytes, err := pdfGen.RenderToPDF()
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(pdfBytes)
	return base64Str, nil
}

//For debugging purposes
func GeneratePDFFile(pdfGen PDFGen, fileName string) error {
	pdfBytes, err := pdfGen.RenderToPDF()
	if err != nil {
		return err
	}
	err = os.MkdirAll("./shared/pdfgen/generated", os.ModePerm)
	if err != nil {
		return err
	}

	formattedPath := fmt.Sprintf("./shared/pdfgen/generated/%s.pdf", fileName)
	err = os.WriteFile(formattedPath, pdfBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
