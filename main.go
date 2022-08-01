package main

import (
	"encoding/xml"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

//The approximate minimum coordinate of the map
const minCoord float64 = 625.0

//The approximate maximum coordinate of the map
const maxCoord float64 = 1375.0

//Id incementer for most things
var ID AutoInc

//Id incrementer for bots
var BotID AutoInc

func main() {
	//readSave()
	template, xSize, zSize := readTemplate("templates/template.txt")

	s := initSave()
	addTechs(&s)
	addStructures(&s, template, xSize, zSize)
	addCharacters(&s, 40, 30, 20, 5, 5)
	addResources(&s)
	writeSave(&s)
}

func readTemplate(s string) ([][]string, int64, int64) {
	template := make([][]string, 0)
	var xSize, zSize int64

	//Could probably use a buffer here rather than read the whole thing
	file, err := os.ReadFile(s)
	checkErr(err)
	lines := strings.Split(string(file), "\n")
	zSize = int64((len(lines) + 1) / 2)

	for i, line := range lines {
		template = append(template, make([]string, 1))

		//Remove offsets
		if i%4 == 2 {
			line = line[2:]
		}
		if i%2 == 1 {
			line = line[1:]
		}
		//Find the longest line for later
		if int64(len(line)+3)/4 > xSize {
			xSize = int64((len(line) + 3) / 4)
		}
		//Put our structure in the template
		for x := 0; x < len(line)-1; x += 2 {
			template[i] = append(template[i], line[x:x+2])
		}
	}

	log.Debug(template)

	//z and x is swapped due to how the template is read,
	//essentially each line is depth in game, not width
	return template, zSize, xSize
}

func initSave() SaveGame {
	s := SaveGame{}
	ID.ID = 1

	//Need a way to input these values
	s.Planet.PlanetIndex.Value = 2
	s.Colony.Latitude.Value = 0
	s.Colony.Longitude.Value = -128

	s.Terrain.Seed.Value = s.Colony.Longitude.Value*1000 + s.Colony.Latitude.Value

	//Hardcode some values
	s.Version = 12
	s.Resources.InmaterialResources.Amount.ResourceType.Value = "Coins"
	s.Camera.Height.Value = 25
	s.Camera.Position.X = 1000
	s.Camera.Position.Y = s.Camera.Height.Value
	s.Camera.Position.Z = 1000
	//Maybe this could be randomized?
	s.VisitorEvents.VisitorEvent.Type = "VisitorEventFugitives"
	s.VisitorEvents.VisitorEvent.VisitorCount.Value = 6

	//Randomize meteor seeds
	meteors := make([]string, 30)
	for i := 0; i < 30; i++ {
		meteors[i] = strconv.Itoa(int(rand.Int31()))
	}
	s.MeteorManager.Seeds.Value = strings.Join(meteors, " ")

	//Default Settings
	s.ShipManager.LandingPermissions.ColonistsAllowed.Value = true
	s.ShipManager.LandingPermissions.MerchantsAllowed.Value = true
	s.ShipManager.LandingPermissions.WorkerPercentage.Value = 40
	s.ShipManager.LandingPermissions.BiologistPercentage.Value = 30
	s.ShipManager.LandingPermissions.EngineerPercentage.Value = 20
	s.ShipManager.LandingPermissions.MedicPercentage.Value = 5
	s.ShipManager.LandingPermissions.GuardPercentage.Value = 5

	//Max Manufacture Limits
	limit := int64(2147483647)
	s.ManufactureLimits.CarrierLimit.Value = limit
	s.ManufactureLimits.ConstructorLimit.Value = limit
	s.ManufactureLimits.DrillerLimit.Value = limit
	s.ManufactureLimits.GunLimit.Value = limit
	s.ManufactureLimits.MedicalSuppliesLimit.Value = limit
	s.ManufactureLimits.SemiconductorsLimit.Value = limit
	s.ManufactureLimits.SparesLimit.Value = limit

	//Push out natural disasters
	nd := 1800.0
	s.Blizzard.TimeToNextBlizzard.Value = nd
	s.SolarFlare.TimeToNextSolarFlare.Value = nd
	s.Sandstorm.TimeToNextSandstorm.Value = nd
	s.ShipManager.TimeToNextIntruder.Value = nd

	return s
}

func addStructures(s *SaveGame, template [][]string, xSize int64, zSize int64) {
	//TODO: This distance should be inputtable
	vertDist := 20.0

	//General calculations based on coordinate system
	horzDist := math.Sin(math.Pi/3.0) * vertDist
	offset := horzDist / 2.0
	isOff := false
	xMapOffset := (maxCoord+minCoord)/2 - ((float64(xSize) - 1.0) / 2.0 * horzDist)
	zMapOffset := (maxCoord+minCoord)/2 - ((float64(zSize) - 1.0) / 2.0 * vertDist)

	log.Debugf("Map offsets are %v, %v", xMapOffset, zMapOffset)
	log.Infof("Making %v by %v structures", xSize, zSize)

	//Some matrices to store values as needed
	//TODO: Make these pointers to the Modules themselves
	idGrid := make([][]int64, xSize)
	posGrid := make([][]Position, xSize)
	for i := range idGrid {
		idGrid[i] = make([]int64, zSize)
		posGrid[i] = make([]Position, zSize)
	}

	//X and Z here are Modules only, not template coordinates
	//Xt = Xm*2
	//Zt = Zm*2 + 1
	for x := 0; x < int(xSize); x++ {
		xPos := float64(x)*horzDist + float64(xMapOffset)
		for z := 0; z < int(zSize); z++ {
			zPos := float64(z)*vertDist + zMapOffset
			if isOff {
				zPos += offset
			}
			p1 := Position{
				X: xPos,
				Z: zPos,
			}

			//Offset Y on planet 3 for water
			if s.Planet.PlanetIndex.Value == 3 {
				p1.Y = 4
			}

			if len(template) >= x*2 && len(template[x*2]) > z*2+1 {
				if _, value := moduleTypes[template[x*2][z*2+1]]; value {
					log.Debugf("(%v,%v) %v", x, z, moduleTypes[template[x*2][z*2+1]])
					c := initModule(template[x*2][z*2+1], p1)
					s.Constructions.Construction = append(s.Constructions.Construction, c)
					idGrid[x][z] = c.ID.Value
					posGrid[x][z] = p1

					//Connections. Only connect with Modules that have been placed,
					//For each module this should be max 3 others
					if z > 0 && compareTemplate(template, x*2, z*2, "==") {
						p := calcLinkPosition(p1, posGrid[x][z-1], template[x*2][z*2+1], template[x*2][z*2-1], vertDist)
						ct := initConnection(p, 0, idGrid[x][z], idGrid[x][z-1])
						s.Constructions.Construction = append(s.Constructions.Construction, ct)
					}
					if x > 0 {
						//If we're offset, look at (x-1, z) and (x-1, z+1)
						if isOff {
							if compareTemplate(template, (x-1)*2+1, z*2+1, "\\\\") {
								log.Debugf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][z*2+1])
								p := calcLinkPosition(p1, posGrid[x-1][z], template[x*2][z*2+1], template[(x-1)*2][z*2+1], vertDist)
								ct := initConnection(p, 60, idGrid[x][z], idGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
							if z < int(zSize)-1 && compareTemplate(template, (x-1)*2+1, (z+1)*2, "//") {
								log.Debugf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][(z+1)*2])
								p := calcLinkPosition(p1, posGrid[x-1][z+1], template[x*2][z*2+1], template[(x-1)*2][(z+1)*2+1], vertDist)
								ct := initConnection(p, 120, idGrid[x][z], idGrid[x-1][z+1])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
						} else { //If we're not offset, look at (x-1, z) and (x-1, z-1)
							if compareTemplate(template, (x-1)*2+1, z*2, "\\\\") {
								log.Debugf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z-1], template[(x-1)*2+1][z*2])
								p := calcLinkPosition(p1, posGrid[x-1][z-1], template[x*2][z*2+1], template[(x-1)*2][(z-1)*2+1], vertDist)
								ct := initConnection(p, 60, idGrid[x][z], idGrid[x-1][z-1])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
							if z < int(zSize)-1 && compareTemplate(template, (x-1)*2+1, z*2+1, "//") {
								log.Debugf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][z*2+1])
								p := calcLinkPosition(p1, posGrid[x-1][z], template[x*2][z*2+1], template[(x-1)*2][z*2+1], vertDist)
								ct := initConnection(p, 120, idGrid[x][z], idGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
						}
					}
				}
			}
		}
		isOff = !isOff
	}

}

