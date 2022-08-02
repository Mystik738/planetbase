package main

import "encoding/xml"

type AutoInc struct {
	ID int64
}

func NextID(ai *AutoInc) int64 {
	ai.ID = ai.ID + 1
	return ai.ID
}

type XMLInt64 struct {
	Text  string `xml:",chardata"`
	Value int64  `xml:"value,attr"`
}

type XMLFloat64 struct {
	Text  string  `xml:",chardata"`
	Value float64 `xml:"value,attr"`
}

type XMLString struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}

type XMLBool struct {
	Text  string `xml:",chardata"`
	Value bool   `xml:"value,attr"`
}

type Position struct {
	Text string  `xml:",chardata"`
	X    float64 `xml:"x,attr"`
	Y    float64 `xml:"y,attr"`
	Z    float64 `xml:"z,attr"`
}

type VisitorEvents struct {
	Text          string       `xml:",chardata"`
	NextEventTime XMLFloat64   `xml:"next-event-time"`
	VisitorEvent  VisitorEvent `xml:"visitor-event"`
}

type VisitorEvent struct {
	Text         string   `xml:",chardata"`
	Type         string   `xml:"type,attr"`
	VisitorCount XMLInt64 `xml:"visitor-count"`
}

type Construction struct {
	Text               string           `xml:",chardata"`
	Type               string           `xml:"type,attr"`
	Enabled            XMLBool          `xml:"enabled"`
	State              XMLInt64         `xml:"state"`
	BuildProgress      XMLFloat64       `xml:"build-progress"`
	Condition          XMLFloat64       `xml:"condition"`
	Oxygen             XMLFloat64       `xml:"oxygen"`
	ID                 XMLInt64         `xml:"id"`
	Position           Position         `xml:"position"`
	Orientation        Position         `xml:"orientation"`
	TimeBuilt          XMLFloat64       `xml:"time-built"`
	Locked             XMLBool          `xml:"locked"`
	HighPriority       XMLBool          `xml:"high-priority"`
	ModuleType         XMLString        `xml:"module-type"`
	SizeIndex          XMLInt64         `xml:"size-index"`
	MobileRotation     Position         `xml:"mobile-rotation"`
	PowerStorage       *XMLFloat64      `xml:"power-storage"`
	Components         *Components      `xml:"components"`
	ProductionProgress *XMLFloat64      `xml:"production-progress"`
	ResourceStorage    *ResourceStorage `xml:"resource-storage"`
	LaserCharge        *XMLFloat64      `xml:"laser-charge"`
	Links              *Links           `xml:"links"`
}

type Components struct {
	Text      string `xml:",chardata"`
	Component []struct {
		Text    string  `xml:",chardata"`
		Type    string  `xml:"type,attr"`
		Enabled XMLBool `xml:"enabled"`
		State   struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"state"`
		BuildProgress XMLFloat64 `xml:"build-progress"`
		ID            XMLInt64   `xml:"id"`
		ComponentType struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"component-type"`
		Position           Position   `xml:"position"`
		Orientation        Position   `xml:"orientation"`
		Condition          XMLFloat64 `xml:"condition"`
		ProductionProgress XMLFloat64 `xml:"production-progress"`
		Time               XMLFloat64 `xml:"time"`
		ProducedItemIndex  struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"produced-item-index"`
		ResourceContainer struct {
			Text     string `xml:",chardata"`
			Capacity struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"capacity"`
			Resource []Resource `xml:"resource"`
		} `xml:"resource-container"`
	} `xml:"component"`
}

type Slot struct {
	Text      string     `xml:",chardata"`
	Position  Position   `xml:"position"`
	MaxHeight XMLFloat64 `xml:"max-height"`
	Resource  []Resource `xml:"resource"`
}

type ResourceStorage struct {
	Text string `xml:",chardata"`
	Slot []Slot `xml:"slot"`
}

type Links struct {
	Text string     `xml:",chardata"`
	ID   []XMLInt64 `xml:"id"`
}

type Resources struct {
	Text                string              `xml:",chardata"`
	Resource            []Resource          `xml:"resource"`
	InmaterialResources InmaterialResources `xml:"inmaterial-resources"`
}

type Resource struct {
	Text        string     `xml:",chardata"`
	Type        string     `xml:"type,attr"`
	ID          XMLInt64   `xml:"id"`
	TraderID    XMLInt64   `xml:"trader-id"`
	Position    Position   `xml:"position"`
	Orientation Position   `xml:"orientation"`
	State       XMLInt64   `xml:"state"`
	Location    XMLInt64   `xml:"location"`
	Subtype     XMLInt64   `xml:"subtype"`
	Condition   XMLFloat64 `xml:"condition"`
	Durability  XMLFloat64 `xml:"durability"`
}

