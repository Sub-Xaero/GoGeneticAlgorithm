package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	mutateChance = 100
	numStrings   = 10
	strLength    = 8
)

// Genome represents a bitstring and associated fitness value
type Genome struct {
	sequence string
}

// Fitness calculates the suitability of a candidate solution and returns an integral score value
func (gene Genome) Fitness() int {
	return strings.Count(gene.sequence, "1")
}

// Crossover returns bitstring pair which is product of two bitstrings with their tails swapped at a random index
func (gene Genome) Crossover(spouse Genome) []Genome {
	offspring := make([]Genome, 0)

	if len(gene.sequence) != len(spouse.sequence) {
		panic(errors.New("strings are not current length"))
	}

	crossover := rand.Int() % len(gene.sequence)

	offspring = append(offspring, Genome{gene.sequence[0:crossover] + spouse.sequence[crossover:]})
	offspring = append(offspring, Genome{spouse.sequence[0:crossover] + gene.sequence[crossover:]})
	return offspring
}

func (self *Genome) mutate(chance int) Genome {
	mutant := ""
	for _, i := range self.sequence {
		if rand.Int()%chance == 1 {
			if string(i) == "1" {
				mutant += "0"
			} else {
				mutant += "1"
			}
		} else {
			mutant += string(i)
		}
	}
	self.sequence = mutant
	return *self
}

// GenerateBitString returns the highest fitness found in a [] Genome population
func GenerateBitString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("strings cannot be zero-length")
	}
	var bitstring string
	for i := 0; i < length; i++ {
		bitstring += strconv.Itoa(rand.Int() % 2)
	}
	return bitstring, nil
}

// Tournament returns a [] Genome population composed of the best out of randomly selected pairs
func Tournament(population []Genome) []Genome {
	offspring := make([]Genome, 0)

	for i := 0; i < len(population); i++ {
		parent1 := population[rand.Int()%len(population)]
		parent2 := population[rand.Int()%len(population)]

		if parent1.Fitness() > parent2.Fitness() {
			offspring = append(offspring, parent1)
		} else {
			offspring = append(offspring, parent2)
		}
	}

	return offspring
}

func fillRandomPopulation(populace []Genome) []Genome {
	for len(populace) < numStrings {
		str, err := generateBitString(strLength)
		if err == nil {
			populace = append(populace, Genome{str})
		} else {
			panic(errors.New("failed to initialise population"))
		}
	}
	fmt.Println("Initial population:", populace)
	return populace
}

func main() {
	rand.Seed(time.Now().Unix())

	// Init
	population := make([]Genome, 0)
	population = fillRandomPopulation(population)

	for y := 0; y < 100; y++ {
		fmt.Println("Interation", y)
		fmt.Println("Start Population:", population)

		breedingGround := make([]Genome, 0)
		breedingGround = append(breedingGround, Tournament(population)...)
		fmt.Println("Tournament Offspring  :", breedingGround, "Average:", AverageFitness(breedingGround), "Max:", MaxFitness(breedingGround))

		crossoverBreedingGround := make([]Genome, 0)
		for i := 0; i+1 < len(breedingGround); i += 2 {
			crossoverBreedingGround = append(crossoverBreedingGround, breedingGround[i].crossover(breedingGround[i+1])...)
		}
		breedingGround = crossoverBreedingGround
		fmt.Println("Crossover:", breedingGround)

		for _, i := range breedingGround {
			i.mutate(mutateChance)
		}
		fmt.Println("Mutation:", breedingGround)

		population = make([]Genome, 0)
		copy(population, breedingGround)
		population = fillRandomPopulation(population)
		fmt.Println("Fill Population:", population)
		fmt.Println()
	}
}
