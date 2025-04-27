package core

const G = 6.67430e-11 // Gravitational constant 

func GravitationalForce(b1, b2 *Body) Vector3 {
	dir := b2.Position.Sub(b1.Position)
	distSq := dir.X*dir.X + dir.Y*dir.Y + dir.Z*dir.Z
	if distSq < 1e-2 {
		distSq = 1e-2 
	}
	forceMag := G * b1.Mass * b2.Mass / distSq
	return dir.Normalize().Mul(forceMag)
}