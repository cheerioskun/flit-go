package utils

type Tuple struct {
	Key string
	Val string
}

func MapFromTuples(list []Tuple) map[string]string {
	m := make(map[string]string)
	for _, tuple := range list {
		m[tuple.Key] = tuple.Val
	}
	return m
}
