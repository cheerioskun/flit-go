package models

import (
	"fmt"
	"time"
)

// Model to hold details about a workout.
type Workout struct {

	// Date when the workout was done
	when time.Time
	// List of exercises in the workout
	exercises []Exercise

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
		when:      dateTime,
		exercises: workout,
	}, nil
}

func (w *Workout) Show() {
	fmt.Printf("%s", w.when.String())
	for _, exercise := range w.exercises {
		exercise.Show()
	}
}
