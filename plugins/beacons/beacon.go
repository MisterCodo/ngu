package beacons

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
)

const (
	ImgSize = 60 // Width and height in pixels of beacon images
)

type Beacon interface {
	Effect() []Effect
	Category() Category
	BType() BType
	Image() image.Image
	Name() string
}

// Category represents the beacon category (Speed, Production or Efficiency)
type Category int

const (
	Speed = iota
	Production
	Efficiency
)

func (c Category) String() string {
	return [...]string{"Speed", "Production", "Efficiency"}[c]
}

// BType represents the beacon type (Box, Knight, Arrow, Wall, Donut)
type BType int

const (
	Box = iota
	Knight
	Arrow
	Wall
	Donut
)

func (bt BType) String() string {
	return [...]string{"Box", "Knight", "Arrow", "Wall", "Donut"}[bt]
}

// Effect indicates the impact at a location.
// X indicates moving left (X < 0) or right (X > 0)
// Y indicates moving down (Y < 0) or up (Y > 0)
// Gain indicates the gain percentage.
type Effect struct {
	X    int
	Y    int
	Gain float64
}

// ImageFromBytes reads a file and returns a decoded png image.
func ImageFromBytes(data []byte) (image.Image, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	if img.Bounds().Max.X != ImgSize || img.Bounds().Max.Y != ImgSize {
		return img, fmt.Errorf("image should be %d by %d pixels", ImgSize, ImgSize)
	}

	return img, nil
}
