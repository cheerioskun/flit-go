package models

import "time"

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
