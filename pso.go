package gpso

import (
	"math"
)

type Option func(*PSO)

func WithT(t int) Option {
	return func(p *PSO) {
		p.t = t
	}
}

func WithC1(c1 float64) Option {
	return func(p *PSO) {
		p.c1 = c1
	}
}

func WithC2(c2 float64) Option {
	return func(p *PSO) {
		p.c2 = c2
	}
}

func WithW(w float64) Option {
	return func(p *PSO) {
		p.w = w
	}
}

func WithLowerLimit(limit []float64) Option {
	return func(p *PSO) {
		p.lowerLimit = limit
	}
}

func WithUpperLimit(limit []float64) Option {
	return func(p *PSO) {
		p.upperLimit = limit
	}
}

func WithLowerVelocity(velocity []float64) Option {
	return func(p *PSO) {
		p.lowerVelocity = velocity
	}
}

func WithUpperVelocity(velocity []float64) Option {
	return func(p *PSO) {
		p.upperVelocity = velocity
	}
}

func WithObjectFunc(objectFunc func([]float64) float64) Option {
	return func(p *PSO) {
		p.objectFunc = objectFunc
	}
}

type PSO struct {
	m             int
	n             int
	t             int
	c1            float64
	c2            float64
	w             float64
	objectFunc    func([]float64) float64
	lowerLimit    []float64
	upperLimit    []float64
	lowerVelocity []float64
	upperVelocity []float64
	gbest         *Particle
	lbest         []*Particle
	population    []*Particle
	historyBests  []*Particle
}

func (p *PSO) setGlobalBest(i int) {
	p.gbest = p.population[i].copy()
}

func (p *PSO) Run() {
	fitnessValues := make([]float64, p.m)
	sortedIndexes := make([]int, p.m)
	for t := 0; t < p.t; t++ {
		for i := 0; i < p.m; i++ {
			p.population[i].UpdateVelocity(p.w, p.c1, p.c2, p.lbest[i], p.gbest, p.lowerVelocity, p.upperVelocity)
			p.population[i].UpdatePosition(p.lowerLimit, p.upperLimit)
			if p.population[i].calculate(p.objectFunc) < p.lbest[i].fitness {
				p.lbest[i] = p.population[i].copy()
			}
		}

		for i := 0; i < p.m; i++ {
			fitnessValues[i] = p.population[i].calculate(p.objectFunc)
		}
		sortedIndexes = argsortFunc(fitnessValues)
		p.setGlobalBest(sortedIndexes[0])
		p.historyBests[t] = p.population[sortedIndexes[0]].copy()
	}
}

func (p PSO) HistoryBests() []*Particle {
	return p.historyBests
}

func New(m, n int, options ...Option) *PSO {
	pso := &PSO{
		m:  m,
		n:  n,
		t:  100,
		w:  0.5,
		c1: 0.5,
		c2: 0.5,
		objectFunc: func(fs []float64) float64 {
			var fitness float64
			for _, f := range fs {
				fitness = math.Pow(f, 2)
			}
			return fitness
		},
		lowerLimit:    floats(-1, n),
		upperLimit:    floats(1, n),
		lowerVelocity: floats(-0.5, n),
		upperVelocity: floats(0.5, n),
		gbest:         nil,
		lbest:         make([]*Particle, m),
	}

	for _, option := range options {
		option(pso)
	}

	pso.population = particles(pso.m, pso.n, pso.lowerLimit, pso.upperLimit, pso.lowerVelocity, pso.upperVelocity)
	for i, particle := range pso.population {
		pso.lbest[i] = particle.copy()
	}

	fitnessValues := make([]float64, pso.m)
	for i := 0; i < pso.m; i++ {
		fitnessValues[i] = pso.population[i].calculate(pso.objectFunc)
	}
	sortedIndexes := argsortFunc(fitnessValues)
	pso.setGlobalBest(sortedIndexes[sortedIndexes[0]])
	pso.historyBests = make([]*Particle, pso.t)

	return pso
}
