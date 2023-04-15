package models

import "github.com/cheerioskun/flit-go/models/user"

type FitnessLog struct {
	// A list of workouts
	entries []*Workout

	// Details about user
	user user.UserDetails
}

func NewFitnessLog(workouts []*Workout) *FitnessLog {
	return &FitnessLog{
		entries: workouts,
	}
}
func (l *FitnessLog) AddWorkout(w Workout) {
	l.entries = append(l.entries, &w)
}

func (l *FitnessLog) Show() {
	for _, workout := range l.entries {
		workout.Show()
	}
}
