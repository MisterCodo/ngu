package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/MisterCodo/ngu/maps"
	"github.com/MisterCodo/ngu/plugins/beacons"
	"github.com/MisterCodo/ngu/plugins/locations"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/cli"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
	"github.com/maxence-charriere/go-app/v8/pkg/logs"
)

var relativePath = "/ngu/web"

type ngu struct {
	app.Compo
	locations  []location
	beacons    []beacon
	goals      []goal
	background string
	location   int
	tiles      []tile
	mask       locations.Mask
	score      float64
	saveFile   *SaveFile
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
	n.location = 0
	n.background = fmt.Sprintf("url(%s/%s.png)", relativePath, n.locations[n.location].label)
	n.beacons = []beacon{
		{id: 0, label: "box", prettyName: "Box", selected: true},
		{id: 1, label: "knight", prettyName: "Knight", selected: true},
		{id: 2, label: "arrow", prettyName: "Arrow", selected: true},
		{id: 3, label: "wall", prettyName: "Wall", selected: true},
		{id: 4, label: "donut", prettyName: "Donut", selected: true},
	}
	n.goals = []goal{
		{id: 0, label: "speed", prettyName: "Speed", selected: true},
		{id: 1, label: "production", prettyName: "Production", selected: true},
	}
	n.tiles = []tile{}
	n.mask = locations.Locations[n.locations[0].label].Mask()
	n.score = 0.0
	n.saveFile = nil
	n.updateTiles()
	n.Update()
}

