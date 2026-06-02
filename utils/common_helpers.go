package utils

import (
	"bytes"
	"image"
	"io"
	"math"
	"mime/multipart"
	"os"
	"strings"
)

func HasDuplicateStrings(slice []string) bool {
	seen := make(map[string]bool)
	for _, val := range slice {
		formattedVal := strings.ToLower(strings.TrimSpace(val))
		formattedVal = strings.ReplaceAll(formattedVal, " ", "")
		if seen[formattedVal] {
			return true 
		}
		seen[formattedVal] = true
	}
	return false
}

func AreEqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		v = strings.ToLower(strings.TrimSpace(v))
		v = strings.ReplaceAll(v, " ", "")
		b[i] = strings.ToLower(strings.TrimSpace(b[i]))
		b[i] = strings.ReplaceAll(b[i], " ", "")
		if v != b[i] {
			return false
		}
	}
	return true
}

func DetectImageType(imageBytes []byte) string {
	_, format, err := image.DecodeConfig(bytes.NewReader(imageBytes))
	if err != nil {
		return "png" 
	}
	return format
}

func GetImageBytesFromMultiPart(file multipart.File) ([]byte, error) {
	defer file.Close()
	return io.ReadAll(file)
}

func CreateMultipartFile(path string) (multipart.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func RoundToDecimals(val float64, place float64) float64 {
	return math.Round(val * 10*(place)) / (10*place)
}