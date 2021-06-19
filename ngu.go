package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MisterCodo/ngu/maps"
	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/MisterCodo/ngu/plugins/locations"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type ngu struct {
	app.Compo
	locations  []location
	beacons    []beacon
	background string
	location   string
	tiles      []tile
	mask       locations.Mask
}

func (n *ngu) OnMount(ctx app.Context) {
	n.initNGU(ctx)
}

func (n *ngu) initNGU(ctx app.Context) {
	n.locations = []location{
		{id: 0, label: "TutorialIsland", prettyName: "Tutorial Island", selected: true},
		{id: 1, label: "FleshWorld", prettyName: "Flesh World", selected: false},
		{id: 2, label: "PlanetTronne", prettyName: "Planet Tronne", selected: false},
		{id: 3, label: "CandyLand", prettyName: "Candy Land", selected: false},
		{id: 4, label: "MansionsAndManagers", prettyName: "Mansions & Managers", selected: false},
	}
	n.location = "TutorialIsland"
	n.background = fmt.Sprintf("url(/web/%s.png)", n.location)
	n.beacons = []beacon{
		{id: 0, label: "box", prettyName: "Box"},
		{id: 0, label: "knight", prettyName: "Knight"},
		{id: 0, label: "arrow", prettyName: "Arrow"},
		{id: 0, label: "wall", prettyName: "Wall"},
		{id: 0, label: "donut", prettyName: "Donut"},
	}
	n.tiles = []tile{}
	n.mask = locations.Locations[n.location].Mask()
	n.updateTiles()
	n.Update()
}

func (n *ngu) Render() app.UI {
	return app.Main().
		Body(
			// map
			app.P().
				Body(
					app.Div().
						Style("width", fmt.Sprintf("%vpx", 600)).
						Style("height", fmt.Sprintf("%vpx", 510)).
						Style("background", n.background).
						Style("background-size", "contain").
						Style("display", "grid").
						Style("grid-template-columns", "auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto auto").
						Body(
							app.Range(n.tiles).Slice(func(i int) app.UI {
								t := n.tiles[i]
								if t.usable == 1 {
									return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").
										Body(app.Img().Style("height", "30px").Style("width", "30px").Src("web/speedbox.png"))
								}
								return app.Div().Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Text("")
							}),
						),
				),
			// location picker
			app.P().
				Body(
					app.Range(n.locations).Slice(func(i int) app.UI {
						l := n.locations[i]
						return app.Div().
							Body(
								app.Input().Type("radio").ID(l.label).Name("location").Value(l.label).Checked(l.selected).OnChange(n.changeLocation(l)),
								app.Label().For(l.label).Text(l.prettyName),
							)
					}),
				),
			// beacons picker
			app.P().
				Body(
					app.Range(n.beacons).Slice(func(i int) app.UI {
						b := n.beacons[i]
						return app.Div().
							Body(
								app.Input().Type("checkbox").ID(b.label).Name(b.label).Checked(true).OnChange(n.changeBeacon(b)),
								app.Label().For(b.label).Text(b.prettyName),
							)
					}),
				),
			// optimization picker
			// start optimization
			app.P().
				Body(
					app.Input().Type("submit").Value("Optimize!").OnClick(n.optimize),
				),
		)
}

func (n *ngu) changeLocation(l location) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Printf("changed location to %s\n", l.label)
		n.background = fmt.Sprintf("url(/web/%s.png)", l.label)
		n.location = l.label
		n.mask = locations.Locations[n.location].Mask()
		n.updateTiles()
		n.Update()
	}
}

func (n *ngu) updateTiles() {
	n.tiles = []tile{}
	for y, row := range n.mask {
		for x, val := range row {
			n.tiles = append(n.tiles, tile{id: y*20 + x, usable: val})
		}
	}
}

func (n *ngu) changeBeacon(b beacon) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		beaconValue := ctx.JSSrc.Get("checked").Bool()
		fmt.Printf("changed beacon %s to %v\n", b.label, beaconValue)
		n.Update()
	}
}

func (n *ngu) optimize(ctx app.Context, e app.Event) {
	fmt.Println("Optimize!")

	// hardcoded for debugging, use ngu values instead
	goal := maps.OptimizationGoal(maps.SpeedGoal)
	beaconTypes := []beacons.BType{beacons.Box}
	locationName := n.location

	optimizer, err := maps.NewOptimizer(goal, beaconTypes, locationName)
	if err != nil {
		fmt.Printf("could not start optimization: %s", err.Error())
		return
	}
	optimizer.Infinite = false

	fmt.Printf("Running %s optimization of map %s\n\n", goal.String(), locationName)

	m, err := optimizer.Run(false)
	if err != nil {
		fmt.Printf("could not run optimization: %s", err.Error())
		return
	}

	fmt.Printf("map scored %.2f\n", m.Score)
}

type location struct {
	id         int
	label      string
	prettyName string
	selected   bool
}

type beacon struct {
	id         int
	label      string
	prettyName string
}

type tile struct {
	id     int
	usable int
}

func main() {
	app.Route("/", &ngu{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "A NGU Map Optimizer",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
