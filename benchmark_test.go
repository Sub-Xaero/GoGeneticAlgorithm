package ga

import (
	"testing"
)

var seed int64 = 3
var (
	globalBestCandidate Genome
	globalNumIterations int
	globalPopulation    Population
)

func benchmarkGATournament(length, generations int, terminateEarly bool, b *testing.B) {
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(seed)

	var (
		bestCandidate Genome
		numIterations int
		candidatePool Population
	)
	for n := 0; n < b.N; n++ {
		genA.Run(length, length, generations, true, true, terminateEarly)
		b.Log("Best Candidate", genA.BestCandidate)
		b.Log("Num Iterations:", genA.Generations)
		b.Log("Population:", genA.Candidates)
	}
	globalBestCandidate = bestCandidate
	globalNumIterations = numIterations
	globalPopulation = candidatePool
}

func BenchmarkGATournamentFull_10(b *testing.B)           { benchmarkGATournament(10, 100, false, b) }
func BenchmarkGATournamentFull_20(b *testing.B)           { benchmarkGATournament(20, 200, false, b) }
func BenchmarkGATournamentFull_50(b *testing.B)           { benchmarkGATournament(50, 500, false, b) }
func BenchmarkGATournamentTerminateEarly_10(b *testing.B) { benchmarkGATournament(10, 100, true, b) }
func BenchmarkGATournamentTerminateEarly_20(b *testing.B) { benchmarkGATournament(20, 200, true, b) }
func BenchmarkGATournamentTerminateEarly_50(b *testing.B) { benchmarkGATournament(50, 500, true, b) }

func benchmarkGARoulette(length, generations int, terminateEarly bool, b *testing.B) {
	var genA = NewGeneticAlgorithm()
	genA.SetSeed(seed)
	var (
		bestCandidate Genome
		numIterations int
		candidatePool Population
	)
	for n := 0; n < b.N; n++ {
		genA.Run(length, length, generations, true, true, terminateEarly)
		b.Log("Best Candidate", genA.BestCandidate)
		b.Log("Num Iterations:", genA.Generations)
		b.Log("Population:", genA.Candidates)
	}
	globalBestCandidate = bestCandidate
	globalNumIterations = numIterations
	globalPopulation = candidatePool
}

func BenchmarkGARouletteFull_10(b *testing.B)           { benchmarkGARoulette(10, 100, false, b) }
func BenchmarkGARouletteFull_20(b *testing.B)           { benchmarkGARoulette(20, 200, false, b) }
func BenchmarkGARouletteFull_50(b *testing.B)           { benchmarkGARoulette(50, 500, false, b) }
func BenchmarkGARouletteTerminateEarly_10(b *testing.B) { benchmarkGARoulette(10, 100, true, b) }
func BenchmarkGARouletteTerminateEarly_20(b *testing.B) { benchmarkGARoulette(20, 200, true, b) }
func BenchmarkGARouletteTerminateEarly_50(b *testing.B) { benchmarkGARoulette(50, 500, true, b) }
