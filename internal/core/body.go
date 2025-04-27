package core

import "github.com/gen2brain/raylib-go/raylib"

type Body struct {
	Name string
	Position Vector3
	Velocity Vector3
	Mass float64
	Radius float64
	Color rl.Color
}