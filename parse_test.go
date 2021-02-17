package parsing

import (
	"testing"

	parser "github.com/LaurentColoma/M-BDX-002/parsing"
)

func TestParsingHandler(t *testing.T) {

	emptyResult := parser.ParsingHandler("colis_a_livrer 2 1 green", 5, 0)

	if emptyResult != true {
		t.Errorf("ParsingHandler(\"colis_a_livrer 2 1 green\") failed, expected true, got false")
	} else {
		t.Logf("the program started properly")
	}
}
