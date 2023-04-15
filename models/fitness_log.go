package models

import "github.com/cheerioskun/flit-go/models/user"

type FitnessLog struct {
	// A list of workouts
	entries []*Workout

	// Details about user
	user user.UserDetails
}

func NewFitnessLog() *FitnessLog {
	return &FitnessLog{
		entries: make([]*Workout, 0),
	}
}
func (l *FitnessLog) AddWorkout(w Workout) {
	l.entries = append(l.entries, &w)
}
