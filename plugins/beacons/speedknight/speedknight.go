package speedknight

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type speedknight struct{}

func (p *speedknight) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -2, Y: 1, Gain: 35.0},
		{X: -1, Y: 2, Gain: 35.0},

		{X: 2, Y: 1, Gain: 35.0},
		{X: 1, Y: 2, Gain: 35.0},

		{X: -2, Y: -1, Gain: 35.0},
		{X: -1, Y: -2, Gain: 35.0},

		{X: 2, Y: -1, Gain: 35.0},
		{X: 1, Y: -2, Gain: 35.0},
	}
}

func (p *speedknight) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedknight) BType() beacons.BType {
	return beacons.Knight
}

func (p *speedknight) Image() image.Image { return img }

func init() {
	beacons.Add("k", func() beacons.Beacon { return &speedknight{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/SpeedKnight.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
