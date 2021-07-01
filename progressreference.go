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
	{Name: "Mansions & Managers", TilesToClearCount: 76},
}

type MaterialsInfo []MaterialInfo
type BeaconsInfo []MaterialInfo

type MaterialInfo struct {
	ID   int
	Name string
}

var armyMaterials = MaterialsInfo{
	{ID: 68, Name: "Cardboard Spitdrone"},
	{ID: 33, Name: "KILLBOT V1"},
	{ID: 57, Name: "Offensive Buttbot"},
	{ID: 58, Name: "Defensive Buttbot"},
	{ID: 84, Name: "Gun Drone"},
	{ID: 85, Name: "Shield Drone"},
	{ID: 104, Name: "Tinder Surprise"},
	{ID: 112, Name: "MINTY TECH"},
	{ID: 134, Name: "Office Drone"},
	{ID: 147, Name: "Killbot V2"},
}

var tier1NormalMaterials = MaterialsInfo{
	{ID: 1, Name: "Iron Ore"},
	{ID: 2, Name: "Copper Ore"},
	{ID: 3, Name: "Iron Bar"},
	{ID: 4, Name: "Copper Bar"},
	{ID: 8, Name: "Cardboard Ore"},
	{ID: 9, Name: "Cardboard Bar"},
	{ID: 10, Name: "Glue Ore"},
	{ID: 11, Name: "Glue Bar"},
	{ID: 7, Name: "Tier 1 Think Juice"},
}

var tier2NormalMaterials = MaterialsInfo{
	{ID: 12, Name: "Crappy Iron Cog"},
	{ID: 13, Name: "Crappy Copper Wires"},
	{ID: 14, Name: "Crappy Staples"},
	{ID: 15, Name: "Crappy Computer Chip"},
	{ID: 16, Name: "Tier 2 Think Juice"},
}

var tier3NormalMaterials = MaterialsInfo{
	{ID: 19, Name: "Metal Frame"},
	{ID: 20, Name: "Crappy Wheel"},
	{ID: 21, Name: "Crappy Hard Drive"},
	{ID: 22, Name: "Shitty Robo Brain"},
	{ID: 23, Name: "Steel Ore"},
	{ID: 24, Name: "Steel Bar"},
	{ID: 25, Name: "Steel Wool"},
	{ID: 26, Name: "Crappy Reactor"},
	{ID: 27, Name: "Tier 3 Think Juice"},
}

var tier4NormalMaterials = MaterialsInfo{
	{ID: 29, Name: "Crappy Rocker Ship"},
	{ID: 30, Name: "Decent Wiring"},
	{ID: 31, Name: "Superglue"},
	{ID: 32, Name: "Basic Fuel"},
	{ID: 34, Name: "Steel Vest"},
	{ID: 35, Name: "Steel Blade"},
	{ID: 36, Name: "Basic Computer Chip"},
	{ID: 28, Name: "Tier 4 Think Juice"},
}

var tier1FleshMaterials = MaterialsInfo{
	{ID: 37, Name: "Meat Ore"},
	{ID: 38, Name: "Meat Bar"},
	{ID: 39, Name: "Bone Ore"},
	{ID: 40, Name: "Bone Bar"},
	{ID: 41, Name: "Bone Beam"},
	{ID: 42, Name: "Beating Heart"},
	{ID: 43, Name: "Tier 1 Flesh Juice"},
}

var tier2FleshMaterials = MaterialsInfo{
	{ID: 47, Name: "Bone Frame"},
	{ID: 48, Name: "Bio-reactor"},
	{ID: 49, Name: "Robo Butt"},
	{ID: 50, Name: "Bone Treads"},
	{ID: 51, Name: "Tier 2 Flesh Juice"},
}

var tier3FleshMaterials = MaterialsInfo{
	{ID: 59, Name: "Basic Shield Generator"},
	{ID: 60, Name: "GUN MK 1"},
	{ID: 61, Name: "BIO-CHIP"},
	{ID: 62, Name: "Tier 3 Flesh Juice"},
}

var tier4FleshMaterials = MaterialsInfo{
	{ID: 63, Name: "Flesh Engine"},
	{ID: 64, Name: "Flesh-Assisted Reactive Transport Solution"},
	{ID: 65, Name: "The Biocomputer"},
	{ID: 66, Name: "The Flesh Rocket"},
	{ID: 67, Name: "Tier 4 Flesh Juice"},
}

var tier1TronneMaterials = MaterialsInfo{
	{ID: 72, Name: "Technetium Ore"},
	{ID: 73, Name: "Technetium Bar"},
	{ID: 74, Name: "Plastic Ore"},
	{ID: 75, Name: "Plastic Bar"},
	{ID: 76, Name: "Tech Alloy"},
	{ID: 86, Name: "Tech Juice 1"},
}

