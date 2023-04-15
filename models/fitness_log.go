package models

import "github.com/cheerioskun/flit-go/models/user"

type FitnessLog struct {
	// A list of workouts
	Entries []*Workout `json:"entries"`

	// Details about user
	User user.UserDetails `json:"user"`
}

func NewFitnessLog(workouts []*Workout) *FitnessLog {
	return &FitnessLog{
		Entries: workouts,
	}
}
func (l *FitnessLog) AddWorkout(w Workout) {
	l.Entries = append(l.Entries, &w)
}

func (l *FitnessLog) Show() {
	for _, workout := range l.Entries {
		workout.Show()
	}
}
