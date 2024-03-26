package main

const (
	BreedDomesticShorthair = "Domestic Shorthair"
	BreedDomesticLonghair  = "Domestic Longhair"
)

type Animal struct {
	Name   string
	Breed  string
	Weight float64
}

func RangePointers(animals []*Animal) (int, int) {
	domesticShorthairs, domesticLonghairs := 0, 0
	for _, animal := range animals {
		switch animal.Breed {
		case BreedDomesticShorthair:
			domesticShorthairs++
		case BreedDomesticLonghair:
			domesticLonghairs++
		}
	}
	return domesticShorthairs, domesticLonghairs
}

func RangeStructs(animals []Animal) (int, int) {
	domesticShorthairs, domesticLonghairs := 0, 0
	for _, animal := range animals {
		switch animal.Breed {
		case BreedDomesticShorthair:
			domesticShorthairs++
		case BreedDomesticLonghair:
			domesticLonghairs++
		}
	}
	return domesticShorthairs, domesticLonghairs
}

func IndexPointers(animals []*Animal) (int, int) {
	domesticShorthairs, domesticLonghairs := 0, 0
	for i := range animals {
		switch animals[i].Breed {
		case BreedDomesticShorthair:
			domesticShorthairs++
		case BreedDomesticLonghair:
			domesticLonghairs++
		}
	}
	return domesticShorthairs, domesticLonghairs
}

func IndexStructs(animals []Animal) (int, int) {
	domesticShorthairs, domesticLonghairs := 0, 0
	for i := range animals {
		switch animals[i].Breed {
		case BreedDomesticShorthair:
			domesticShorthairs++
		case BreedDomesticLonghair:
			domesticLonghairs++
		}
	}
	return domesticShorthairs, domesticLonghairs
}
