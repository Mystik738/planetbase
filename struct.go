package main

import "encoding/xml"

type AutoInc struct {
	ID int64
}

func NextID(ai *AutoInc) int64 {
	ai.ID = ai.ID + 1
	return ai.ID
}

type Position struct {
	Text string  `xml:",chardata"`
	X    float64 `xml:"x,attr"`
	Y    float64 `xml:"y,attr"`
	Z    float64 `xml:"z,attr"`
}

type Tech struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
}

type VisitorEvents struct {
	Text          string `xml:",chardata"`
	NextEventTime struct {
		Text  string  `xml:",chardata"`
		Value float64 `xml:"value,attr"`
	} `xml:"next-event-time"`
	VisitorEvent VisitorEvent `xml:"visitor-event"`
}

type VisitorEvent struct {
	Text         string `xml:",chardata"`
	Type         string `xml:"type,attr"`
	VisitorCount struct {
		Text  string `xml:",chardata"`
		Value int64  `xml:"value,attr"`
	} `xml:"visitor-count"`
}

type Construction struct {
	Text    string `xml:",chardata"`
	Type    string `xml:"type,attr"`
	Enabled struct {
		Text  string `xml:",chardata"`
		Value bool   `xml:"value,attr"`
	} `xml:"enabled"`
	State struct {
		Text  string `xml:",chardata"`
		Value int64  `xml:"value,attr"`
	} `xml:"state"`
	BuildProgress struct {
		Text  string  `xml:",chardata"`
		Value float64 `xml:"value,attr"`
	} `xml:"build-progress"`
	Condition struct {
		Text  string  `xml:",chardata"`
		Value float64 `xml:"value,attr"`
	} `xml:"condition"`
	Oxygen struct {
		Text  string  `xml:",chardata"`
		Value float64 `xml:"value,attr"`
	} `xml:"oxygen"`
	ID          XMLInt64 `xml:"id"`
	Position    Position `xml:"position"`
	Orientation Position `xml:"orientation"`
	TimeBuilt   struct {
		Text  string  `xml:",chardata"`
		Value float64 `xml:"value,attr"`
	} `xml:"time-built"`
	Locked struct {
		Text  string `xml:",chardata"`
		Value bool   `xml:"value,attr"`
	} `xml:"locked"`
	HighPriority struct {
		Text  string `xml:",chardata"`
		Value bool   `xml:"value,attr"`
	} `xml:"high-priority"`
	ModuleType struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"module-type"`
	SizeIndex struct {
		Text  string `xml:",chardata"`
		Value int64  `xml:"value,attr"`
	} `xml:"size-index"`
	MobileRotation     Position            `xml:"mobile-rotation"`
	PowerStorage       *PowerStorage       `xml:"power-storage"`
	Components         *Components         `xml:"components"`
	ProductionProgress *ProductionProgress `xml:"production-progress"`
	ResourceStorage    *ResourceStorage    `xml:"resource-storage"`
	LaserCharge        *LaserCharge        `xml:"laser-charge"`
	Links              *Links              `xml:"links"`
}

type Components struct {
	Text      string `xml:",chardata"`
	Component []struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		Enabled []struct {
			Text  string `xml:",chardata"`
			Value bool   `xml:"value,attr"`
		} `xml:"enabled"`
		State struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"state"`
		BuildProgress struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"build-progress"`
		ID struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"id"`
		ComponentType struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"component-type"`
		Position    Position `xml:"position"`
		Orientation Position `xml:"orientation"`
		Condition   struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"condition"`
		ProductionProgress struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"production-progress"`
		Time struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"time"`
		ProducedItemIndex struct {
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

type ProductionProgress struct {
	Text  string  `xml:",chardata"`
	Value float64 `xml:"value,attr"`
}

type ResourceStorage struct {
	Text string `xml:",chardata"`
	Slot []struct {
		Text      string   `xml:",chardata"`
		Position  Position `xml:"position"`
		MaxHeight struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"max-height"`
		Resource []Resource `xml:"resource"`
	} `xml:"slot"`
}

type LaserCharge struct {
	Text  string  `xml:",chardata"`
	Value float64 `xml:"value,attr"`
}

type Links struct {
	Text string     `xml:",chardata"`
	ID   []XMLInt64 `xml:"id"`
}

type XMLInt64 struct {
	Text  string `xml:",chardata"`
	Value int64  `xml:"value,attr"`
}

type PowerStorage struct {
	Text  string  `xml:",chardata"`
	Value float64 `xml:"value,attr"`
}

type Resources struct {
	Text                string              `xml:",chardata"`
	Resource            []Resource          `xml:"resource"`
	InmaterialResources InmaterialResources `xml:"inmaterial-resources"`
}

type Resource struct {
	Text     string   `xml:",chardata"`
	Type     string   `xml:"type,attr"`
	ID       XMLInt64 `xml:"id"`
	TraderID struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"trader-id"`
	Position    Position `xml:"position"`
	Orientation Position `xml:"orientation"`
	State       struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"state"`
	Location struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"location"`
	Subtype struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"subtype"`
	Condition struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"condition"`
	Durability struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"durability"`
}

type InmaterialResources struct {
	Text          string `xml:",chardata"`
	ContainerName struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value,attr"`
	} `xml:"container-name"`
	Amount Amount `xml:"amount"`
}

type Amount struct {
	Text         string       `xml:",chardata"`
	ResourceType ResourceType `xml:"resource-type"`
	Amount       struct {
		Text  string `xml:",chardata"`
		Value int64  `xml:"value,attr"`
	} `xml:"amount"`
}

type ResourceType struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value,attr"`
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
		Text        string `xml:",chardata"`
		PlanetIndex struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"planet-index"`
	} `xml:"planet"`
	Milestones struct {
		Text      string `xml:",chardata"`
		Milestone []struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"milestone"`
	} `xml:"milestones"`
	Techs struct {
		Text string `xml:",chardata"`
		Tech []Tech `xml:"tech"`
	} `xml:"techs"`
	Environment struct {
		Text      string `xml:",chardata"`
		TimeOfDay struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-of-day"`
		WindIndicator struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"wind-indicator"`
	} `xml:"environment"`
	Terrain struct {
		Text string `xml:",chardata"`
		Seed struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"seed"`
	} `xml:"terrain"`
	Camera struct {
		Text   string `xml:",chardata"`
		Height struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"height"`
		Position    Position `xml:"position"`
		Orientation Position `xml:"orientation"`
	} `xml:"camera"`
	Sandstorm struct {
		Text                string `xml:",chardata"`
		SandstormInProgress struct {
			Text  string `xml:",chardata"`
			Value bool   `xml:"value,attr"`
		} `xml:"sandstorm-in-progress"`
		TimeToNextSandstorm struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-to-next-sandstorm"`
		Time struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time"`
		SandstormTime struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"sandstorm-time"`
	} `xml:"sandstorm"`
	Blizzard struct {
		Text               string `xml:",chardata"`
		BlizzardInProgress struct {
			Text  string `xml:",chardata"`
			Value bool   `xml:"value,attr"`
		} `xml:"blizzard-in-progress"`
		TimeToNextBlizzard struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-to-next-blizzard"`
		Time struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time"`
		BlizzardTime struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"blizzard-time"`
	} `xml:"blizzard"`
	SolarFlare struct {
		Text                 string `xml:",chardata"`
		SolarFlareInProgress struct {
			Text  string `xml:",chardata"`
			Value bool   `xml:"value,attr"`
		} `xml:"solar-flare-in-progress"`
		TimeToNextSolarFlare struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-to-next-solar-flare"`
		Time struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time"`
		SolarFlareTime struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"solar-flare-time"`
	} `xml:"solar-flare"`
	Colony struct {
		Text          string `xml:",chardata"`
		ExtraPrestige struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"extra-prestige"`
		GameTime struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"game-time"`
		RealGameTime struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"real-game-time"`
		Name struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"name"`
		Latitude struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"latitude"`
		Longitude struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"longitude"`
	} `xml:"colony"`
	ShipManager struct {
		Text                         string `xml:",chardata"`
		Type                         string `xml:"type,attr"`
		TimeSinceLastColonistLanding struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-since-last-colonist-landing"`
		TimeSinceLastVisitorLanding struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-since-last-visitor-landing"`
		TimeSinceLastMerchantLanding struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-since-last-merchant-landing"`
		TimeToNextIntruder struct {
			Text  string  `xml:",chardata"`
			Value float64 `xml:"value,attr"`
		} `xml:"time-to-next-intruder"`
		LandingPermissions struct {
			Text             string `xml:",chardata"`
			ColonistsAllowed struct {
				Text  string `xml:",chardata"`
				Value bool   `xml:"value,attr"`
			} `xml:"colonists-allowed"`
			MerchantsAllowed struct {
				Text  string `xml:",chardata"`
				Value bool   `xml:"value,attr"`
			} `xml:"merchants-allowed"`
			VisitorsAllowed struct {
				Text  string `xml:",chardata"`
				Value bool   `xml:"value,attr"`
			} `xml:"visitors-allowed"`
			WorkerPercentage struct {
				Text  string `xml:",chardata"`
				Value int64  `xml:"value,attr"`
			} `xml:"Worker-percentage"`
			BiologistPercentage struct {
				Text  string `xml:",chardata"`
				Value int64  `xml:"value,attr"`
			} `xml:"Biologist-percentage"`
			EngineerPercentage struct {
				Text  string `xml:",chardata"`
				Value int64  `xml:"value,attr"`
			} `xml:"Engineer-percentage"`
			MedicPercentage struct {
				Text  string `xml:",chardata"`
				Value int64  `xml:"value,attr"`
			} `xml:"Medic-percentage"`
			GuardPercentage struct {
				Text  string `xml:",chardata"`
				Value int64  `xml:"value,attr"`
			} `xml:"Guard-percentage"`
		} `xml:"landing-permissions"`
	} `xml:"ship-manager"`
	Stats struct {
		Text    string `xml:",chardata"`
		Counter []struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"type,attr"`
			TypeName struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"type-name"`
			Counts struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"counts"`
		} `xml:"counter"`
	} `xml:"stats"`
	VisitorEvents VisitorEvents `xml:"visitor-events"`
	GameHints     struct {
		Text       string `xml:",chardata"`
		ShownHints struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"shown-hints"`
	} `xml:"game-hints"`
	MeteorManager struct {
		Text  string `xml:",chardata"`
		Seeds struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"seeds"`
	} `xml:"meteor-manager"`
	ManufactureLimits struct {
		Text         string `xml:",chardata"`
		CarrierLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Carrier-limit"`
		ConstructorLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Constructor-limit"`
		DrillerLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Driller-limit"`
		MedicalSuppliesLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"MedicalSupplies-limit"`
		SparesLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Spares-limit"`
		SemiconductorsLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Semiconductors-limit"`
		GunLimit struct {
			Text  string `xml:",chardata"`
			Value int64  `xml:"value,attr"`
		} `xml:"Gun-limit"`
	} `xml:"manufacture-limits"`
	ChallengeManager string `xml:"challenge-manager"`
	Constructions    struct {
		Text         string         `xml:",chardata"`
		Construction []Construction `xml:"construction"`
	} `xml:"constructions"`
	Characters struct {
		Text      string `xml:",chardata"`
		Character []struct {
			Text        string   `xml:",chardata"`
			Type        string   `xml:"type,attr"`
			Position    Position `xml:"position"`
			Orientation Position `xml:"orientation"`
			Location    struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"location"`
			Name struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"name"`
			Specialization struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"specialization"`
			StatusFlags struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"status-flags"`
			State struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"state"`
			ID struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"id"`
			WanderTime struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"wander-time"`
			Health struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Health"`
			Nutrition struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Nutrition"`
			Hydration struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Hydration"`
			Oxygen struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Oxygen"`
			Sleep struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Sleep"`
			Morale struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Morale"`
			Gender struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"gender"`
			BasicMealCount struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"basic-meal-count"`
			HeadIndex struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"head-index"`
			SkinColorIndex struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"skin-color-index"`
			HairColorIndex struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"hair-color-index"`
			Doctor struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"doctor"`
			InmunityToContagionTime struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"inmunity-to-contagion-time"`
			LoadedResource struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"loaded-resource"`
			Condition struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Condition"`
			Integrity struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"Integrity"`
			IntegrityDecayRate struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"integrity-decay-rate"`
		} `xml:"character"`
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
