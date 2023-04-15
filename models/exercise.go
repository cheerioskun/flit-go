package models

type Exercise struct {
	name     string
	settings map[string]string
	sets     []Set
}

// Each set of an exercise
type Set struct {
	reps    RepsPair
	weight  WeightPair
	comment string
}

type RepsPair struct {
	left  float32
	right float32
}

type WeightPair struct {
	left  float32
	right float32
}
