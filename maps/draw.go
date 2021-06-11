package maps

import (
	"embed"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//go:embed assets/*
var assets embed.FS

const (
	imgX     = 1200
	imgY     = 1020
	tileSize = 60
)

var BaseMapImages = map[string]image.Image{}
var MappingMapImageName = map[string]string{
	"Tutorial Island":     "TutorialIsland",
	"Flesh World":         "FleshWorld",
	"Planet Tronne":       "PlanetTronne",
	"Candy Land":          "CandyLand",
	"Mansions & Managers": "MansionsAndManagers",
}

var BeaconImages = map[string]image.Image{}
var MappingBeaconSymbolImageName = map[string]string{
	"*": "SpeedBox",
	"k": "SpeedKnight",
	"^": "SpeedArrowUp",
	"v": "SpeedArrowDown",
	"<": "SpeedArrowLeft",
	">": "SpeedArrowRight",
	"-": "SpeedWallHorizontal",
	"|": "SpeedWallVertical",
	"o": "SpeedDonut",
	"b": "ProductionBox",
	"&": "ProductionKnight",
	"u": "ProductionArrowUp",
	"d": "ProductionArrowDown",
	"l": "ProductionArrowLeft",
	"r": "ProductionArrowRight",
	"h": "ProductionWallHorizontal",
	"w": "ProductionWallVertical",
	"O": "ProductionDonut",
}

// init populates image assets from the assets/ directory.
func init() {
	entries, err := fs.Glob(assets, "assets/maps/*")
	if err != nil {
		log.Fatalln("bad assets map:", err)
	}
	for _, entry := range entries {
		img, err := fileToImage(assets, entry, imgX, imgY)
		if err != nil {
			log.Fatalln("bad assets map data:", err)
		}
		base := filepath.Base(entry)
		extension := filepath.Ext(base)
		BaseMapImages[base[0:len(base)-len(extension)]] = img
	}

	entries, err = fs.Glob(assets, "assets/beacons/*")
	if err != nil {
		log.Fatalln("bad assets beacon:", err)
	}
	for _, entry := range entries {
		img, err := fileToImage(assets, entry, tileSize, tileSize)
		if err != nil {
			log.Fatalln("bad assets beacon data:", err)
		}
		base := filepath.Base(entry)
		extension := filepath.Ext(base)
		BeaconImages[base[0:len(base)-len(extension)]] = img
	}
}

// fileToImage opens a png file and decodes it.
func fileToImage(fsys fs.FS, filename string, maxX int, maxY int) (img image.Image, err error) {
	f, err := fsys.Open(filename)
	if err != nil {
		return img, fmt.Errorf("open failed: %w", err)
	}
	defer f.Close()

	img, err = png.Decode(f)
	if err != nil {
		return img, fmt.Errorf("image decode failed: %w", err)
	}

	if img.Bounds().Max.X != maxX || img.Bounds().Max.Y != maxY {
		return img, fmt.Errorf("image %s should be %d by %d pixels", filename, maxX, maxY)
	}

	return img, nil
}

// DrawMap draws the map image.
func DrawMap(m *Map, mapMaskName string, goal string, score float64) error {
	// Find base image
	mapImageName, ok := MappingMapImageName[mapMaskName]
	if !ok {
		return fmt.Errorf("could not find map mask name to map file name")
	}
	img, ok := BaseMapImages[mapImageName]
	if !ok {
		return fmt.Errorf("could not open image %s", mapMaskName)
	}

	// Initialize output image
	outputImg := image.NewRGBA(image.Rect(0, 0, imgX, imgY))
	sr := img.Bounds()
	draw.Draw(outputImg, sr, img, image.Point{}, draw.Src)

	// Go through each tile and if it's a beacon, print on top of the loaded image
	for y, row := range m.Tiles {
		for x := range row {
			if m.Tiles[y][x].Type == UnusableTile || m.Tiles[y][x].Type == ProductionTile {
				continue
			}
			beaconImageName, ok := MappingBeaconSymbolImageName[m.Tiles[y][x].Type]
			if !ok {
				return fmt.Errorf("could not find symbol %s to beacon file name", m.Tiles[y][x].Type)
			}
			beaconImg, ok := BeaconImages[beaconImageName]
			if !ok {
				return fmt.Errorf("could not find beacon %s to draw map", beaconImageName)
			}
			sr = beaconImg.Bounds()
			r := sr.Sub(sr.Min).Add(image.Point{x * tileSize, y * tileSize})
			draw.Draw(outputImg, r, beaconImg, image.Point{}, draw.Over)
		}
	}

	// Save image to disk
	outName := strings.Join([]string{mapImageName, goal, fmt.Sprintf("%.0f", score*100), fmt.Sprintf("%d", time.Now().Unix())}, "_") + ".png"
	out, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, outputImg)
	if err != nil {
		return err
	}

	fmt.Printf("Generated output image %s\n", outName)
	return nil
}
