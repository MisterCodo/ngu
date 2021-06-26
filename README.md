# NGU Optimizer

## How to Use

Go to [NGU Optimizer](https://mistercodo.github.io/ngu/) or download the latest release and use command `cli -h` to see usage instructions:

### Optimize Map

Optimize a map with command `optimize`.

```ascii
Optimize placement of beacons on NGU Industries map.

Usage:
  cli optimize [flags]

Flags:
  -b, --beacons int   optimization beacon types available: (1)Box, (2)Box & Knight, (3)Box, Knight & Arrow, (4) Box, Knight, Arrow & Wall, (5)All (default 5)
  -g, --goal int      optimization goal: (1)Speed&Production, (2)Speed, (3)Production (default 1)
  -h, --help          help for optimize
  -m, --map int       map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers (default 1)
```

For instance, to launch a speed optimization for Candy Land with box and knight beacons, use command `cli optimize -m 4 -b 2 -g 2`. It will run forever and everytime a new higher scoring map is found it will be both printed to console and saved as a png image.

If you want to make modifications to the code, the only dependency is [GoLang](https://golang.org/). Once GoLang is installed, clone the repo and from its root use command `go run github.com/MisterCodo/ngu/cli`.

### Draw Map

If you already have an optimized map layout, you can draw an image of it by using command `draw`.

```ascii
Draw map to disk for choosen location according to beacons file provided.

Usage:
  cli draw [flags]

Flags:
  -f, --file string   file consisting of characters defining placement of beacons. These beacons will be drawn on top of selected map.        
  -h, --help          help for draw
  -m, --map int       map to optimize: (1)Tutorial Island, (2)Flesh World, (3)Planet Tronne, (4)Candy Land, (5)Mansions & Managers (default 1)
```

The file should contain the console map layout results, for instance running `cli draw -m 1 -f beacons.txt` draws the following image:

![Optimized Map Image](/_images/TutorialIslandSample.png)

Where `beacons.txt` was:

```ascii
        ww
        www    .
       v dv   *.*O<
       vvvvd O*.*O<
        dd  >O*.*O<
   hr>> ..   < . O<
   hr>r.... <<hw.O<.
   hr>r.... l  w
   h>>>...OO<  w
   .>>O^uu^O<  w
      ^^ ^u^>.*.*O<
             O*.  <
             O*.*O<
  .    www   O*.*O<
  .h   .*.
       .*.
```

## Design

The design relies on 3 steps. The first step consists of generating randomized maps and only keeping the highest scoring randomized map. The second step consists on generating a good map (local optimization) by making slight adjustments to the randomized map, both randomly and with beam search. The third step consists of generating a bunch of good candidate maps and keeping the best candidate (global optimization), followed by one last beam search on this best candidate.

The resulting map is presented to the console where tile types are:

```ascii
  = unusable (space character)
. = resource

* = box (speed)
k = knight (speed)
^v<> = up, down, left, right arrows (speed)
-| = horizontal, vertical wall (speed)
o = donut (speed)

b = box (production)
& = knight (production)
udlr = up, down, left, right arrows (production)
hw = horizontal, vertical wall (production)
O = donut (production)
```

## Custom Map

If you have not yet unlocked all the tiles of a given map, you can go into subfolder `plugins/locations/<mapname>/data` and edit the corresponding txt file before running the optimizer. This is currently only available by running the optimizer from code with Go.

## TODO / Upcoming Features

- Improve performance
- Reactivate beam after fixing it
- Add efficiency beacons
- Add heat maps of both speed and production gains
- Clean code
- Multithread UI
- Add stats based on loaded save file
- Change locked tile image according to map

## Best Maps

Best generated maps from this tool, and other tools, can be found at [NGU Industries Fandom](https://ngu-industries.fandom.com/wiki/Optimal_beacons_configurations).

## License

This software is licensed under MIT.

Save file reader is using [https://github.com/kms70847/NGUI-Save-Reader](https://github.com/kms70847/NGUI-Save-Reader).
