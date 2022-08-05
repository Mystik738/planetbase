# Planetbase Base Templates

This is a utility for creating Planetbase .sav files from a template. 

```
Flags:
  -d float
        The distance to space buildings. (default 25)
  -h    Display the help.
  -la int
        The latitude to use. (default 0)
  -lo int
        The longitude to use. (default -128)
  -o string
        The output filename to use. (default "save.sav")
  -p int
        The planet to use, between 0 and 3. (default 3)
  -r    Read the save instead and output some information from it.
  -t string
        The template filename to use. (default "template.txt")
```

A small template looks like the following:

```
Ox==Ai
 \\//
  So
```

Where Ox, Ai, and So are modules and ==, //, and \\\\ are connections between them. Modules are expected to be connected hexagonally. Module size is determined by the binary representation of the capitalization in the template, i.e. ox, oX, Ox, and OX are 00, 01, 10, 11 or zero through three, respectively, starting at the module's smallest size. Anything above the module's largest size will remain at its largest size. The full list of module values are:
- ai: Airlock
- an: AntiMeteorLaser
- ba: Bar
- bi: BioDome
- bs: BasePad
- ca: Cabin
- cn: Canteen
- co: ControlCenter
- do: Dorm
- fa: Factory
- la: Lab
- li: LightningRod
- ln: LandingPad
- mi: Mine
- mo: Monolith
- mu: MultiDome
- ox: OxygenGenerator
- po: PowerCollector
- pr: ProcessingPlant
- py: Pyramid
- ra: RadioAntenna
- ro: RoboticsFacility
- sg: Signpost
- si: SickBay
- so: SolarPanel
- sr: Storage
- st: Starport
- te: Telescope
- wa: WaterExtractor
- wi: WindTurbine
- wt: WaterTank