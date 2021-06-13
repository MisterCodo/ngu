package beacons

var Beacons = map[string]Beacon{}

func Add(name string, beacon Beacon) {
	Beacons[name] = beacon
}
