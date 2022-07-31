package main

import (
	"encoding/xml"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	readSave()

	s := initSave()
	addTechs(&s)
	addStructures(&s)
	writeSave(&s)
}

const minCoord float64 = 625.0
const maxCoord float64 = 1375.0

var ID AutoInc
var BotID AutoInc

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

	return s
}

func addStructures(s *SaveGame) {
	vertDist := 25.0
	dist := math.Sin(math.Pi/3.0) * vertDist
	offset := vertDist / 2.0
	isOff := false

	xSize := int64((maxCoord - minCoord) / dist)
	zSize := int64((maxCoord - minCoord) / vertDist)

	log.Printf("Making %v by %v structures", xSize, zSize)

	idGrid := make([][]int64, xSize)
	for i := range idGrid {
		idGrid[i] = make([]int64, zSize)
	}

	for x := 0; x < int(xSize); x++ {
		xPos := minCoord + float64(x)*dist
		for z := 0; z < int(zSize); z++ {
			zPos := minCoord + float64(z)*vertDist
			if isOff {
				zPos += offset
			}
			p := Position{
				X: xPos,
				Z: zPos,
			}

			c := initConstruction("Module", p)
			c.ModuleType.Value = "ModuleTypeWindTurbine"
			if isOff {
				c.ModuleType.Value = "ModuleTypePowerCollector"
			}
			s.Constructions.Construction = append(s.Constructions.Construction, c)

			idGrid[x][z] = c.ID.Value

			//Connections
			if z > 0 {
				p = Position{
					X: xPos,
					Z: zPos - (vertDist / 2),
				}
				c = initConnection(p, 0, idGrid[x][z], idGrid[x][z-1])
				s.Constructions.Construction = append(s.Constructions.Construction, c)
			}
			if x > 0 {
				if isOff {
					p = Position{
						X: xPos - (dist / 2),
						Z: zPos - (vertDist / 4),
					}
					c = initConnection(p, 60, idGrid[x][z], idGrid[x-1][z])
					s.Constructions.Construction = append(s.Constructions.Construction, c)
					if z < int(zSize)-1 {
						p = Position{
							X: xPos - (dist / 2),
							Z: zPos + (vertDist / 4),
						}
						c = initConnection(p, 120, idGrid[x][z], idGrid[x-1][z+1])
						s.Constructions.Construction = append(s.Constructions.Construction, c)
					}
				} else {
					if z > 0 {
						p = Position{
							X: xPos - (dist / 2),
							Z: zPos - (vertDist / 4),
						}
						c = initConnection(p, 60, idGrid[x][z], idGrid[x-1][z-1])
						s.Constructions.Construction = append(s.Constructions.Construction, c)
					}
					p = Position{
						X: xPos - (dist / 2),
						Z: zPos + (vertDist / 4),
					}
					c = initConnection(p, 120, idGrid[x][z], idGrid[x-1][z])
					s.Constructions.Construction = append(s.Constructions.Construction, c)
				}
			}
		}
		isOff = !isOff
	}

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
	c.Oxygen.Value = -1
	c.Condition.Value = 1.0

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

	for i, c := range save.Constructions.Construction {
		//fmt.Println(c.Type)
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

	log.Println(max)
	log.Println(min)
	log.Println(minID)
	log.Println(maxID)
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
		tec := Tech{}
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
