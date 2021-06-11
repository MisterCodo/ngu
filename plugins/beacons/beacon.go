package beacons

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/png"
	"io/fs"
)

const (
	ImgSize = 60 // Width and height in pixels of beacon images
)

type Beacon interface {
	Effect() []Effect
	Category() Category
	BType() BType
	Image() image.Image
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

// FileToImage reads a file and returns a decoded png image.
func FileToImage(assets embed.FS, filename string) (image.Image, error) {
	f, err := fs.ReadFile(assets, filename)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(f))
	if err != nil {
		return nil, err
	}

	if img.Bounds().Max.X != ImgSize || img.Bounds().Max.Y != ImgSize {
		return img, fmt.Errorf("image %s should be %d by %d pixels", filename, ImgSize, ImgSize)
	}

	return img, nil
}
