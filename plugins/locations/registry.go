package locations

type Creator func() Location

var Locations = map[string]Creator{}

func Add(name string, creator Creator) {
	Locations[name] = creator
}
