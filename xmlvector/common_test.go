package xmlvector

import (
	"fmt"
	"os"
)

var (
	// small, medium and large fixtures are from https://www.w3schools.com/xml/xml_examples.asp
	smallFixture  = getFromFile("testdata/small.xml")
	mediumFixture = getFromFile("testdata/medium.xml")
	largeFixture  = getFromFile("testdata/large.xml")

	// sigmod and mondial fixtures are from https://aiweb.cs.washington.edu/research/projects/xmltk/xmldata/
	sigmodFixture  = getFromFile("testdata/sigmod.xml")
	mondialFixture = getFromFile("testdata/mondial.xml")
)

func getFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("cannot read %s: %s", filename, err))
	}
	return string(data)
}