// Calculates the position of a connection between two Modules
func calcLinkPosition(p1 Position, p2 Position, s1 string, s2 string, dist float64) Position {
	//Calc percentage from p1
	linkSize := dist - sizeToFloat[moduleSizes[s1]] - sizeToFloat[moduleSizes[s2]]
	perc := (sizeToFloat[moduleSizes[s1]] + linkSize/2) / dist

	//Get vector from p1 to p2, mult by perc, add to p1
	p := Position{
		X: p1.X + perc*(p2.X-p1.X),
		Y: p1.Y + perc*(p2.Y-p1.Y),
		Z: p1.Z + perc*(p2.Z-p1.Z),
	}
	return p
}

//Initializes a Connection
func initConnection(p Position, rotation float64, id1 int64, id2 int64) Construction {
	c := initConstruction("Connection", p)
	c.Orientation.Y = rotation
	c.Oxygen.Value = -1.0
	l := make([]XMLInt64, 2)
	l[1].Value = id1
	l[0].Value = id2
	c.Links = &Links{
		ID: l,
	}

	return c
}

//Initializes a Module
func initModule(t string, p Position) Construction {
	c := initConstruction("Module", p)
	c.ModuleType.Value = moduleTypes[t]
	c.SizeIndex.Value = moduleSizes[t]

	if t == "Po" {
		c.PowerStorage = &XMLFloat64{
			Value: 20000000,
		}
	}

	return c
}

