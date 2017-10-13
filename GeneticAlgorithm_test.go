package ga

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestGeneticAlgorithm(t *testing.T) {
	rand.Seed(3)
	SetCrossoverFunc(DefaultCrossoverFunc)
	SetMutateFunc(DefaultMutateFunc)
	fmt.Println(GeneticAlgorithm(10, 10, 50, true, true, false))
}

func TestFillRandomPopulation(t *testing.T) {
	SetCrossoverFunc(DefaultCrossoverFunc)
	SetMutateFunc(DefaultMutateFunc)

	population := make([]Genome, 0)
	population = FillRandomPopulation(population, 10, 10)

	if len(population) != 10 {
		t.Error("Population not filled to specified size")
	}

	for _, val := range population {
		if len(val.Sequence) == 0 {
			t.Error("Bitstrings in population are empty")
		}
	}
}

func TestTournament(t *testing.T) {
	population := []Genome{
		{[]int{1, 1, 1, 1}},
		{[]int{0, 1, 1, 1}},
		{[]int{0, 0, 1, 1}},
		{[]int{0, 0, 0, 1}},
	}
	avgFitnessBefore := AverageFitness(population)
	population = Tournament(population)
	avgFitnessAfter := AverageFitness(population)

	if avgFitnessAfter < avgFitnessBefore {
		t.Error("Fitness decreased after tournament")
	}
}