func (n *ngu) Render() app.UI {
	// measure statistics
	var settingsTotalPlayedTime time.Duration
	settingsSteamAchievements := 0
	settingsMaxSteamAchievements := 10
	mapsUnlocked := 0
	mapsMaxUnlocked := 5
	mapsClearedTiles := 0
	mapsMaxClearedTiles := 432
	if n.saveFile != nil {
		settingsTotalPlayedTime = time.Duration(int(n.saveFile.Settings.TotalTimePlayed)) * time.Second

		for _, s := range n.saveFile.Settings.SteamAchievements {
			if s {
				settingsSteamAchievements++
			}
		}

		for _, m := range n.saveFile.FactoryData.Maps {
			if m.Unlocked {
				mapsUnlocked++
			}
			mapsClearedTiles += len(m.ClearedTiles)
		}
	}

	return app.Main().
		Body(
			// about
			app.Div().ID("about").
				Body(
					app.H1().Text("NGU Industries Progress Analyzer & Map Optimizer"),
					app.H2().Text("About"),
					app.P().
						Body(
							app.Text("Welcome. You have stumbled upon a "),
							app.A().Href("https://store.steampowered.com/app/1433990/NGU_INDUSTRIES/").Text("NGU Industries"),
							app.Text(` file save analyzer and map optimizer.
								You can load your save file to view some game progress statistics.
								You can also use the beacons optimizer section to optimize your beacons placement on maps, be it for
								speed, production or a combination of speed&production.`),
						),
					app.P().
						Body(
							app.Text("NOTE: Expect spoilers."),
						),
					app.P().
						Body(
							app.Text("This tool is written in GoLang and relies on "),
							app.A().Href("https://go-app.dev/").Text("Go-App"),
							app.Text(" and "),
							app.A().Href("https://github.com/kms70847/NGUI-Save-Reader").Text("NGUI-Save-Reader"),
							app.Text(". Source code can be found at "),
							app.A().Href("https://github.com/MisterCodo/ngu").Text("https://github.com/MisterCodo/ngu"),
							app.Text("."),
						),
				),
			// save file
			app.Div().ID("savefile").
				Body(
					app.Script().Src("/web/ngui_reader.bundled.js"),
					app.H2().Text("Select Save File"),
					app.Text(`This section allows you to see statistics on your NGU Industries progress. Furthermore, loading a save file
					    will automatically add blocked sections on maps to represent tiles you have yet to unlock in the game.`),
					app.Br(),
					app.Br(),
					app.Label().For("fileinput").Text("Select save file "),
					app.Input().Type("file").ID("fileinput").Name("fileinput").OnChange(n.gameFileLoad),
				),
			// statistics
			app.If(n.saveFile != nil,
				app.Div().ID("statistics").
					Body(
						app.H2().Text("Progress Details"),
						// general
						app.H3().Text("General"),
						app.Ul().
							Body(
								app.Li().Text(fmt.Sprintf("Played for %s", settingsTotalPlayedTime)),
								app.Li().Text(fmt.Sprintf("Completed %d of %d Steam achievements", settingsSteamAchievements, settingsMaxSteamAchievements)),
							),
						// maps
						app.H3().Text("Maps"),
						app.Ul().
							Body(
								app.Li().Text(fmt.Sprintf("Unlocked %d of %d maps", mapsUnlocked, mapsMaxUnlocked)),
								app.Li().Text(fmt.Sprintf("Unlocked %d of %d tiles", mapsClearedTiles, mapsMaxClearedTiles)),
							),
					),
			),
			// optimizer
			app.Div().ID("optimizer").
				Body(
					app.H2().Text("Beacons Optimizer"),
					app.Div().
						Body(
							app.Text("This section allows you to optimize the placement of beacons on NGU Industries maps."),
						),
					app.Br(),
					// location picker
					app.Div().
						Body(
							app.Text("Select the map to optimize:"),
							app.Range(n.locations).Slice(func(i int) app.UI {
								l := n.locations[i]
								return app.Div().
									Body(
										app.Input().Type("radio").ID(l.label).Name("location").Value(l.label).Checked(l.selected).OnChange(n.changeLocation(l)),
										app.Label().For(l.label).Text(l.prettyName),
									)
							}),
						),
					app.Br(),
					// beacons picker
					app.Div().
						Body(
							app.Text("Select the beacon types to use during optimization:"),
							app.Range(n.beacons).Slice(func(i int) app.UI {
								b := n.beacons[i]
								return app.Div().
									Body(
										app.Input().Type("checkbox").ID(b.label).Name(b.label).Checked(b.selected).OnChange(n.changeBeacon(b)),
										app.Label().For(b.label).Text(b.prettyName),
									)
							}),
						),
					app.Br(),
					// optimization goal picker
					app.Div().
						Body(
							app.Text(`Select the optimization goal. Generally speaking, speed is useful for labs, production is useful for
								materials with low speed caps and speed&production is useful otherwise.
								`),
							app.Range(n.goals).Slice(func(i int) app.UI {
								g := n.goals[i]
								return app.Div().
									Body(
										app.Input().Type("checkbox").ID(g.label).Name(g.label).Checked(g.selected).OnChange(n.changeGoal(g)),
										app.Label().For(g.label).Text(g.prettyName),
									)
							}),
						),
					app.Br(),
					// click on map
					app.Div().
						Body(
							app.Text("You can click on the map to block sections you have yet to unlock in the game. If you load a save file, this will be done automatically for you."),
						),
					app.Br(),
					// start optimization
					app.Div().
						Body(
							app.Text("Press Start to optimize the map. This will block the application for approximately 15 seconds before displaying results on the map below."),
							app.Br(),
							app.Input().Type("submit").Value("Start").OnClick(n.optimize),
							app.Br(),
							app.Br(),
							app.Label().Text(fmt.Sprintf("Optimized Map Score: %.2f", n.score)),
						),
					// map
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
								// usable tiles, either beacon or production
								if t.usable == 1 {
									// production tile
									if t.image == "" {
										return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").OnClick(n.clickTile(t)).
											Body()
									}
									// beacon tile
									return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").OnClick(n.clickTile(t)).
										Body(app.Img().Style("height", "30px").Style("width", "30px").Src(t.image))
								}
								// blocked tile
								if t.usable == 2 {
									return app.Button().Style("cursor", "pointer").Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Style("background-color", "transparent").OnClick(n.clickTile(t)).
										Body(app.Img().Style("height", "30px").Style("width", "30px").Src(fmt.Sprintf("%s/Unusable.png", relativePath)))
								}
								// unusable tile (e.g. water)
								return app.Div().Style("padding", "0").Style("border", "0").Style("height", "30px").Style("width", "30px").Text("")
							}),
						),
				),
		)
}

type SaveFile struct {
	Arbitrary   interface{} `json:"arbitrary"`
	Combat      interface{} `json:"combat"`
	Experiments interface{} `json:"experiments"`
	FactoryData FactoryData `json:"factoryData"`
	Farm        interface{} `json:"farm"`
	Journal     interface{} `json:"journal"`
	Materials   interface{} `json:"materials"`
	Pit         interface{} `json:"pit"`
	PlayerBase  interface{} `json:"playerBase"`
	Relics      interface{} `json:"relics"`
	Research    interface{} `json:"research"`
	Settings    Settings    `json:"settings"`
	Spin        interface{} `json:"spin"`
	Tutorial    interface{} `json:"tutorial"`
	WorkOrders  interface{} `json:"workOrders"`
}