type InmaterialResources struct {
	Text          string    `xml:",chardata"`
	ContainerName XMLString `xml:"container-name"`
	Amount        Amount    `xml:"amount"`
}

type Amount struct {
	Text         string    `xml:",chardata"`
	ResourceType XMLString `xml:"resource-type"`
	Amount       XMLInt64  `xml:"amount"`
}

type Character struct {
	Text                    string      `xml:",chardata"`
	Type                    string      `xml:"type,attr"`
	Position                Position    `xml:"position"`
	Orientation             Position    `xml:"orientation"`
	Location                XMLInt64    `xml:"location"`
	Name                    XMLString   `xml:"name"`
	Specialization          XMLString   `xml:"specialization"`
	StatusFlags             XMLInt64    `xml:"status-flags"`
	State                   XMLInt64    `xml:"state"`
	ID                      XMLInt64    `xml:"id"`
	WanderTime              XMLFloat64  `xml:"wander-time"`
	Health                  *XMLFloat64 `xml:"Health"`
	Nutrition               *XMLFloat64 `xml:"Nutrition"`
	Hydration               *XMLFloat64 `xml:"Hydration"`
	Oxygen                  *XMLFloat64 `xml:"Oxygen"`
	Sleep                   *XMLFloat64 `xml:"Sleep"`
	Morale                  *XMLFloat64 `xml:"Morale"`
	Gender                  *XMLInt64   `xml:"gender"`
	BasicMealCount          *XMLInt64   `xml:"basic-meal-count"`
	HeadIndex               *XMLInt64   `xml:"head-index"`
	SkinColorIndex          *XMLInt64   `xml:"skin-color-index"`
	HairColorIndex          *XMLInt64   `xml:"hair-color-index"`
	Doctor                  *XMLBool    `xml:"doctor"`
	InmunityToContagionTime *XMLFloat64 `xml:"inmunity-to-contagion-time"`
	LoadedResource          *XMLInt64   `xml:"loaded-resource"`
	Condition               *XMLFloat64 `xml:"Condition"`
	Integrity               *XMLFloat64 `xml:"Integrity"`
	IntegrityDecayRate      *XMLFloat64 `xml:"integrity-decay-rate"`
}

