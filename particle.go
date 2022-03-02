package gpso

import "math/rand"

type Particle struct {
	position []float64
	velocity []float64
	fitness float64
}

func (p *Particle) calculate(f func([]float64)float64) float64 {
	p.fitness = f(p.position)
	return p.fitness
}

func (p Particle) copy() *Particle {
	newParticle := &Particle{
		position: make([]float64, len(p.position)),
		velocity: make([]float64, len(p.velocity)),
		fitness:  p.fitness,
	}
	copy(newParticle.position, p.position)
	copy(newParticle.velocity, p.velocity)
	return newParticle
}

func (p Particle) Fitness() float64 {
	return p.fitness
}

func (p Particle) Values() []float64 {
	return p.position
}

func (p *Particle) UpdateVelocity(w float64, c1, c2 float64, hbest, gbest *Particle, lowerVelocity, upperVelocity []float64) {
	for i := 0; i < len(p.position); i++ {
		p.velocity[i] = w * p.velocity[i] + c1 * rand.Float64() * (hbest.velocity[i] - p.velocity[i]) + c2 * rand.Float64() * (gbest.velocity[i] - p.velocity[i])
		if p.velocity[i] > upperVelocity[i] {
			p.velocity[i] = upperVelocity[i]
		}
		if p.velocity[i] < lowerVelocity[i] {
			p.velocity[i] = lowerVelocity[i]
		}
	}
}

func (p *Particle) UpdatePosition(lowerLimit, upperLimit []float64) {
	for i := 0; i < len(p.position); i++ {
		p.position[i] = p.position[i] + p.velocity[i]
		if p.position[i] > upperLimit[i] {
			p.position[i] = upperLimit[i]
		}
		if p.position[i] < lowerLimit[i] {
			p.position[i] = lowerLimit[i]
		}
	}
}

func NewParticle(n int, lowerLimit, upperLimit, lowerVelocity, upperVelocity []float64) *Particle {
	produce := func(lower, upper []float64) []float64 {
		diff := float64sSub(upper, lower)
		ret := make([]float64, n)
		for i := 0; i < n; i++ {
			ret[i] = lower[i] + rand.Float64() * diff[i]
		}
		return ret
	}
	p := &Particle{
		position: produce(lowerLimit, upperLimit),
		velocity: produce(lowerVelocity, upperVelocity),
	}
	return p
}

func particles(m, n int, lowerLimit, upperLimit, lowerVelocity, upperVelocity []float64) []*Particle {
	ps := make([]*Particle, m)
	for i := 0; i < m; i++ {
		ps[i] = NewParticle(n, lowerLimit, upperLimit, lowerVelocity, upperVelocity)
	}
	return ps
}