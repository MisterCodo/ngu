package locations

var Locations = map[string]Location{}

func Add(name string, location Location) {
	Locations[name] = location
}