type SaveGame struct {
	XMLName     xml.Name `xml:"save-game"`
	Text        string   `xml:",chardata"`
	Version     int64    `xml:"version,attr"`
	IDGenerator struct {
		Text      string   `xml:",chardata"`
		NextID    XMLInt64 `xml:"next-id"`
		NextBotID XMLInt64 `xml:"next-bot-id"`
	} `xml:"id-generator"`
	Planet struct {
		Text        string   `xml:",chardata"`
		PlanetIndex XMLInt64 `xml:"planet-index"`
	} `xml:"planet"`
	Milestones struct {
		Text      string      `xml:",chardata"`
		Milestone []XMLString `xml:"milestone"`
	} `xml:"milestones"`
	Techs struct {
		Text string      `xml:",chardata"`
		Tech []XMLString `xml:"tech"`
	} `xml:"techs"`
	Environment struct {
		Text          string     `xml:",chardata"`
		TimeOfDay     XMLFloat64 `xml:"time-of-day"`
		WindIndicator XMLFloat64 `xml:"wind-indicator"`
	} `xml:"environment"`
	Terrain struct {
		Text string   `xml:",chardata"`
		Seed XMLInt64 `xml:"seed"`
	} `xml:"terrain"`
	Camera struct {
		Text        string     `xml:",chardata"`
		Height      XMLFloat64 `xml:"height"`
		Position    Position   `xml:"position"`
		Orientation Position   `xml:"orientation"`
	} `xml:"camera"`
	Sandstorm struct {
		Text                string     `xml:",chardata"`
		SandstormInProgress XMLBool    `xml:"sandstorm-in-progress"`
		TimeToNextSandstorm XMLFloat64 `xml:"time-to-next-sandstorm"`
		Time                XMLFloat64 `xml:"time"`
		SandstormTime       XMLFloat64 `xml:"sandstorm-time"`
	} `xml:"sandstorm"`
	Blizzard struct {
		Text               string     `xml:",chardata"`
		BlizzardInProgress XMLBool    `xml:"blizzard-in-progress"`
		TimeToNextBlizzard XMLFloat64 `xml:"time-to-next-blizzard"`
		Time               XMLFloat64 `xml:"time"`
		BlizzardTime       XMLFloat64 `xml:"blizzard-time"`
	} `xml:"blizzard"`
	SolarFlare struct {
		Text                 string     `xml:",chardata"`
		SolarFlareInProgress XMLBool    `xml:"solar-flare-in-progress"`
		TimeToNextSolarFlare XMLFloat64 `xml:"time-to-next-solar-flare"`
		Time                 XMLFloat64 `xml:"time"`
		SolarFlareTime       XMLFloat64 `xml:"solar-flare-time"`
	} `xml:"solar-flare"`
	Colony struct {
		Text          string     `xml:",chardata"`
		ExtraPrestige XMLInt64   `xml:"extra-prestige"`
		GameTime      XMLFloat64 `xml:"game-time"`
		RealGameTime  XMLFloat64 `xml:"real-game-time"`
		Name          XMLString  `xml:"name"`
		Latitude      XMLInt64   `xml:"latitude"`
		Longitude     XMLInt64   `xml:"longitude"`
	} `xml:"colony"`
	ShipManager struct {
		Text                         string     `xml:",chardata"`
		Type                         string     `xml:"type,attr"`
		TimeSinceLastColonistLanding XMLFloat64 `xml:"time-since-last-colonist-landing"`
		TimeSinceLastVisitorLanding  XMLFloat64 `xml:"time-since-last-visitor-landing"`
		TimeSinceLastMerchantLanding XMLFloat64 `xml:"time-since-last-merchant-landing"`
		TimeToNextIntruder           XMLFloat64 `xml:"time-to-next-intruder"`
		LandingPermissions           struct {
			Text                string   `xml:",chardata"`
			ColonistsAllowed    XMLBool  `xml:"colonists-allowed"`
			MerchantsAllowed    XMLBool  `xml:"merchants-allowed"`
			VisitorsAllowed     XMLBool  `xml:"visitors-allowed"`
			WorkerPercentage    XMLInt64 `xml:"Worker-percentage"`
			BiologistPercentage XMLInt64 `xml:"Biologist-percentage"`
			EngineerPercentage  XMLInt64 `xml:"Engineer-percentage"`
			MedicPercentage     XMLInt64 `xml:"Medic-percentage"`
			GuardPercentage     XMLInt64 `xml:"Guard-percentage"`
		} `xml:"landing-permissions"`
	} `xml:"ship-manager"`
	Stats struct {
		Text    string `xml:",chardata"`
		Counter []struct {
			Text     string    `xml:",chardata"`
			Type     string    `xml:"type,attr"`
			TypeName XMLString `xml:"type-name"`
			Counts   XMLString `xml:"counts"`
		} `xml:"counter"`
	} `xml:"stats"`
	VisitorEvents VisitorEvents `xml:"visitor-events"`
	GameHints     struct {
		Text       string    `xml:",chardata"`
		ShownHints XMLString `xml:"shown-hints"`
	} `xml:"game-hints"`
	MeteorManager struct {
		Text  string    `xml:",chardata"`
		Seeds XMLString `xml:"seeds"`
	} `xml:"meteor-manager"`
	ManufactureLimits struct {
		Text                 string   `xml:",chardata"`
		CarrierLimit         XMLInt64 `xml:"Carrier-limit"`
		ConstructorLimit     XMLInt64 `xml:"Constructor-limit"`
		DrillerLimit         XMLInt64 `xml:"Driller-limit"`
		MedicalSuppliesLimit XMLInt64 `xml:"MedicalSupplies-limit"`
		SparesLimit          XMLInt64 `xml:"Spares-limit"`
		SemiconductorsLimit  XMLInt64 `xml:"Semiconductors-limit"`
		GunLimit             XMLInt64 `xml:"Gun-limit"`
	} `xml:"manufacture-limits"`
	ChallengeManager string `xml:"challenge-manager"`
	Constructions    struct {
		Text         string         `xml:",chardata"`
		Construction []Construction `xml:"construction"`
	} `xml:"constructions"`
	Characters struct {
		Text      string      `xml:",chardata"`
		Character []Character `xml:"character"`
	} `xml:"characters"`
	Resources    Resources `xml:"resources"`
	Ships        string    `xml:"ships"`
	Interactions struct {
		Text        string `xml:",chardata"`
		Interaction []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Character struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				TypeName string `xml:"type-name,attr"`
			} `xml:"character"`
			Selectable struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				TypeName string `xml:"type-name,attr"`
			} `xml:"selectable"`
			InteractionPoint struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"interaction-point"`
			Stage struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"stage"`
			StageProgress struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"stage-progress"`
			Target struct {
				Text string `xml:",chardata"`
				X    string `xml:"x,attr"`
				Y    string `xml:"y,attr"`
				Z    string `xml:"z,attr"`
			} `xml:"target"`
		} `xml:"interaction"`
	} `xml:"interactions"`
	Screenshot string `xml:"screenshot"`
}
