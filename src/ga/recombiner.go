package ga

import (
	"math/rand"
)

type Recombine func(parents <-chan Individuum, children chan<- Individuum, popsize int)

func DummyRecombiner(parents <-chan Individuum, children chan<- Individuum, popsize int) {
	for i:=0; i<popsize; i++ {
		children <- <-parents
	}
}

func OnePointCrossOver(parents <-chan Individuum, children chan<- Individuum, popsize int) {
	i := 0
	for {
		if popsize > 0 {
			if i >= popsize {
				break
			}
		}
		parent1 := <-parents
		parent2 := <-parents
		child := MakeIndividuum(parent1)
		for i := range child {
			mask := 0x1FFFF >> uint(rand.Intn(16) + 1)
			child[i] = parent2[i] & mask | parent1[i] & (mask^0xFFFFFFFF)
		}
		children <- child
		i++
	}
}