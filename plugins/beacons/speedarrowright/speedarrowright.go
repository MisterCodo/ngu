package speedarrowright

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type speedarrowright struct{}

func (p *speedarrowright) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: 3, Y: 2, Gain: 26.0},

		{X: 4, Y: 1, Gain: 26.0},
		{X: 3, Y: 1, Gain: 26.0},

		{X: 5, Y: 0, Gain: 26.0},
		{X: 4, Y: 0, Gain: 26.0},
		{X: 3, Y: 0, Gain: 26.0},
		{X: 2, Y: 0, Gain: 26.0},
		{X: 1, Y: 0, Gain: 26.0},

		{X: 4, Y: -1, Gain: 26.0},
		{X: 3, Y: -1, Gain: 26.0},

		{X: 3, Y: -2, Gain: 26.0},
	}
}

func (p *speedarrowright) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedarrowright) BType() beacons.BType {
	return beacons.Arrow
}

func (p *speedarrowright) Image() image.Image { return img }

func init() {
	beacons.Add(">", func() beacons.Beacon { return &speedarrowright{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/SpeedArrowRight.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
