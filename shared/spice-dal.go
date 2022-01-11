package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type SpiceDal interface {
	ReadSpices() []Spice
	WriteSpice(Spice)
}

//SpiceDal implementation for file IO
type FsSpiceDal struct {
}

func (fsDal FsSpiceDal) ReadSpices() []Spice {
	var spices []Spice

	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			var spice Spice

			data, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
			}

			err = json.Unmarshal([]byte(data), &spice)
			if err != nil {
				fmt.Printf("Error converting from %v to spice: %v\n", data, err)
			}

			spices = append(spices, spice)
		}
	}

	return spices
}

func (fsDal FsSpiceDal) WriteSpice(s Spice) {
	file, _ := json.MarshalIndent(s, "", " ")

	err := ioutil.WriteFile(s.Name+".json", file, 0644)
	if err != nil {
		fmt.Printf("Error writing spice: %v\n", s)
	}
}
