package graphics

import (
	"github.com/Cladkoewka/gravity-simulation/internal/core"

	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

const scale = 2e10
var cnt = 0

func DrawBodies(bodies []*core.Body) {
	for _, body := range bodies {
			pos := rl.NewVector3(
					float32(body.Position.X/scale),
					float32(body.Position.Y/scale),
					float32(body.Position.Z/scale),
			)
			rl.DrawSphere(pos, float32(body.Radius), body.Color)
			debug(body, pos)
	}
}

func debug(b *core.Body, p rl.Vector3) {
	cnt ++
	if cnt >= 60 {
		fmt.Printf("Name: %s Position: %v \n", b.Name, p)
		cnt = 0
	}
}