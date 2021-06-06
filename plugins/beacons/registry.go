package beacons

type Creator func() Beacon

var Beacons = map[string]Creator{}

func Add(name string, creator Creator) {
	Beacons[name] = creator
}
