package models

import "github.com/cheerioskun/flit-go/models/user"

type FitnessLog struct {
	// A list of workouts
	entries []*Workout

	// Details about user
	user user.UserDetails
}