var tier2TronneMaterials = MaterialsInfo{
	{ID: 77, Name: "Tech Frame"},
	{ID: 78, Name: "Tech CPU"},
	{ID: 79, Name: "Giant Coil"},
	{ID: 87, Name: "Tech Juice 2"},
}

var tier3TronneMaterials = MaterialsInfo{
	{ID: 80, Name: "Rail Gun"},
	{ID: 81, Name: "Defensive Cover"},
	{ID: 82, Name: "SHART Reactor"},
	{ID: 83, Name: "Tech Drone Base"},
	{ID: 88, Name: "Tech Juice 3"},
}

var tier4TronneMaterials = MaterialsInfo{
	{ID: 90, Name: "Tronne Nacelle"},
	{ID: 91, Name: "Tronne Computer"},
	{ID: 92, Name: "Cybercorn Head"},
	{ID: 93, Name: "Tronne Ship"},
	{ID: 89, Name: "Tech Juice 4"},
}

var tier1CandyMaterials = MaterialsInfo{
	{ID: 94, Name: "Candy Ore"},
	{ID: 95, Name: "Candy Bar"},
	{ID: 96, Name: "Chocolate Ore"},
	{ID: 97, Name: "Chocolate Bar"},
	{ID: 98, Name: "Candy Compound"},
	{ID: 99, Name: "Chocolate Compound"},
	{ID: 100, Name: "Candy Jello 1"},
}

var tier2CandyMaterials = MaterialsInfo{
	{ID: 101, Name: "Nitroglucoserin"},
	{ID: 102, Name: "Licorice Wire"},
	{ID: 103, Name: "Confectionary Frame"},
	{ID: 105, Name: "Nougat CPU"},
	{ID: 106, Name: "Candy Jello 2"},
}

var tier3CandyMaterials = MaterialsInfo{
	{ID: 107, Name: "D.O.N.U.T. Reactor"},
	{ID: 108, Name: "Chokets"},
	{ID: 109, Name: "Mint Mech Body"},
	{ID: 110, Name: "Cinnamon Flamethrower"},
	{ID: 111, Name: "Candy Jello 3"},
}

var tier4CandyMaterials = MaterialsInfo{
	{ID: 113, Name: "Gumball Computer"},
	{ID: 114, Name: "The Candy Ship"},
	{ID: 115, Name: "MEGA-GIGA-HYPERSPACE COMMUNICATOR PROBE"},
	{ID: 116, Name: "Candy Juice 4"}}

var tier1MMMaterials = MaterialsInfo{
	{ID: 123, Name: "Dixie Cup"},
	{ID: 124, Name: "Coffee Grounds"},
	{ID: 125, Name: "Pens"},
	{ID: 126, Name: "Printer Paper"},
	{ID: 127, Name: "Glitch"},
	{ID: 128, Name: "Clipboard"},
	{ID: 129, Name: "Coffee Cup"},
	{ID: 130, Name: "Think Dice 1"},
}

var tier2MMMaterials = MaterialsInfo{
	{ID: 131, Name: "Exploit"},
	{ID: 132, Name: "PenSword"},
	{ID: 133, Name: "Office Armour"},
	{ID: 135, Name: "Think Dice 2"}}

var tier3MMMaterials = MaterialsInfo{
	{ID: 136, Name: "Missingno"},
	{ID: 137, Name: "The Power Remote"},
	{ID: 138, Name: "The Soul Virtual Pet"},
	{ID: 139, Name: "The Time Sundial"},
	{ID: 140, Name: "The Space Pretzel"},
	{ID: 141, Name: "The Mind Abacus"},
	{ID: 142, Name: "The Reality Shroom"},
	{ID: 143, Name: "Infinity Underpants"},
	{ID: 144, Name: "Think Dice 3"},
}

var tier4MMMaterials = MaterialsInfo{
	{ID: 145, Name: "Mega CPU"},
	{ID: 146, Name: "Mega Rocket"},
	{ID: 148, Name: "SMGHCP"},
	{ID: 149, Name: "Think Dice 4"},
}

var beaconsInfo = BeaconsInfo{
	// {ID: 6, Name: "Crappy Lab"},

	{ID: 5, Name: "Box"}, // Box Speed
	// {ID: 17, Name: "Box Production"},
	// {ID: 18, Name: "Box Efficiency"},

	{ID: 44, Name: "Knight"}, // Knight Speed
	// {ID: 45, Name: "Knight Production"},
	// {ID: 46, Name: "Knight Efficiency"},

	{ID: 69, Name: "Arrow"}, // Arrow Speed
	// {ID: 70, Name: "Arrow Production"},
	// {ID: 71, Name: "Arrow Efficiency"},

	{ID: 117, Name: "Wall"}, // Wall Speed
	// {ID: 118, Name: "Wall Production"},
	// {ID: 119, Name: "Wall Efficiency"},

	{ID: 120, Name: "Donut"}, // Donut Speed
	// {ID: 121, Name: "Donut Production"},
	// {ID: 122, Name: "Donut Efficiency"},
}
