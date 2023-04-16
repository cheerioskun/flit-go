package language

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/cheerioskun/flit-go/language/lexer"
	"github.com/cheerioskun/flit-go/language/parser"
	"github.com/cheerioskun/flit-go/models"
	"github.com/nsf/jsondiff"
)

type T1 struct {
	name   string
	src    string
	expect *models.FitnessLog
}

var src = `
2023-04-14
1. Pushups(grip:wide) = [13 10]+12[minor shoulder pain]
2. Shoulder Shrugs = 12([12 -12])+14([12 +10])
`

var dateToUse, _ = time.Parse("2006-01-02", "2023-04-14")
var testData = []T1{
	{
		name: "Single Entry, Multiple Exercises, Separate reps, Separate weights, comments, settings, negative weights",
		src:  src,
		expect: &models.FitnessLog{
			Entries: []*models.Workout{
				{
					When: dateToUse,
					Exercises: []models.Exercise{
						{
							Name: "Pushups",
							Sets: []models.Set{
								{
									Reps:    models.RepsPair{13, 10},
									Weights: models.WeightPair{0, 0},
								},
								{
									Reps:    models.RepsPair{12, 12},
									Weights: models.WeightPair{0, 0},
									Comment: "minor shoulder pain",
								},
							},
							Settings: map[string]string{
								"grip": "wide",
							},
						},
						{
							Name: "Shoulder Shrugs",
							Sets: []models.Set{
								{
									Reps:    models.RepsPair{12, 12},
									Weights: models.WeightPair{12, -12},
								},
								{
									Reps:    models.RepsPair{14, 14},
									Weights: models.WeightPair{12, 10},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestParsing(t *testing.T) {
	p := parser.NewParser()
	for _, test := range testData {
		s := lexer.NewLexer([]byte(test.src))
		// for i := 0; i < 15; i++ {
		// 	t.Logf("%s", string(s.Scan().Lit))
		// }
		fl, err := p.Parse(s)
		if err != nil {
			t.Logf("FAILED: %q", err)
			t.Fail()
		}

		if !reflect.DeepEqual(fl, test.expect) {
			jma, _ := json.MarshalIndent(fl, "", "  ")
			jmb, _ := json.MarshalIndent(test.expect, "", "  ")
			opts := jsondiff.DefaultJSONOptions()
			_, explanation := jsondiff.Compare(jma, jmb, &opts)
			t.Logf("%s", explanation)
			t.Logf("t: %s\nNot deeply equal", test.name)
			t.Fail()
		}
	}

}
