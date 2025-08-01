package canvas

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"strings"
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/utils"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/httpimg"
)

type ImageElement struct {
	X, Y    float64
	Width   float64
	Height  float64
	URL     string
	Bytes   []byte
	Flow    bool
	Link    int
	LinkStr string
}

func (c *Canvas) GetCorrectByteImageDimensions(img image.Image) (float64, float64) {
	origWidth := float64(img.Bounds().Dx())
	origHeight := float64(img.Bounds().Dy())

	maxWidth, maxHeight := c.PDF.GetPageSize()

	// Scale while preserving aspect ratio
	scale := math.Min(maxWidth/origWidth, maxHeight/origHeight)
	scaledWidth := origWidth * scale
	scaledHeight := origHeight * scale

	return scaledWidth, scaledHeight
}

func (c *Canvas) DrawImageFromURL(image ImageElement) {
	if image.URL == "" {
		return
	}
	httpimg.Register(c.PDF, image.URL, "png")
	c.PDF.ImageOptions(image.URL, image.X, image.Y, image.Width, image.Height, image.Flow, gofpdf.ImageOptions{}, image.Link, image.LinkStr)
}

func (c *Canvas) DrawImageFromBytes(image ImageElement) {
	if image.Bytes == nil || len(image.Bytes) == 0 {
		return
	}
	options := gofpdf.ImageOptions{
		ImageType:             strings.ToUpper(utils.DetectImageType(image.Bytes)),
		ReadDpi:               true,
		AllowNegativePosition: false,
	}
	imageName := fmt.Sprintf("%d", time.Now().UnixNano())
	c.PDF.RegisterImageOptionsReader(imageName, options, bytes.NewReader(image.Bytes))
	c.PDF.ImageOptions(imageName, image.X, image.Y, image.Width, image.Height, image.Flow, options, image.Link, image.LinkStr)
}
