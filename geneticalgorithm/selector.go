package geneticalgorithm

import (
	"math"
	"math/rand"
)

type OptimizeFunction func(DiscreteIndividuum) float64

type Select func(individuals <-chan DiscreteIndividuum, selected chan<- DiscreteIndividuum, window int, optimizer OptimizeFunction)

func DummySelector(individuals <-chan DiscreteIndividuum, selected chan<- DiscreteIndividuum, window int, optimizer OptimizeFunction) {
	for individuum := range individuals {
		selected <- individuum
	}
}

func collectSample(individuals <-chan DiscreteIndividuum, window int, optimizer OptimizeFunction) (sampleWindow []DiscreteIndividuum, distribution []float64) {
	sampleWindow = make([]DiscreteIndividuum, window)
	distribution = make([]float64, window)
	var sum float64
	for i := 0; i < window; i++ {
		individuum := <-individuals
		fitness := optimizer(individuum)
		sampleWindow[i] = individuum
		distribution[i] = fitness
		sum += fitness
	}
	for i := range distribution {
		distribution[i] *= 1 / sum
	}
	return
}

func RemainderStochasticSampling(individuals <-chan DiscreteIndividuum, selected chan<- DiscreteIndividuum, window int, optimizer OptimizeFunction) {
	for {
		sampleWindow, distribution := collectSample(individuals, window, optimizer)
		var totalAmout int
		for i, individuum := range sampleWindow {
			amount := int(math.Floor(distribution[i] * float64(window)))
			for j := 0; j < amount; j++ {
				selected <- individuum
			}
			totalAmout += amount
		}
		for i := 0; i < (window - totalAmout); i++ {
			selected <- sampleWindow[rand.Intn(window)]
		}
	}
}
