package main

import (
	"math"

	"github.com/Cladkoewka/gravity-simulation/internal/core"
	"github.com/Cladkoewka/gravity-simulation/internal/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	G             = 6.67430e-11
	sunMass       = 1.989e30
	earthMass     = 5.972e24
	mercuryMass   = 3.301e23
	venusMass     = 4.867e24
	marsMass      = 6.417e23
	jupiterMass   = 1.898e27
	saturnMass    = 5.683e26
	uranusMass    = 8.681e25
	neptuneMass   = 1.024e26
	plutoMass = 1.309e22
)

func main() {
	earthOrbitRadius := 1.496e11      
	mercuryOrbitRadius := 5.791e10       
	venusOrbitRadius := 1.082e11         
	marsOrbitRadius := 2.279e11
	jupiterOrbitRadius := 7.785e11
	saturnOrbitRadius := 1.433e12
	uranusOrbitRadius := 2.877e12
	neptuneOrbitRadius := 4.503e12
	plutoOrbitRadius := 5.906e12

	earthOrbitVelocity := math.Sqrt(G * sunMass / earthOrbitRadius)
	mercuryOrbitVelocity := math.Sqrt(G * sunMass / mercuryOrbitRadius)
	venusOrbitVelocity := math.Sqrt(G * sunMass / venusOrbitRadius)
	marsOrbitVelocity := math.Sqrt(G * sunMass / marsOrbitRadius)
	jupiterOrbitVelocity := math.Sqrt(G * sunMass / jupiterOrbitRadius)
	saturnOrbitVelocity := math.Sqrt(G * sunMass / saturnOrbitRadius)
	uranusOrbitVelocity := math.Sqrt(G * sunMass / uranusOrbitRadius)
	neptuneOrbitVelocity := math.Sqrt(G * sunMass / neptuneOrbitRadius)
	plutoOrbitVelocity := math.Sqrt(G * sunMass / plutoOrbitRadius)


	world := &core.World{
		Bodies: []*core.Body{
			{
				Name: "Sun",
				Position: core.Vector3{0, 0, 0},
				Mass: sunMass,
				Radius: 5.0,
				Color: rl.Yellow,
			},
			{
				Name: "Mercury",
				Position: core.Vector3{mercuryOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, mercuryOrbitVelocity},
				Mass: mercuryMass,
				Radius: 1.0,
				Color: rl.LightGray,
			},
			{
				Name: "Venus",
				Position: core.Vector3{venusOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, venusOrbitVelocity},
				Mass: venusMass,
				Radius: 1.5,
				Color: rl.Orange,
			},
			{
				Name: "Earth",
				Position: core.Vector3{earthOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, earthOrbitVelocity},
				Mass: earthMass,
				Radius: 2.0,
				Color: rl.Blue,
			},
			{
				Name: "Mars",
				Position: core.Vector3{marsOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, marsOrbitVelocity},
				Mass: marsMass,
				Radius: 1.3,
				Color: rl.Red,
			},
			{
				Name: "Jupiter",
				Position: core.Vector3{jupiterOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, jupiterOrbitVelocity},
				Mass: jupiterMass,
				Radius: 5.0,
				Color: rl.Orange,
			},
			{
				Name: "Saturn",
				Position: core.Vector3{saturnOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, saturnOrbitVelocity},
				Mass: saturnMass,
				Radius: 2.5,
				Color: rl.Beige,
			},
			{
				Name: "Uranus",
				Position: core.Vector3{uranusOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, uranusOrbitVelocity},
				Mass: uranusMass,
				Radius: 2.0,
				Color: rl.SkyBlue,
			},
			{
				Name: "Neptune",
				Position: core.Vector3{neptuneOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, neptuneOrbitVelocity},
				Mass: neptuneMass,
				Radius: 2.0,
				Color: rl.DarkBlue,
			},
			{
				Name: "Pluto",
				Position: core.Vector3{plutoOrbitRadius, 0, 0},
				Velocity: core.Vector3{0, 0, plutoOrbitVelocity},
				Mass: plutoMass,
				Radius: 1, 
				Color: rl.Gray,
			},
		},
	}

	e := engine.NewEngine(world)
	e.Run()
}
