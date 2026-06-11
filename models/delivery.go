package models

import (
	"bytes"
	"image"
	"image/png"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

type Delivery struct {
	Order          *Order
	ReceivedBy     string
	DeliveredBy    string
	Signature      []byte
	DeliveryImages [][]byte
	DeliveredAt    time.Time //In UTC
	Timezone       string
}

// For some reason on ios, when an image is taken and sent to backend,
// it is always rotated 90 degrees when baked in the pdf. This function corrects
// that to make sure it is not rotated.
func (d *Delivery) GetCorrectlyRotatedImages() [][]byte {
	deliveryImages := make([][]byte, 0)
	for _, imageBytes := range d.DeliveryImages {
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
		return d.DeliveryImages
	}
	return deliveryImages
}
