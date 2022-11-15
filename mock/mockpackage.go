package mock

import (
	"fmt"
	"io"
)

//Package is to mimic io.Closer and provide extra mocking details
type Package struct {
	io.Closer
	SaveCalled      bool
	SaveAsInterface interface{}
	ForceError      bool
}

//NewPackage creates a new package struct
func NewPackage() *Package {
	return &Package{}
}

//Save will mock the save method
func (m *Package) Save() error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveCalled = true
	return nil
}

//SaveAs will mock the saveas method
func (m *Package) SaveAs(target interface{}) error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveAsInterface = target
	return nil
}
