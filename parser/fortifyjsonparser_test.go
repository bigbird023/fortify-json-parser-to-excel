package parser

import (
	"testing"
)

func TestJsonParseBadFile(t *testing.T) {

	cfgInput := "../local/examples/testBadFile.json"

	fxp := NewFortifyJSONParser()

	_, err := fxp.JSONParse(cfgInput)
	if err != nil {
		if err.Error() != "open ../local/examples/testBadFile.json: no such file or directory" {
			t.Log(err)
			t.Fail()
		}
	} else {
		t.Log("Error expected")
		t.Fail()
	}

}

func TestJsonParseEmpty(t *testing.T) {

	cfgInput := "../local/examples/testEmpty.json"
	fxp := NewFortifyJSONParser()

	fortifyJSON, err := fxp.JSONParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if fortifyJSON == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if fortifyJSON != nil && len(*fortifyJSON) != 0 {
		t.Log("FortifyJson length should be 0")
		t.Fail()
	}

}

func TestJsonParse(t *testing.T) {

	cfgInput := "../local/examples/test.json"
	fxp := NewFortifyJSONParser()

	fortifyJSON, err := fxp.JSONParse(cfgInput)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if fortifyJSON == nil {
		t.Log("Return should not be nil")
		t.Fail()
	}

	if fortifyJSON != nil && len(*fortifyJSON) == 0 {
		t.Log("fortifyjson length should not be 0")
		t.Fail()
	}
}
