package engine

import (
	"fmt"

	"github.com/Cladkoewka/gravity-simulation/internal/core"
	"github.com/Cladkoewka/gravity-simulation/internal/graphics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Engine struct {
	World *core.World
	Camera rl.Camera
	TimeAcceleration float64
}
const timeAcceleration = 60 * 60 * 24
var cnt = 0


func NewEngine(world *core.World) *Engine {
	camera := rl.Camera{
		Position: rl.NewVector3(0, 200, 400),
		Target:  rl.NewVector3(0, 0, 0),
		Up:      rl.NewVector3(0, 1, 0),
		Fovy: 	45.0,
		Projection: 	rl.CameraPerspective,
	}

	return &Engine{
		World: world,
		Camera: camera,
		TimeAcceleration: 60 * 60 * 24,
	}
}

func (e *Engine) Run() {
	rl.InitWindow(1280, 720, "Gravity Simulation")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		scaleDt := float64(dt) * e.TimeAcceleration
		e.World.Update(scaleDt)

		//rl.UpdateCamera(&e.Camera, rl.CameraOrbital)

		e.handleTimeSlider()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(e.Camera)
		rl.DrawGrid(1000, 50)
		graphics.DrawBodies(e.World.Bodies)

		//debug(e)
		rl.EndMode3D()

		rl.DrawText("Gravity Simulation", 10, 10, 20, rl.White)
		rl.DrawText(fmt.Sprintf("Time x%.0f", e.TimeAcceleration), 10, 40, 20, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (e *Engine) handleTimeSlider() {
	var sliderX int32 = 200
	var sliderY int32 = 680
	var sliderWidth int32 = 400
	var sliderHeight int32 = 10

	rl.DrawRectangle(int32(sliderX), int32(sliderY), int32(sliderWidth), int32(sliderHeight), rl.Gray)

	normalized := e.TimeAcceleration / (60 * 60 * 24 * 365)
	if normalized > 1 {
		normalized = 1
	}
	if normalized < 0 {
		normalized = 0
	}
	handleX := sliderX + int32(normalized*float64(sliderWidth))

	rl.DrawRectangle(int32(handleX-5), int32(sliderY-5), 10, 20, rl.White)

	// Перетаскивание
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		if mouseY >= sliderY-10 && mouseY <= sliderY+20 {
			if mouseX < int32(sliderX) {
				mouseX = int32(sliderX)
			}
			if mouseX > int32(sliderX+sliderWidth) {
				mouseX = int32(sliderX + sliderWidth)
			}
			newNormalized := float64(mouseX - int32(sliderX)) / float64(sliderWidth)
			e.TimeAcceleration = (60 * 60 * 24) + newNormalized*(60*60*24*365) 
		}
	}
}

func debug(e *Engine) {
	cnt ++
	if cnt >= 60 {
		for _, body := range e.World.Bodies {
			fmt.Printf("Name: %s Position: %v Velocity: %v \n", body.Name, body.Position, body.Velocity)
		}
		cnt = 0
	}
}