package models

import (
	"errors"
	"fmt"
	"strings"
)

type Exercise struct {
	Name     string            `json:"name"`
	Settings map[string]string `json:"settings"`
	Sets     []Set             `json:"sets"`
}

type Set struct {
	Reps    RepsPair   `json:"reps"`
	Weights WeightPair `json:"weight"`
	Comment string     `json:"comment"`
}

type RepsPair struct {
	Left  int64 `json:"left"`
	Right int64 `json:"right"`
}

type WeightPair struct {
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

func NewSet(reps RepsPair, weight interface{}, comment string) (Set, error) {
	if weight == nil {
		weight = WeightPair{0, 0}
	}
	return Set{
		Reps:    reps,
		Weights: weight.(WeightPair),
		Comment: comment,
	}, nil
}

func NewExercise(name string, sets []Set, settings interface{}) (Exercise, error) {

	if name == "" {
		return Exercise{}, errors.New("name must not be empty")
	}
	if len(sets) == 0 {
		return Exercise{}, errors.New("sets cannot be empty")
	}
	name = strings.TrimSpace(name)
	ex := Exercise{name, nil, sets}
	if settings != nil {
		ex.Settings = settings.(map[string]string)
	}
	return ex, nil
}

func (e *Exercise) Show() {
	fmt.Printf("Name: %s, Vals: %+v", e.Name, e.Sets)

}
