package main

type RelicsInfo map[int]RelicInfo

type RelicInfo struct {
	Name    string
	Impacts string
	Levels  []RelicsLevel
}

type RelicsLevel struct {
	Number int
	Cost   string
}

var relicsInfo = []RelicInfo{
	{
		Name:    "T-Rex Skull",
		Impacts: "Tier 1 Normal",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
			{Number: 11, Cost: ""},
			{Number: 12, Cost: ""},
			{Number: 13, Cost: ""},
			{Number: 14, Cost: ""},
			{Number: 15, Cost: ""},
		}},
	{
		Name:    "Used Beer Bottle",
		Impacts: "Tier 2 Normal",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
		}},
	{
		Name:    "Used Paperclip",
		Impacts: "Tier 3 Normal",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
		}},
	{
		Name:    "Hard Hat",
		Impacts: "Tier 4 Normal",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
		}},
	{
		Name:    "Wishbone",
		Impacts: "Tier 1 Flesh",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
			{Number: 11, Cost: ""},
		}},
	{
		Name:    "Dusty \"W\"",
		Impacts: "Tier 2 Flesh",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
			{Number: 11, Cost: ""},
		}},
	{
		Name:    "Grill",
		Impacts: "Tier 3 Flesh",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
		}},
	{
		Name:    "Ketchup Bottle",
		Impacts: "Tier 4 Flesh",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
		}},
	{
		Name:    "Raygun",
		Impacts: "Tier 1 Tronne",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
		}},
	{
		Name:    "Mini-UFO",
		Impacts: "Tier 2 Tronne",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
		}},
	{
		Name:    "Soldering Iron",
		Impacts: "Tier 3 Tronne",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
		}},
	{
		Name:    "Dokia Fone",
		Impacts: "Tier 4 Tronne",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
		}},
	{
		Name:    "Pure Elementium",
		Impacts: "Elementium Gain",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
		}},
	{
		Name:    "MEGA-BDSM Coin",
		Impacts: "BDSM Gain",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
		}},
	{
		Name:    "Stick (?)",
		Impacts: "Offense Power",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
			{Number: 11, Cost: ""},
			{Number: 12, Cost: ""},
			{Number: 13, Cost: ""},
			{Number: 14, Cost: ""},
		}},
	{
		Name:    "Tattered Shirt",
		Impacts: "Defense Power",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
			{Number: 9, Cost: ""},
			{Number: 10, Cost: ""},
			{Number: 11, Cost: ""},
			{Number: 12, Cost: ""},
			{Number: 13, Cost: ""},
			{Number: 14, Cost: ""},
		}},
	{
		Name:    "Smart Little Snail",
		Impacts: "Lab Power",
		Levels:  []RelicsLevel{}},
	{
		Name:    "Water Sprinkler",
		Impacts: "Plant Power",
		Levels:  []RelicsLevel{}},
	{
		Name:    "Bag of BoomBooms",
		Impacts: "Breed Power",
		Levels:  []RelicsLevel{}},
	{
		Name:    "A Chocolate Bar!",
		Impacts: "Tier 1 Candy",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
			{Number: 8, Cost: ""},
		}},
	{
		Name:    "Broken Ice Cream Machine",
		Impacts: "Tier 2 Candy",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
			{Number: 7, Cost: ""},
		}},
	{
		Name:    "The Gummiest Bear",
		Impacts: "Tier 3 Candy",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
			{Number: 6, Cost: ""},
		}},
	{
		Name:    "Golder Ticket",
		Impacts: "Tier 4 Candy",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
			{Number: 5, Cost: ""},
		}},
	{
		Name:    "Coffee Holder",
		Impacts: "Tier 1 M&M",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
			{Number: 4, Cost: ""},
		}},
	{
		Name:    "Ur'Kell's phone!",
		Impacts: "Tier 2 M&M",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
			{Number: 3, Cost: ""},
		}},
	{
		Name:    "Minku's Arm (?)",
		Impacts: "Tier 3 M&M",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
			{Number: 2, Cost: ""},
		}},
	{
		Name:    "Randawl's DM Screen",
		Impacts: "Tier 4 M&M",
		Levels: []RelicsLevel{
			{Number: 1, Cost: ""},
		}},
}

type AchievementsInfo []AchievementInfo

type AchievementInfo struct {
	Name    string
	Details string
}

var achievementsInfo = AchievementsInfo{
	{Name: "Unlocked Flesh World!", Details: ""},
	{Name: "Unlocked Planet Tronne", Details: ""},
	{Name: "Unlocked Candy Land!", Details: ""},
	{Name: "Hidden Achievement", Details: ""},
	{Name: "How Immature", Details: "Threw more than 69420 of something in the pit"},
	{Name: "Turn your first material INFINITE", Details: ""},
	{Name: "First Steps!", Details: "Combat 100 ISOPOD"},
	{Name: "Getting Higher", Details: "Combat 500 ISOPOD"},
	{Name: "Look at you go!", Details: "Combat 1000 ISOPOD"},
	{Name: "Unlocked M&M!", Details: ""},
}

type MapsInfo []MapInfo

type MapInfo struct {
	Name              string
	TilesToClearCount int
}

var mapsInfo = MapsInfo{
	{Name: "Tutorial Island", TilesToClearCount: 76},
	{Name: "Flesh World", TilesToClearCount: 113},
	{Name: "Planet Tronne", TilesToClearCount: 67},
	{Name: "Candy Land", TilesToClearCount: 100},
	{Name: "Mansions & Managers", TilesToClearCount: 73},
}
