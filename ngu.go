package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/MisterCodo/ngu/maps"
	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/MisterCodo/ngu/plugins/locations"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/cli"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
	"github.com/maxence-charriere/go-app/v8/pkg/logs"
)

var relativePath = "/ngu/web/"

type ngu struct {
	app.Compo
	locations  []location
	beacons    []beacon
	background string
	location   string
	tiles      []tile
	mask       locations.Mask
	score      float64
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
	n.background = fmt.Sprintf("url(%s/%s.png)", relativePath, n.location)
	n.beacons = []beacon{
		{id: 0, label: "box", prettyName: "Box"},
		{id: 0, label: "knight", prettyName: "Knight"},
		{id: 0, label: "arrow", prettyName: "Arrow"},
		{id: 0, label: "wall", prettyName: "Wall"},
		{id: 0, label: "donut", prettyName: "Donut"},
	}
	n.tiles = []tile{}
	n.mask = locations.Locations[n.location].Mask()
	n.score = 0.0
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
									if t.image == "" {
										return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").
											Body()
									}
									return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").
										Body(app.Img().Style("height", "30px").Style("width", "30px").Src(t.image))
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
					app.Label().Text(fmt.Sprintf("Score: %.2f", n.score)),
				),
		)
}

func (n *ngu) changeLocation(l location) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Printf("changed location to %s\n", l.label)
		n.background = fmt.Sprintf("url(%s%s.png)", relativePath, l.label)
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
			n.tiles = append(n.tiles, tile{id: y*20 + x, usable: val, image: ""})
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
	goal := maps.OptimizationGoal(maps.SpeedAndProductionGoal)
	beaconTypes := []beacons.BType{beacons.Box, beacons.Knight, beacons.Arrow, beacons.Wall, beacons.Donut}
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
	n.score = m.Score

	for y, row := range m.Tiles {
		for x, val := range row {
			if val.Type == maps.UnusableTile || val.Type == maps.ProductionTile {
				n.tiles[y*20+x].image = ""
				continue
			}
			imgName := beacons.Beacons[val.Type].Name()
			n.tiles[y*20+x].image = fmt.Sprintf("%s%s.png", relativePath, imgName)
		}
	}

	n.Update()
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
	image  string
}

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	app.Route("/", &ngu{})
	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	h := app.Handler{
		Author:      "Mister Codo",
		Name:        "NGU Optimizer",
		Description: "Yet another NGU Industries map optimizer",
		Title:       "NGU Optimizer",
	}

	opts := options{Port: 4000}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&opts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run NGU Optimizer app on GitHub Pages.`).
		Options(&githubOpts)
	cli.Load()

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, opts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}
}

func runLocal(ctx context.Context, h http.Handler, opts options) {
	app.Logf("%s", logs.New("starting ngu optimizer app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	h.Resources = app.GitHubPages("ngu")
	if err := app.GenerateStaticWebsite(opts.Output, h); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
