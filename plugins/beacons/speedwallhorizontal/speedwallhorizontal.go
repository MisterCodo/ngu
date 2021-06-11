package speedwallhorizontal

import (
	"embed"
	"image"
	"log"

	"github.com/MisterCodo/ngu/plugins/beacons"
)

//go:embed data/*
var assets embed.FS
var img image.Image

type speedwallhorizontal struct{}

func (p *speedwallhorizontal) Effect() []beacons.Effect {
	return []beacons.Effect{
		{X: -6, Y: 0, Gain: 27.0},
		{X: -5, Y: 0, Gain: 27.0},
		{X: -4, Y: 0, Gain: 27.0},
		{X: -3, Y: 0, Gain: 27.0},
		{X: -2, Y: 0, Gain: 27.0},
		{X: -1, Y: 0, Gain: 27.0},

		{X: 1, Y: 0, Gain: 27.0},
		{X: 2, Y: 0, Gain: 27.0},
		{X: 3, Y: 0, Gain: 27.0},
		{X: 4, Y: 0, Gain: 27.0},
		{X: 5, Y: 0, Gain: 27.0},
		{X: 6, Y: 0, Gain: 27.0},
	}
}

func (p *speedwallhorizontal) Category() beacons.Category {
	return beacons.Speed
}

func (p *speedwallhorizontal) BType() beacons.BType {
	return beacons.Wall
}

func (p *speedwallhorizontal) Image() image.Image { return img }

func init() {
	beacons.Add("-", func() beacons.Beacon { return &speedwallhorizontal{} })

	var err error
	img, err = beacons.FileToImage(assets, "data/SpeedWallHorizontal.png")
	if err != nil {
		log.Fatalf("beacon image not found: %s", err.Error())
	}
}
