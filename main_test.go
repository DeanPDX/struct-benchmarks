package main

import (
	"fmt"
	"testing"
)

// GetAnimals returns contrived animal data for benchmarking
func GetAnimals() []Animal {
	count := 1000
	animals := make([]Animal, count)
	for i := 0; i < count; i++ {
		animals[i] = Animal{
			Name:   fmt.Sprintf("Animal #%v", i),
			Breed:  GetBreed(i),
			Weight: float64(i),
		}
	}
	return animals
}

// GetBreed is just a deterministic way to get a breed based on
// whether the current iterator value is even.
func GetBreed(ctr int) string {
	if ctr%2 == 0 {
		return BreedDomesticLonghair
	}
	return BreedDomesticShorthair
}

// GetAnimalPointers calls GetAnimals then returns a slice with
// pointers to the animal structs.
func GetAnimalPointers() []*Animal {
	animals := GetAnimals()
	animalPtrs := make([]*Animal, len(animals))
	for i := range animals {
		animalPtrs[i] = &animals[i]
	}
	return animalPtrs
}

var (
	finalResultShort int
	finalResultLong  int
)

func BenchmarkRangePointers(b *testing.B) {
	animals := GetAnimalPointers()
	// Don't incude our setup time in this benchmark.
	b.ResetTimer()
	domesticShort, domesticLong := 0, 0
	for i := 0; i < b.N; i++ {
		domesticShort, domesticLong = RangePointers(animals)
	}
	// Store final result just to make sure compiler doesn't optimize.
	finalResultShort = domesticShort
	finalResultLong = domesticLong
}

func BenchmarkRangeStructs(b *testing.B) {
	animals := GetAnimals()
	// Don't incude our setup time in this benchmark.
	b.ResetTimer()
	domesticShort, domesticLong := 0, 0
	for i := 0; i < b.N; i++ {
		domesticShort, domesticLong = RangeStructs(animals)
	}
	// Store final result just to make sure compiler doesn't optimize.
	finalResultShort = domesticShort
	finalResultLong = domesticLong
}

func BenchmarkIndexPointers(b *testing.B) {
	animals := GetAnimalPointers()
	// Don't incude our setup time in this benchmark.
	b.ResetTimer()
	domesticShort, domesticLong := 0, 0
	for i := 0; i < b.N; i++ {
		domesticShort, domesticLong = IndexPointers(animals)
	}
	// Store final result just to make sure compiler doesn't optimize.
	finalResultShort = domesticShort
	finalResultLong = domesticLong
}

func BenchmarkIndexStructs(b *testing.B) {
	animals := GetAnimals()
	// Don't incude our setup time in this benchmark.
	b.ResetTimer()
	domesticShort, domesticLong := 0, 0
	for i := 0; i < b.N; i++ {
		domesticShort, domesticLong = IndexStructs(animals)
	}
	// Store final result just to make sure compiler doesn't optimize.
	finalResultShort = domesticShort
	finalResultLong = domesticLong
}
