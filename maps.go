package main

type modInfo struct {
	Name string
	Ox   bool
	Min  int64
	Max  int64
}

var moduleTypes = map[string]modInfo{
	"ai": {Ox: true, Min: 0, Max: 0, Name: "ModuleTypeAirlock"},
	"an": {Ox: false, Min: 2, Max: 2, Name: "ModuleTypeAntiMeteorLaser"},
	"ba": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeBar"},
	"bs": {Ox: false, Min: 0, Max: 0, Name: "ModuleTypeBasePad"},
	"bi": {Ox: true, Min: 1, Max: 4, Name: "ModuleTypeBioDome"},
	"ca": {Ox: true, Min: 0, Max: 1, Name: "ModuleTypeCabin"},
	"cn": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeCanteen"},
	"co": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeControlCenter"},
	"do": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeDorm"},
	"fa": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeFactory"},
	"la": {Ox: true, Min: 0, Max: 1, Name: "ModuleTypeLab"},
	"ln": {Ox: false, Min: 2, Max: 2, Name: "ModuleTypeLandingPad"},
	"li": {Ox: false, Min: 0, Max: 1, Name: "ModuleTypeLightningRod"},
	"mi": {Ox: false, Min: 1, Max: 1, Name: "ModuleTypeMine"},
	"mo": {Ox: false, Min: 2, Max: 2, Name: "ModuleTypeMonolith"},
	"mu": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeMultiDome"},
	"ox": {Ox: true, Min: 0, Max: 1, Name: "ModuleTypeOxygenGenerator"},
	"po": {Ox: false, Min: 0, Max: 3, Name: "ModuleTypePowerCollector"},
	"pr": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeProcessingPlant"},
	"py": {Ox: false, Min: 4, Max: 4, Name: "ModuleTypePyramid"},
	"ra": {Ox: false, Min: 1, Max: 2, Name: "ModuleTypeRadioAntenna"},
	"ro": {Ox: true, Min: 1, Max: 2, Name: "ModuleTypeRoboticsFacility"},
	"si": {Ox: true, Min: 0, Max: 1, Name: "ModuleTypeSickBay"},
	"sg": {Ox: false, Min: 0, Max: 1, Name: "ModuleTypeSignpost"},
	"so": {Ox: false, Min: 1, Max: 4, Name: "ModuleTypeSolarPanel"},
	"st": {Ox: false, Min: 4, Max: 4, Name: "ModuleTypeStarport"},
	"sr": {Ox: true, Min: 1, Max: 4, Name: "ModuleTypeStorage"},
	"te": {Ox: false, Min: 1, Max: 2, Name: "ModuleTypeTelescope"},
	"wa": {Ox: false, Min: 1, Max: 2, Name: "ModuleTypeWaterExtractor"},
	"wt": {Ox: false, Min: 0, Max: 1, Name: "ModuleTypeWaterTank"},
	"wi": {Ox: false, Min: 0, Max: 3, Name: "ModuleTypeWindTurbine"},
}

var sizeToFloat = map[int64]float64{
	4: 8.75,
	3: 7.75,
	2: 6.25,
	1: 4.75,
	0: 3.75,
}
