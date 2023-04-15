package models

import (
	"errors"
	"fmt"
)

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
	Left  int64
	Right int64
}

type WeightPair struct {
	Left  float64
	Right float64
}

func NewSet(reps RepsPair, weight interface{}, comment string) (Set, error) {
	if weight == nil {
		weight = WeightPair{0, 0}
	}
	return Set{
		reps:    reps,
		weight:  weight.(WeightPair),
		comment: comment,
	}, nil
}

func NewExercise(name string, sets []Set, settings interface{}) (Exercise, error) {
	if settings == nil {
		settings = make(map[string]string)
	}
	if name == "" {
		return Exercise{}, errors.New("name must not be empty")
	}
	if len(sets) == 0 {
		return Exercise{}, errors.New("sets cannot be empty")
	}
	return Exercise{name, settings.(map[string]string), sets}, nil
}

func (e *Exercise) Show() {
	fmt.Printf("Name: %s, Vals: %+v", e.name, e.sets)

}
