package core

import "sync"

type World struct {
	Bodies []*Body
}

func (w *World) Update(dt float64) {
	var wg sync.WaitGroup
	forces := make([]Vector3, len(w.Bodies))

	for i := range w.Bodies {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var totalForce Vector3
			for j := range w.Bodies {
				if i == j {
					continue
				}
				totalForce = totalForce.Add(GravitationalForce(w.Bodies[i], w.Bodies[j]))
			}
			forces[i] = totalForce
		}(i)
	}
	wg.Wait()

	for i, b := range w.Bodies {
		acceleration := forces[i].Div(b.Mass)
		b.Velocity = b.Velocity.Add(acceleration.Mul(dt))
		b.Position = b.Position.Add(b.Velocity.Mul(dt))
	}
}