package models

type CostRule struct {
	Quality   string
	Wastetype string
	Cost      float64
}

var CostRules = []CostRule{
	{Quality: "Good", Wastetype: "Lithium Cells", Cost: 5000},
	{Quality: "Medium", Wastetype: "Lithium Cells", Cost: 3500},
	{Quality: "Bad", Wastetype: "Lithium Cells", Cost: 1000},
	{Quality: "Good", Wastetype: "Plastic Casings", Cost: 1500},
	{Quality: "Medium", Wastetype: "Plastic Casings", Cost: 700},
	{Quality: "Bad", Wastetype: "Plastic Casings", Cost: 200},
	{Quality: "Good", Wastetype: "Circuit Boards", Cost: 1500},
	{Quality: "Medium", Wastetype: "Circuit Boards", Cost: 750},
	{Quality: "Bad", Wastetype: "Circuit Boards", Cost: 150},
	{Quality: "Good", Wastetype: "Batteries", Cost: 500},
	{Quality: "Medium", Wastetype: "Batteries", Cost: 250},
	{Quality: "Bad", Wastetype: "Batteries", Cost: 50},
	{Quality: "Good", Wastetype: "Electrolyte", Cost: 1600},
	{Quality: "Medium", Wastetype: "Electrolyte", Cost: 800},
	{Quality: "Bad", Wastetype: "Electrolyte", Cost: 200},
	{Quality: "Good", Wastetype: "Rotor", Cost: 3000},
	{Quality: "Medium", Wastetype: "Rotor", Cost: 1500},
	{Quality: "Bad", Wastetype: "Rotor", Cost: 500},
	{Quality: "Good", Wastetype: "Stator", Cost: 1200},
	{Quality: "Medium", Wastetype: "Stator", Cost: 600},
	{Quality: "Bad", Wastetype: "Stator", Cost: 200},
	{Quality: "Good", Wastetype: "Worn Brushes", Cost: 300},
	{Quality: "Medium", Wastetype: "Worn Brushes", Cost: 150},
	{Quality: "Bad", Wastetype: "Worn Brushes", Cost: 30},
	{Quality: "Good", Wastetype: "Connectors", Cost: 140},
	{Quality: "Medium", Wastetype: "Connectors", Cost: 70},
	{Quality: "Bad", Wastetype: "Connectors", Cost: 10},
	{Quality: "Good", Wastetype: "Charging Cables", Cost: 600},
	{Quality: "Medium", Wastetype: "Charging Cables", Cost: 300},
	{Quality: "Bad", Wastetype: "Charging Cables", Cost: 50},
	{Quality: "Good", Wastetype: "Charging Ports", Cost: 800},
	{Quality: "Medium", Wastetype: "Charging Ports", Cost: 400},
	{Quality: "Bad", Wastetype: "Charging Ports", Cost: 100},
	{Quality: "Good", Wastetype: "Outdated Chips", Cost: 300},
	{Quality: "Medium", Wastetype: "Outdated Chips", Cost: 150},
	{Quality: "Bad", Wastetype: "Outdated Chips", Cost: 40},
	{Quality: "Good", Wastetype: "Damaged Wiring", Cost: 400},
	{Quality: "Medium", Wastetype: "Damaged Wiring", Cost: 200},
	{Quality: "Bad", Wastetype: "Damaged Wiring", Cost: 80},
	{Quality: "Good", Wastetype: "Scrap Metal", Cost: 1200},
	{Quality: "Medium", Wastetype: "Scrap Metal", Cost: 600},
	{Quality: "Bad", Wastetype: "Scrap Metal", Cost: 300},
	{Quality: "Good", Wastetype: "Frames", Cost: 5000},
	{Quality: "Medium", Wastetype: "Frames", Cost: 3500},
	{Quality: "Bad", Wastetype: "Frames", Cost: 1000},
	{Quality: "Good", Wastetype: "Plastic Trim Parts", Cost: 300},
	{Quality: "Medium", Wastetype: "Plastic Trim Parts", Cost: 150},
	{Quality: "Bad", Wastetype: "Plastic Trim Parts", Cost: 60},
	{Quality: "Good", Wastetype: "Refrigerant Gas", Cost: 700},
	{Quality: "Medium", Wastetype: "Refrigerant Gas", Cost: 350},
	{Quality: "Bad", Wastetype: "Refrigerant Gas", Cost: 100},
	{Quality: "Good", Wastetype: "Radiator Components", Cost: 1400},
	{Quality: "Medium", Wastetype: "Radiator Components", Cost: 700},
	{Quality: "Bad", Wastetype: "Radiator Components", Cost: 200},
	{Quality: "Good", Wastetype: "Worn Gearbox Components", Cost: 2000},
	{Quality: "Medium", Wastetype: "Worn Gearbox Components", Cost: 1000},
	{Quality: "Bad", Wastetype: "Worn Gearbox Components", Cost: 500},
	{Quality: "Good", Wastetype: "Clutch Plates", Cost: 800},
	{Quality: "Medium", Wastetype: "Clutch Plates", Cost: 400},
	{Quality: "Bad", Wastetype: "Clutch Plates", Cost: 100},
	{Quality: "Good", Wastetype: "Air Filters", Cost: 250},
	{Quality: "Medium", Wastetype: "Air Filters", Cost: 125},
	{Quality: "Bad", Wastetype: "Air Filters", Cost: 50},
	{Quality: "Good", Wastetype: "Compressor Unit", Cost: 3000},
	{Quality: "Medium", Wastetype: "Compressor Unit", Cost: 1500},
	{Quality: "Bad", Wastetype: "Compressor Unit", Cost: 700},
	{Quality: "Good", Wastetype: "Refrigerant Lines", Cost: 500},
	{Quality: "Medium", Wastetype: "Refrigerant Lines", Cost: 250},
	{Quality: "Bad", Wastetype: "Refrigerant Lines", Cost: 80},
	{Quality: "Good", Wastetype: "LED Lights", Cost: 600},
	{Quality: "Medium", Wastetype: "LED Lights", Cost: 300},
	{Quality: "Bad", Wastetype: "LED Lights", Cost: 50},
	{Quality: "Good", Wastetype: "Damaged Bulbs", Cost: 300},
	{Quality: "Medium", Wastetype: "Damaged Bulbs", Cost: 150},
	{Quality: "Bad", Wastetype: "Damaged Bulbs", Cost: 20},
	{Quality: "Good", Wastetype: "Tires", Cost: 4000},
	{Quality: "Medium", Wastetype: "Tires", Cost: 2000},
	{Quality: "Bad", Wastetype: "Tires", Cost: 800},
	{Quality: "Good", Wastetype: "Rims", Cost: 800},
	{Quality: "Medium", Wastetype: "Rims", Cost: 400},
	{Quality: "Bad", Wastetype: "Rims", Cost: 100},
	{Quality: "Good", Wastetype: "Wires", Cost: 400},
	{Quality: "Medium", Wastetype: "Wires", Cost: 200},
	{Quality: "Bad", Wastetype: "Wires", Cost: 90},
	{Quality: "Good", Wastetype: "Electronic Modules", Cost: 800},
	{Quality: "Medium", Wastetype: "Electronic Modules", Cost: 600},
	{Quality: "Bad", Wastetype: "Electronic Modules", Cost: 300},
	{Quality: "Good", Wastetype: "Dashboard Components", Cost: 800},
	{Quality: "Medium", Wastetype: "Dashboard Components", Cost: 400},
	{Quality: "Bad", Wastetype: "Dashboard Components", Cost: 100},
}
