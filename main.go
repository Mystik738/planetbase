package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Pallinder/go-randomdata"

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
	log.SetLevel(log.InfoLevel)

	distance := flag.Float64("d", 25.0, "The distance to space buildings.")
	templateName := flag.String("t", "template.txt", "The template filename to use.")
	outputName := flag.String("o", "save.sav", "The output filename to use.")
	help := flag.Bool("h", false, "Display the help.")
	read := flag.Bool("r", false, "Read the save instead and output some information from it.")
	planet := flag.Int64("p", 3, "The planet to use, between 0 and 3.")
	latitude := flag.Int64("la", 0, "The latitude to use. (default 0)")
	longitude := flag.Int64("lo", -128, "The longitude to use.")

	flag.Parse()

	if *help {
		helpText()
	} else if *read {
		readSave(*outputName)
	} else {
		template, xSize, zSize := readTemplate(*templateName)

		s := initSave(*planet, *latitude, *longitude)
		addTechs(&s)
		addStructures(&s, template, xSize, zSize, *distance)
		addCharacters(&s, 40, 30, 20, 5, 5)
		addBots(&s, 100, 20, 40)
		addResources(&s)
		writeSave(&s, *outputName)
	}
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
		template = append(template, make([]string, 0))

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

func initSave(planet, latitude, longitude int64) SaveGame {
	s := SaveGame{}
	ID.ID = 1

	//Need a way to input these values
	s.Planet.PlanetIndex.Value = planet
	s.Colony.Latitude.Value = latitude
	s.Colony.Longitude.Value = longitude
	s.Colony.Name.Value = randomdata.City()
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

func addStructures(s *SaveGame, template [][]string, xSize int64, zSize int64, dist float64) int64 {
	//General calculations based on coordinate system
	horizDist := math.Sin(math.Pi/3.0) * dist
	offset := dist / 2.0
	isOff := false
	xMapOffset := (maxCoord+minCoord)/2 - ((float64(xSize) - 1.0) / 2.0 * horizDist)
	zMapOffset := (maxCoord+minCoord)/2 - ((float64(zSize) - 1.0) / 2.0 * dist)

	log.Debugf("Map offsets are %v, %v", xMapOffset, zMapOffset)
	log.Infof("Template is %v by %v modules", xSize, zSize)

	totalModules := 0

	//Some matrices to store values as needed
	modGrid := make([][]*Construction, xSize)
	for i := range modGrid {
		modGrid[i] = make([]*Construction, zSize)
	}

	//X and Z here are Modules only, not template coordinates
	//Xt = Xm*2
	//Zt = Zm*2
	for x := 0; x < int(xSize); x++ {
		xPos := float64(x)*horizDist + float64(xMapOffset)
		for z := 0; z < int(zSize); z++ {
			zPos := float64(z)*dist + zMapOffset
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

			if len(template) >= x*2 && len(template[x*2]) > z*2 {
				if _, value := moduleTypes[template[x*2][z*2]]; value {
					log.Debugf("(%v,%v) %v", x, z, moduleTypes[template[x*2][z*2]])
					totalModules++
					c := initModule(template[x*2][z*2], p1)
					c.Orientation = orientModule(template, x*2, z*2)
					s.Constructions.Construction = append(s.Constructions.Construction, c)
					modGrid[x][z] = &c

					//Connections. Only connect with Modules that have been placed,
					//For each module this should be max 3 others
					if z > 0 && compareTemplate(template, x*2, z*2-1, "==") {
						ct := initConnection(*modGrid[x][z], *modGrid[x][z-1])
						s.Constructions.Construction = append(s.Constructions.Construction, ct)
					}
					if x > 0 {
						if isOff { //If we're offset, look at (x-1, z) and (x-1, z+1)
							if compareTemplate(template, (x-1)*2+1, z*2, "\\\\") {
								ct := initConnection(*modGrid[x][z], *modGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
							if z < int(zSize)-1 && compareTemplate(template, (x-1)*2+1, (z+1)*2-1, "//") {
								ct := initConnection(*modGrid[x][z], *modGrid[x-1][z+1])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
						} else { //If we're not offset, look at (x-1, z-1) and (x-1, z)
							if z > 0 && compareTemplate(template, (x-1)*2+1, (z-1)*2+1, "\\\\") {
								ct := initConnection(*modGrid[x][z], *modGrid[x-1][z-1])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
							if compareTemplate(template, (x-1)*2+1, z*2, "//") {
								ct := initConnection(*modGrid[x][z], *modGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, ct)
							}
						}
					}
				}
			}
		}
		isOff = !isOff
	}

	log.Infof("Created %v modules", totalModules)

	return int64(totalModules)
}

// Calculates the position of a connection between two Modules
func calcLinkPosition(p1 Position, p2 Position, s1 int64, s2 int64) Position {
	dist := math.Sqrt(math.Pow(p2.X-p1.X, 2.0) + math.Pow(p2.Y-p1.Y, 2.0) + math.Pow(p2.Z-p1.Z, 2.0))

	//Calc percentage from p1
	linkSize := dist - sizeToFloat[s1] - sizeToFloat[s2]
	perc := (sizeToFloat[s1] + linkSize/2) / dist

	//Get vector from p1 to p2, mult by perc, add to p1
	p := Position{
		X: p1.X + perc*(p2.X-p1.X),
		Y: p1.Y + perc*(p2.Y-p1.Y),
		Z: p1.Z + perc*(p2.Z-p1.Z),
	}
	return p
}

//Initializes a Connection
func initConnection(m1 Construction, m2 Construction) Construction {
	log.Debugf("Connecting %v and %v", m1.ID.Value, m2.ID.Value)

	p := calcLinkPosition(m1.Position, m2.Position, m1.SizeIndex.Value, m2.SizeIndex.Value)
	c := initConstruction("Connection", p)
	c.Orientation.Y = math.Atan((m2.Position.X-m1.Position.X)/(m2.Position.Z-m1.Position.Z)) * 180 / math.Pi
	c.Oxygen.Value = -1.0
	if m1.Oxygen.Value > 0.0 && m2.Oxygen.Value > 0.0 {
		c.Oxygen.Value = 1.0
	}
	l := make([]XMLInt64, 2)
	l[1].Value = m1.ID.Value
	l[0].Value = m2.ID.Value
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
	if moduleOx[t] {
		c.Oxygen.Value = 1.0
	}

	return c
}

//Orient a Module based on its connections
func orientModule(template [][]string, x int, z int) Position {
	if !compareTemplate(template, x, z-1, "  ") && !compareTemplate(template, x, z+1, "  ") {
		return Position{}
	}

	//Offset
	offset := 0
	if z%4 == 2 {
		offset = 1
	}

	if compareTemplate(template, x-1, z-1+offset, "\\\\") && compareTemplate(template, x+1, z+offset, "\\\\") {
		return Position{
			Y: 60,
		}
	}
	if compareTemplate(template, x+1, z-1+offset, "//") && compareTemplate(template, x-1, z+offset, "//") {
		return Position{
			Y: 120,
		}
	}

	return Position{}
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
	c.Oxygen.Value = -1
	c.Condition.Value = 1.0

	return c
}

//Safely compares a template index to a value.
func compareTemplate(template [][]string, x int, z int, val string) bool {
	if x < 0 || z < 0 {
		return false
	}
	if len(template) > x && len(template[x]) > z && template[x][z] == val {
		return true
	}
	return false
}

//Adds people to the SaveGame.
func addCharacters(s *SaveGame, workers int64, biologists int64, engineers int64, medics int64, guards int64) {
	if s.Characters.Character == nil {
		s.Characters.Character = make([]Character, 0)
	}

	var charMap = map[string]int64{
		"Worker":    workers,
		"Biologist": biologists,
		"Engineer":  engineers,
		"Medic":     medics,
		"Guard":     guards,
	}
	for charType, amt := range charMap {
		for i := 0; i < int(amt); i++ {
			s.Characters.Character = append(s.Characters.Character, initCharacter("Colonist", charType))
		}
	}
}

//Adds bots to the SaveGame
func addBots(s *SaveGame, carrier int64, constructor int64, driller int64) {
	if s.Characters.Character == nil {
		s.Characters.Character = make([]Character, 0)
	}

	var botMap = map[string]int64{
		"Carrier":     carrier,
		"Constructor": constructor,
		"Driller":     driller,
	}

	for charType, amt := range botMap {
		for i := 0; i < int(amt); i++ {
			s.Characters.Character = append(s.Characters.Character, initCharacter("Bot", charType))
		}
	}
}

//Initializes a character
func initCharacter(t string, s string) Character {
	log.Debugf("Creating %v of type %v", t, s)
	c := Character{
		Type: t,
	}
	p := Position{
		X: 1000,
		Y: 0,
		Z: 1000,
	}
	c.Position = p
	c.Specialization.Value = s
	if c.Specialization.Value == "Medic" {
		c.Doctor = &XMLBool{
			Value: true,
		}
	}
	c.State.Value = 3
	c.ID.Value = NextID(&ID)
	c.WanderTime.Value = 1

	vals := &XMLFloat64{
		Value: 1.0,
	}
	if t == "Colonist" {
		c.Health = vals
		c.Nutrition = vals
		c.Hydration = vals
		c.Oxygen = vals
		c.Sleep = vals
		c.Morale = vals

		//Cosmetic
		c.Gender = &XMLInt64{
			Value: rand.Int63n(2),
		}
		c.HeadIndex = &XMLInt64{
			Value: rand.Int63n(11),
		}
		c.SkinColorIndex = &XMLInt64{
			Value: rand.Int63n(11),
		}
		c.HairColorIndex = &XMLInt64{
			Value: rand.Int63n(11),
		}

		if c.Gender.Value == 0 {
			c.Name.Value = randomdata.FullName(randomdata.Male)
		} else {
			c.Name.Value = randomdata.FullName(randomdata.Female)
		}
	}

	if t == "Bot" {
		c.Name.Value = randomdata.SillyName()
		//If we want real bot names
		/*switch c.Specialization.Value {
		case "Driller":
			c.Name.Value = "DR-" + strconv.Itoa(int(NextID(&BotID)))
			break
		case "Constructor":
			c.Name.Value = "CNT-" + strconv.Itoa(int(NextID(&BotID)))
			break
		case "Carrier":
			c.Name.Value = "CR-" + strconv.Itoa(int(NextID(&BotID)))
			break
		}*/

		c.Condition = vals
		c.Integrity = vals

		//Never decay?
		c.IntegrityDecayRate = &XMLFloat64{
			Value: 100000,
		}
	}

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
		amt := 200
		if resourceName == "Bioplastic" {
			amt *= 2
		}
		for i := 0; i < amt; i++ {
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
func writeSave(s *SaveGame, o string) {
	s.IDGenerator.NextID.Value = NextID(&ID)
	s.IDGenerator.NextBotID.Value = NextID(&BotID)

	xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	checkErr(err)
	err = os.WriteFile(o, xmlBytes, 0775)
	checkErr(err)
}

//Reads the save from file and analyzes a few characteristics
func readSave(fileName string) {
	var save SaveGame
	fileContents, err := os.ReadFile(fileName)
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

	var maxGender, maxHead, maxSkin, maxHair int64

	for _, c := range save.Characters.Character {
		if c.Gender != nil && c.Gender.Value > maxGender {
			maxGender = c.Gender.Value
		}
		if c.HeadIndex != nil && c.HeadIndex.Value > maxHead {
			maxHead = c.HeadIndex.Value
		}
		if c.SkinColorIndex != nil && c.SkinColorIndex.Value > maxSkin {
			maxSkin = c.SkinColorIndex.Value
		}
		if c.HairColorIndex != nil && c.HairColorIndex.Value > maxHair {
			maxHair = c.HairColorIndex.Value
		}
	}

	sort.Strings(moduleTypeSlice)
	log.Infof("Coordinates (%v, %v)", max, min)
	log.Infof("IDs (%v, %v)", minID, maxID)
	//log.Info("\"", strings.Join(moduleTypeSlice, "\",\n\""), "\"")
	//log.Info(moduleTypeSize)
	log.Infof("Gender: %v Head: %v Skin: %v Hair: %v", maxGender, maxHead, maxSkin, maxHair)
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

func helpText() {
	fmt.Println(`This is a utility for creating Planetbase .sav files from a template. Modules in the template will be output at their largest size.

Flags:`)
	flag.PrintDefaults()
	fmt.Println(`
A small template looks like the following:

Ox==Ai
 \\//
  So

Where Ox, Ai, and So are modules and ==, //, and \\ are connections between them. Modules are expected to be connected hexagonally. The full list of module values are:`)
	keys := make([]string, 0)
	for k := range moduleTypes {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, v := range keys {
		fmt.Printf("- %v: %v\n", v, moduleTypes[v][10:])
	}
}
