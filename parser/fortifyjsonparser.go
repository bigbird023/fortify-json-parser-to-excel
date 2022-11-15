/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bigbird023/fortify-json-parser-to-excel/data"
)

type (
	FortifyJsonParserInterface interface {
		JsonParse(inputFile string) (*FortifyJson, error)
	}

	FortifyJsonParser struct {
	}

	FortifyJson []data.FortifyIssue
)

func NewFortifyJsonParser() FortifyJsonParserInterface {
	return &FortifyJsonParser{}
}

func (f *FortifyJsonParser) JsonParse(inputFile string) (*FortifyJson, error) {

	// Open file
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(file)

	// we initialize our Users array
	var fortifyJson FortifyJson
	// we unmarshal our byteArray which contains our
	// jsonFiles content into 'users' which we defined above
	json.Unmarshal(byteValue, &fortifyJson)

	return &fortifyJson, nil

}
