package language

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cheerioskun/flit-go/language/lexer"
	"github.com/cheerioskun/flit-go/language/parser"
	"github.com/cheerioskun/flit-go/models"
)

type T1 struct {
	src    string
	expect *models.FitnessLog
}

var src = `
2023-04-14
1. Pushups = 13+12+10+4
2. Shoulder Shrugs = 12([12 12])+14([12 12])+12([14 14])
3. Butterfly(setting:4) = 12(30)+12(30)+8(30)+4(23)
4. Bench Press(grip:wide) = 12(+30)[minor shoulder pain]+12(+30)
5. Forearm Curls = [10 12]([5 5])+10(5)
`
var testData = T1{
	src:    src,
	expect: &models.FitnessLog{},
}

func TestParsing(t *testing.T) {
	p := parser.NewParser()
	s := lexer.NewLexer([]byte(testData.src))
	// for i := 0; i < 15; i++ {
	// 	t.Logf("%s", string(s.Scan().Lit))
	// }
	fl, err := p.Parse(s)
	if err != nil {
		t.Fatalf("FAILED: %q", err)
	} else {
		bytes, _ := json.MarshalIndent(fl.(*models.FitnessLog), "", "    ")
		fmt.Printf("%s", string(bytes))
		t.Fatalf("%+v", fl)
	}
}

// 2. Shoulder Shrugs = 12([12 12])+14([12 12])+12([14 14])
// 3. Butterfly(setting:4) = 12(30)+12(30)+8(30)+4(23)
// 4. Bench Press(grip:wide) = 12(+30)[minor shoulder pain]+12(+30)
// 5. Forearm Curls = [10 12]([5 5])+10(5)