//Initializes a construction
func initConstruction(t string, p Position) Construction {
	c := Construction{
		Type:     t,
		Position: p,
	}
	c.BuildProgress.Value = -1
	c.Enabled.Value = true
	c.ID.Value = NextID(&ID)
	c.SizeIndex.Value = 3
	c.State.Value = 3
	c.Oxygen.Value = 1.0
	c.Condition.Value = 1.0

	return c
}

//Safely compares a template index to a value.
func compareTemplate(template [][]string, x int, z int, val string) bool {
	if len(template) > x && len(template[x]) > z && template[x][z] == val {
		return true
	}
	return false
}

//Adds people to the SaveGame.
func addCharacters(s *SaveGame, workers int64, biologists int64, engineers int64, medics int64, guards int64) {
	s.Characters.Character = make([]Character, 0)

	for i := 0; i < int(workers); i++ {
		s.Characters.Character = append(s.Characters.Character, initCharacter("Worker"))
	}
	for i := 0; i < int(biologists); i++ {
		s.Characters.Character = append(s.Characters.Character, initCharacter("Biologist"))
	}
	for i := 0; i < int(engineers); i++ {
		s.Characters.Character = append(s.Characters.Character, initCharacter("Engineer"))
	}
	for i := 0; i < int(medics); i++ {
		s.Characters.Character = append(s.Characters.Character, initCharacter("Medic"))
	}
	for i := 0; i < int(guards); i++ {
		s.Characters.Character = append(s.Characters.Character, initCharacter("Guard"))
	}
}

//Initializes a character
func initCharacter(s string) Character {
	c := Character{
		Type: "Colonist",
	}
	p := Position{
		X: 1000,
		Y: 0,
		Z: 1000,
	}
	c.Position = p
	c.Specialization.Value = s
	c.Name.Value = s
	c.State.Value = 3
	c.ID.Value = NextID(&ID)
	c.Health.Value = 1
	c.Nutrition.Value = 1
	c.Hydration.Value = 1
	c.Oxygen.Value = 1
	c.Sleep.Value = 1
	c.Morale.Value = 1
	c.WanderTime.Value = 1

	return c
}

