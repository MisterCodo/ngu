package locations

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"strings"
)

const (
	ImgSizeX = 1200 // Width in pixels of map images
	ImgSizeY = 1020 // Width in pixels of map images
	MapX     = 20   // How many X tiles in a map
	MapY     = 17   // How many Y tiles in a map
)

type Location interface {
	Image() image.Image
	Mask() Mask
	PrettyName() string
	UglyName() string
}

type Mask [MapY][MapX]int

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

	if img.Bounds().Max.X != ImgSizeX || img.Bounds().Max.Y != ImgSizeY {
		return img, fmt.Errorf("image %s should be %d by %d pixels", filename, ImgSizeX, ImgSizeY)
	}

	return img, nil
}

// FileToMask converts a file in the data/ directory into a MapMask.
func FileToMask(fsys fs.FS, filename string) (name string, mask Mask, err error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return "", mask, fmt.Errorf("open failed: %w", err)
	}
	defer f.Close()
	hasName := false
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	y := 0
	for scanner.Scan() {
		if !hasName {
			name = scanner.Text()
			hasName = true
			continue
		}
		if scanner.Text() == "" {
			continue
		}
		for x, char := range strings.Split(scanner.Text(), "") {
			result := 0
			if char == "1" {
				result = 1
			} else if char == "2" {
				result = 2
			}
			mask[y][x] = result
		}
		y += 1
	}
	if err := scanner.Err(); err != nil {
		return "", mask, fmt.Errorf("scan failed: %w", err)
	}
	return name, mask, nil
}
