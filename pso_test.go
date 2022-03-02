package gpso

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	pso := New(10, 2)
	fmt.Println(pso)

	for i := 0; i < pso.m; i++ {
		fmt.Println(pso.population[i].position)
	}
}

func TestNew2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pso := New(20, 2,
		WithT(500),
		WithC1(0.1), WithC2(0.1), WithW(0.8),
		WithLowerLimit([]float64{0, 0}), WithUpperLimit([]float64{5, 5}),
		WithLowerVelocity([]float64{0, 0}), WithUpperVelocity([]float64{0.1, 0.1}),
		WithObjectFunc(func(fs []float64) float64 {
			return math.Pow(fs[0] - 3.14, 2) + math.Pow(fs[1] - 2.72, 2) + math.Sin(3*fs[0]+1.41) + math.Sin(4*fs[1] - 1.73)
		}))

	pso.Run()

	for i, particle := range pso.HistoryBests() {
		fmt.Printf("%diter %v, fitness = %f\n", i, particle.Values(), particle.Fitness())
	}
}