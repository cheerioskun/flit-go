package app

import (
	"encoding/json"

	"github.com/cheerioskun/flit-go/language/lexer"
	"github.com/cheerioskun/flit-go/language/parser"
	"github.com/cheerioskun/flit-go/models"
)

func ConvertFileToFmt(inpath string, format string) []byte {

	s, err := lexer.NewLexerFile(inpath)
	if err != nil {
		panic(err)
	}
	parsedStruct, err := parser.NewParser().Parse(s)
	if err != nil {
		panic(err)
	}

	fl := parsedStruct.(*models.FitnessLog)

	switch format {
	case "json":
		bytes, err := json.MarshalIndent(fl, "", "    ")
		if err != nil {
			panic(err)
		}
		return bytes
	default:
		panic("Unsupported output format " + format)
	}
}
