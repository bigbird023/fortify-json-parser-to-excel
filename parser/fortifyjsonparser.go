package parser

import (
	"encoding/json"
	"io"
	"os"

	"github.com/bigbird023/fortify-json-parser-to-excel/data"
)

type (
	//FortifyJSONParserInterface interface for json parser
	FortifyJSONParserInterface interface {
		JSONParse(inputFile string) (*FortifyJSON, error)
	}

	//FortifyJSONParser struct to create an object
	FortifyJSONParser struct {
	}

	//FortifyJSON array of FortifyIssues
	FortifyJSON []data.FortifyIssue
)

//NewFortifyJSONParser create a New FortifyJSONParser
func NewFortifyJSONParser() FortifyJSONParserInterface {
	return &FortifyJSONParser{}
}

//JSONParse parsing json to struct
func (f *FortifyJSONParser) JSONParse(inputFile string) (*FortifyJSON, error) {

	// Open file
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	// we initialize our Users array
	var fortifyJSON FortifyJSON
	// we unmarshal our byteArray which contains our
	// jsonFiles content into 'users' which we defined above
	json.Unmarshal(byteValue, &fortifyJSON)

	return &fortifyJSON, nil

}
