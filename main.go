package main

import (
	"encoding/xml"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

const minCoord float64 = 625.0
const maxCoord float64 = 1375.0

var ID AutoInc
var BotID AutoInc

func main() {
	//readSave()
	template, xSize, zSize := readTemplate("templates/template.txt")

	s := initSave()
	addTechs(&s)
	addStructures(&s, template, xSize, zSize)
	addCharacters(&s, 40, 30, 20, 5, 5)
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
		if int64(len(line))/4 > xSize {
			xSize = int64(len(line) / 4)
		}
		for x := 0; x < len(line)-1; x += 2 {
			template[i] = append(template[i], line[x:x+2])
		}
	}

	return template, zSize, xSize
}

func initSave() SaveGame {
	s := SaveGame{}
	ID.ID = 1

	//Hardcode some values
	s.Version = 12
	s.Resources.InmaterialResources.Amount.ResourceType.Value = "Coins"
	s.Camera.Height.Value = 21
	s.Camera.Position.X = 1128
	s.Camera.Position.Y = s.Camera.Height.Value
	s.Camera.Position.Z = 993
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

	return s
}

func addStructures(s *SaveGame, template [][]string, xSize int64, zSize int64) {
	vertDist := 25.0
	dist := math.Sin(math.Pi/3.0) * vertDist
	offset := vertDist / 2.0
	isOff := false

	xMapOffset := (maxCoord+minCoord)/2 - ((float64(xSize) - 1.0) / 2.0 * dist)
	zMapOffset := (maxCoord+minCoord)/2 - ((float64(zSize) - 1.0) / 2.0 * vertDist)

	//log.Printf("Map offsets are %v, %v", xMapOffset, zMapOffset)
	//log.Printf("Making %v by %v structures", xSize, zSize)

	idGrid := make([][]int64, xSize)
	for i := range idGrid {
		idGrid[i] = make([]int64, zSize)
	}

	for x := 0; x < int(xSize); x++ {
		xPos := float64(x)*dist + float64(xMapOffset)
		log.Println()
		for z := 0; z < int(zSize); z++ {
			zPos := float64(z)*vertDist + zMapOffset
			if isOff {
				zPos += offset
			}
			p := Position{
				X: xPos,
				Z: zPos,
			}

			if len(template) >= x*2 && len(template[x*2]) > z*2+1 {
				log.Printf("Placing [%v]", template[x*2][z*2+1])
				if _, value := moduleTypes[template[x*2][z*2+1]]; value {
					//log.Printf("(%v,%v) %v", x, z, moduleTypes[template[x*2][z*2+1]])
					c := initModule(template[x*2][z*2+1], p)
					s.Constructions.Construction = append(s.Constructions.Construction, c)
					idGrid[x][z] = c.ID.Value

					//Connections
					if z > 0 && compareTemplate(template, x*2, z*2, "==") {
						//log.Printf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[x*2][z*2])
						p = Position{
							X: xPos,
							Z: zPos - (vertDist / 2),
						}
						c = initConnection(p, 0, idGrid[x][z], idGrid[x][z-1])
						s.Constructions.Construction = append(s.Constructions.Construction, c)
					}
					if x > 0 {
						if isOff {
							if compareTemplate(template, (x-1)*2+1, z*2+1, "\\\\") {
								//log.Printf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][z*2+1])
								p = Position{
									X: xPos - (dist / 2),
									Z: zPos - (vertDist / 4),
								}
								c = initConnection(p, 60, idGrid[x][z], idGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, c)
							}
							if z < int(zSize)-1 && compareTemplate(template, (x-1)*2+1, (z+1)*2, "//") {
								//log.Printf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][(z+1)*2])
								p = Position{
									X: xPos - (dist / 2),
									Z: zPos + (vertDist / 4),
								}
								c = initConnection(p, 120, idGrid[x][z], idGrid[x-1][z+1])
								s.Constructions.Construction = append(s.Constructions.Construction, c)
							}
						} else {
							if compareTemplate(template, (x-1)*2+1, z*2, "\\\\") {
								//log.Printf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z-1], template[(x-1)*2+1][z*2])
								p = Position{
									X: xPos - (dist / 2),
									Z: zPos - (vertDist / 4),
								}
								c = initConnection(p, 60, idGrid[x][z], idGrid[x-1][z-1])
								s.Constructions.Construction = append(s.Constructions.Construction, c)
							}
							if z < int(zSize)-1 && compareTemplate(template, (x-1)*2+1, z*2+1, "//") {
								//log.Printf("Connecting %v and %v via %v", idGrid[x][z], idGrid[x-1][z], template[(x-1)*2+1][z*2+1])
								p = Position{
									X: xPos - (dist / 2),
									Z: zPos + (vertDist / 4),
								}
								c = initConnection(p, 120, idGrid[x][z], idGrid[x-1][z])
								s.Constructions.Construction = append(s.Constructions.Construction, c)
							}
						}
					}
				}
			}
		}
		isOff = !isOff
	}

}

func compareTemplate(template [][]string, x int, z int, val string) bool {
	if len(template) >= x && len(template[x]) > z && template[x][z] == val {
		return true
	}
	return false
}

func initConnection(p Position, rotation float64, id1 int64, id2 int64) Construction {
	c := initConstruction("Connection", p)
	c.Orientation.Y = rotation
	l := make([]XMLInt64, 2)
	l[1].Value = id1
	l[0].Value = id2
	c.Links = &Links{
		ID: l,
	}

	return c
}

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
	c.BuildProgress.Value = -1
	c.Oxygen.Value = 1.0
	c.Condition.Value = 1.0

	if c.Type == "ModuleTypePowerCollector" {
		c.PowerStorage.Value = 20000000
	}

	return c
}

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

func writeSave(s *SaveGame) {
	s.IDGenerator.NextID.Value = NextID(&ID)
	s.IDGenerator.NextBotID.Value = NextID(&BotID)

	xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	checkErr(err)
	err = os.WriteFile("saves/s.sav", xmlBytes, 0775)
	checkErr(err)
}

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
	log.Println(max)
	log.Println(min)
	log.Println(minID)
	log.Println(maxID)
	//log.Println("\"", strings.Join(moduleTypeSlice, "\",\n\""), "\"")
	log.Println(moduleTypeSize)
}

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
