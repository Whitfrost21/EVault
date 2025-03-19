package models

var LogStatus = false

type Sendreq struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Phone       string  `json:"phone"`
	Wastetype   string  `json:"wastetype"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}

var Wastetypes = []string{
	"Lithium Cells",
	"Plastic Casings",
	"Circuit Boards",
	"Battery Management System Chips",
	"Electrolyte",
	"Rotor",
	"Stator",
	"Worn Brushes",
	"Connectors",
	"Charging Cables",
	"Charging Ports",
	"Outdated Chips",
	"Damaged Wiring",
	"Scrap Metal",
	"Broken Frames",
	"Plastic Trim Parts",
	"Refrigerant Gas",
	"Radiator Components",
	"Worn Gearbox Components",
	"Clutch Plates",
	"Air Filters",
	"Compressor Unit",
	"Refrigerant Lines",
	"Broken LED Lights",
	"Damaged Bulbs",
	"Worn Tires",
	"Rims",
	"Wiring Harnesses",
	"Electronic Modules",
	"Dashboard Components",
	"Other",
}