//Adds Resources to the SaveGame
func addResources(s *SaveGame) {
	p := Position{
		X: 1000,
		Y: 0,
		Z: 1000,
	}

	for _, resourceName := range []string{"Metal", "Bioplastic", "Meal", "MedicalSupplies", "Spares", "Vegetables", "Vitromeat", "Gun", "Semiconductors"} {
		for i := 0; i < 300; i++ {
			s.Resources.Resource = append(s.Resources.Resource, initResource(resourceName, p))
		}
	}
}

//Initializes a resource
func initResource(s string, p Position) Resource {
	r := Resource{
		Type:     s,
		Position: p,
	}
	r.ID.Value = NextID(&ID)
	r.Condition.Value = 1
	r.TraderID.Value = -1

	return r
}

//Write the save to file
func writeSave(s *SaveGame) {
	s.IDGenerator.NextID.Value = NextID(&ID)
	s.IDGenerator.NextBotID.Value = NextID(&BotID)

	xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	checkErr(err)
	err = os.WriteFile("/Users/rbeckman/Documents/Planetbase/Saves/s.sav", xmlBytes, 0775)
	checkErr(err)
}

//Reads the save from file and analyzes a few characteristics
func readSave() {
	var save SaveGame
	fileContents, err := os.ReadFile("saves/save2.sav")
	checkErr(err)

	err = xml.Unmarshal(fileContents, &save)
	checkErr(err)

	var min, max Position
	var minID, maxID int64
	moduleTypeMap := make(map[string]bool)
	moduleTypeSize := make(map[string]int64)
	moduleTypeSlice := make([]string, 5)

	for i, c := range save.Constructions.Construction {

		if c.Type != "Connection" {
			modType := c.ModuleType.Value
			if _, value := moduleTypeMap[modType]; !value {
				moduleTypeMap[modType] = true
				moduleTypeSlice = append(moduleTypeSlice, modType)
				moduleTypeSize[modType] = c.SizeIndex.Value
			} else if moduleTypeSize[modType] < c.SizeIndex.Value {
				moduleTypeSize[modType] = c.SizeIndex.Value
			}
		}

		if i == 0 {
			min = c.Position
			max = c.Position
			minID = c.ID.Value
			maxID = c.ID.Value
		}

		if c.ID.Value < minID {
			minID = c.ID.Value
		}
		if c.ID.Value > maxID {
			maxID = c.ID.Value
		}

		if c.Position.X < min.X {
			min.X = c.Position.X
		}
		if c.Position.Y < min.Y {
			min.Y = c.Position.Y
		}
		if c.Position.Z < min.Z {
			min.Z = c.Position.Z
		}

		if c.Position.X > max.X {
			min.X = c.Position.X
		}
		if c.Position.Y > max.Y {
			min.Y = c.Position.Y
		}
		if c.Position.Z > max.Z {
			min.Z = c.Position.Z
		}
	}

	sort.Strings(moduleTypeSlice)
	log.Info(max)
	log.Info(min)
	log.Info(minID)
	log.Info(maxID)
	log.Info("\"", strings.Join(moduleTypeSlice, "\",\n\""), "\"")
	log.Info(moduleTypeSize)
}

//Adds all techs to the SaveGame
func addTechs(save *SaveGame) {
	techs := []string{
		"TechDrillerBot",
		"TechColossalPanel",
		"TechGmOnions",
		"TechMegaCollector",
		"TechMassiveStorage",
		"TechGoliathTurbine",
		"TechGmTomatoes",
		"TechFarmDome",
		"TechSuperExtractor",
		"TechConstructorBot",
	}

	for _, t := range techs {
		tec := XMLString{}
		tec.Value = t
		save.Techs.Tech = append(save.Techs.Tech, tec)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
