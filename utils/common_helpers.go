package utils

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"math"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
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

func MultiPartFileToBytes(file multipart.File) ([]byte, error) {
	defer file.Close()
	return io.ReadAll(file)
}

func GetLocalTimeFromTimezone(t time.Time,timezone string) time.Time {
	loc, _ := time.LoadLocation(timezone)
   	return t.In(loc)
}

// For some reason on ios, when an image is taken and sent to backend,
// it is always rotated 90 degrees when baked in the pdf. This function corrects
// that to make sure it is not rotated.
func GetCorrectlyRotatedImages(images [][]byte ) [][]byte {
	deliveryImages := make([][]byte, 0)
	for _, imageBytes := range images {
		img, _, err := image.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			continue
		}

		//Get the original orientation of the image
		orientation := 1
		x, err := exif.Decode(bytes.NewReader(imageBytes))
		if err == nil {
			tag, err := x.Get(exif.Orientation)
			if err == nil {
				o, err := tag.Int(0)
				if err == nil {
					orientation = o
				}
			}
		}

		// Correct the orientation
		switch orientation {
		case 1: // Normal
			break
		case 2: // Flipped horizontally
			img = imaging.FlipH(img)
		case 3: // Rotated 180°
			img = imaging.Rotate180(img)
		case 4: // Flipped vertically
			img = imaging.FlipV(img)
		case 5: // Transposed
			img = imaging.Transpose(img)
		case 6: // 90° clockwise
			img = imaging.Rotate270(img)
		case 7: // Transverse
			img = imaging.Transverse(img)
		case 8: // 270° clockwise (90° CCW)
			img = imaging.Rotate90(img)
		}

		// Create a new image with the correct orientation
		img = imaging.Clone(img)
		buf := new(bytes.Buffer)
		err = png.Encode(buf, img)
		if err != nil {
			continue
		}
		deliveryImages = append(deliveryImages, buf.Bytes())
	}
	// If no images were fixed in case of errors, return the original images as is
	if len(deliveryImages) == 0 {
		return images
	}
	return deliveryImages
}
