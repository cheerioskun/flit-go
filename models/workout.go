package models

import (
	"fmt"
	"time"
)

// Model to hold details about a workout.
type Workout struct {

	// Date When the workout was done
	When time.Time `json:"date"`
	// List of Exercises in the workout
	Exercises []Exercise `json:"exercises"`

	// // Images like progress and full-body scans
	// images []Image

	// Computed fields
	// Amount of calories burnt
	calories int
}

func NewWorkout(date string, workout []Exercise) (*Workout, error) {
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return &Workout{
		When:      dateTime,
		Exercises: workout,
	}, nil
}

func (w *Workout) Show() {
	fmt.Printf("%s", w.When.String())
	for _, exercise := range w.Exercises {
		exercise.Show()
	}
}
