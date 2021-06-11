package speedbox

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type speedbox struct{}

func (p *speedbox) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -1, Y: 1, Gain: 40.0},
		{X: 0, Y: 1, Gain: 40.0},
		{X: 1, Y: 1, Gain: 40.0},
		{X: -1, Y: 0, Gain: 40.0},
		{X: 1, Y: 0, Gain: 40.0},
		{X: -1, Y: -1, Gain: 40.0},
		{X: 0, Y: -1, Gain: 40.0},
		{X: 1, Y: -1, Gain: 40.0},
	}
}

func (p *speedbox) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedbox) BType() beacons.BType {
	return beacons.Box
}

func (p *speedbox) Image() image.Image { return img }

func init() {
	beacons.Add("*", func() beacons.Beacon { return &speedbox{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/SpeedBox.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