type Settings struct {
	BeaconLock            bool    `json:"beaconLock"`
	ColourblindMode       bool    `json:"colourblindMode"`
	DamageFloatersOn      bool    `json:"damageFloatersOn"`
	FancierGainLossColour bool    `json:"fancierGainLossColour"`
	FancySpins            bool    `json:"fancyspins"`
	GainLossBuffer        int     `json:"gainLossBuffer"`
	LastTimeStamp         int     `json:"lastTimeStamp"`
	Number                int     `json:"number"`
	NumberDisplayMode     int     `json:"numberDisplayMode"`
	SteamAchievements     []bool  `json:"steamAchievements"`
	Theme                 int     `json:"theme"`
	TileOverlayID         int     `json:"tileOverlayID"`
	TimedTooltipsOn       bool    `json:"timedTooltipsOn"`
	ToolTipsOn            bool    `json:"tooltipsOn"`
	TotalTimePlayed       float64 `json:"totalTimePlayed"`
	TutorialState         int     `json:"tutorialState"`
	Version               Version `json:"version"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
	RC    int `json:"rc"`
}

type FactoryData struct {
	BlueprintsUnlocked bool        `json:"blueprintsUnlocked"`
	FavsList           interface{} `json:"favsList"`
	Maps               []SaveMap   `json:"maps"`
}

type SaveMap struct {
	Beacons      interface{} `json:"beacons"`
	Blueprints   interface{} `json:"blueprints"`
	Buildings    interface{} `json:"buildings"`
	ClearedTiles []int       `json:"clearedTiles"`
	Labs         interface{} `json:"labs"`
	Unlocked     bool        `json:"unlocked"`
}

func (n *ngu) gameFileLoad(ctx app.Context, e app.Event) {
	fmt.Printf("loading game file using NGUI Save Reader\n")

	files := ctx.JSSrc.Get("files")
	if !files.Truthy() || files.Get("length").Int() == 0 {
		fmt.Println("file not found")
		return
	}

	file := files.Index(0)
	var close func()

	onFileLoad := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		event := args[0]
		content := event.Get("target").Get("result")
		decoded := app.Window().Call("decode_ngu_save", content, true).String()
		// fmt.Println(decoded)

		// Make a JSON out of decoded
		in := []byte(decoded)

		var saveFile SaveFile
		if err := json.Unmarshal(in, &saveFile); err != nil {
			return err
		}
		n.saveFile = &saveFile

		// call updateTiles func to clear unlocked tiles according to save file
		n.updateTiles()

		// render map again
		n.score = 0.0
		n.Update()
		close()
		return nil
	})

	onFileLoadError := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		// Your error handling...
		fmt.Println("oh no, something went boom")
		close()
		return nil
	})

	// To release resources when callback are called.
	close = func() {
		onFileLoad.Release()
		onFileLoadError.Release()
	}

	reader := app.Window().Get("FileReader").New()
	reader.Set("onload", onFileLoad)
	reader.Set("onerror", onFileLoadError)
	reader.Call("readAsArrayBuffer", file)
}

func (n *ngu) changeLocation(l location) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		fmt.Printf("changed location to %s\n", l.label)
		n.background = fmt.Sprintf("url(%s/%s.png)", relativePath, l.label)
		n.location = l.id
		n.mask = locations.Locations[n.locations[n.location].label].Mask()
		n.score = 0.0
		n.updateTiles()
		n.Update()
	}
}

func (n *ngu) changeBeacon(b beacon) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		beaconValue := ctx.JSSrc.Get("checked").Bool()
		fmt.Printf("changed beacon %s to %v\n", b.label, beaconValue)
		n.beacons[b.id].selected = beaconValue
		n.score = 0.0
		n.Update()
	}
}

func (n *ngu) changeGoal(g goal) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		goalValue := ctx.JSSrc.Get("checked").Bool()
		fmt.Printf("changed goal %s to %v\n", g.label, goalValue)
		n.goals[g.id].selected = goalValue
		n.score = 0.0
		n.Update()
	}
}

func (n *ngu) clickTile(t tile) app.EventHandler {
	if t.usable == 2 {
		return func(ctx app.Context, e app.Event) {
			fmt.Printf("unblocked tile %d\n", t.id)
			n.tiles[t.id].image = ""
			n.tiles[t.id].usable = 1
			n.Update()
		}
	}
	return func(ctx app.Context, e app.Event) {
		fmt.Printf("blocked tile %d\n", t.id)
		// n.tiles[t.id].image is handled in Render given it will change based on location (eventually, still have to code this)
		n.tiles[t.id].usable = 2
		n.Update()
	}
}

func (n *ngu) updateTiles() {
	n.tiles = []tile{}
	for y, row := range n.mask {
		for x, val := range row {
			// if user has not loaded a save file and the mask says it's a block tile, overwrite it to a usable tile
			if n.saveFile == nil && val == 2 {
				val = 1
			}

			n.tiles = append(n.tiles, tile{id: y*20 + x, usable: val, image: ""})
		}
	}

	// change usable from 2 to 1 for unlocked tiles based on save file
	if n.saveFile != nil {
		fmt.Printf("save file is being used, handle unlocked tiles\n")

		// TODO: add proper error handling
		for _, clearedTile := range n.saveFile.FactoryData.Maps[n.location].ClearedTiles {
			n.tiles[clearedTile].usable = 1
		}
	}
}

func (n *ngu) optimize(ctx app.Context, e app.Event) {
	fmt.Println("Optimize!")

	var goal maps.OptimizationGoal
	if n.goals[0].selected && n.goals[1].selected {
		goal = maps.OptimizationGoal(maps.SpeedAndProductionGoal)
	} else if n.goals[0].selected {
		goal = maps.OptimizationGoal(maps.SpeedGoal)
	} else if n.goals[1].selected {
		goal = maps.OptimizationGoal(maps.ProductionGoal)
	} else {
		// nothing to do
		return
	}

	var beaconTypes []beacons.BType
	if n.beacons[0].selected {
		beaconTypes = append(beaconTypes, beacons.Box)
	}
	if n.beacons[1].selected {
		beaconTypes = append(beaconTypes, beacons.Knight)
	}
	if n.beacons[2].selected {
		beaconTypes = append(beaconTypes, beacons.Arrow)
	}
	if n.beacons[3].selected {
		beaconTypes = append(beaconTypes, beacons.Wall)
	}
	if n.beacons[4].selected {
		beaconTypes = append(beaconTypes, beacons.Donut)
	}
	if len(beaconTypes) == 0 {
		// nothing to do
		return
	}

	locationName := n.locations[n.location].label

	// blocked tiles clicked on map by user
	blockedTiles := []int{}
	for _, t := range n.tiles {
		if t.usable == 2 {
			blockedTiles = append(blockedTiles, t.id)
		}
	}

	optimizer, err := maps.NewOptimizer(goal, beaconTypes, locationName, blockedTiles)
	if err != nil {
		fmt.Printf("could not start optimization: %s", err.Error())
		return
	}

	fmt.Printf("Running %s optimization of map %s\n\n", goal.String(), locationName)

	m, err := optimizer.Run(false, 10*time.Second)
	if err != nil {
		fmt.Printf("could not run optimization: %s", err.Error())
		return
	}

	fmt.Printf("map scored %.2f\n", m.Score)
	if m.Score > n.score {
		n.score = m.Score

		for y, row := range m.Tiles {
			for x, val := range row {
				// Don't update blocked tiles by user
				if n.tiles[y*20+x].usable == 2 {
					continue
				}
				if val.Type == maps.UnusableTile || val.Type == maps.ProductionTile {
					n.tiles[y*20+x].image = ""
					continue
				}
				imgName := beacons.Beacons[val.Type].Name()
				n.tiles[y*20+x].image = fmt.Sprintf("%s/%s.png", relativePath, imgName)
			}
		}

		n.Update()
	}
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
	selected   bool
}

type goal struct {
	id         int
	label      string
	prettyName string
	selected   bool
}

type tile struct {
	id     int
	usable int // 0 means mask says not usable, 1 means mask says usable, 2 means user did not unlock tile yet
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
		Icon:        app.Icon{Default: "/web/SpeedKnight.png"},
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
