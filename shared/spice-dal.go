package shared

import "fmt"

type SpiceDal interface {
	ReadSpices() []Spice
	WriteSpice(Spice)
}

type FsSpiceDal struct {
}

func (fsDal FsSpiceDal) ReadSpices() []Spice {
	//seed data
	var spices = []Spice{
		{Id: 1, Name: "Oregano", Level: 50, Substitute: "Rosemary"},
		{Id: 2, Name: "Rosemary", Level: 25, Substitute: "Oregano"},
		{Id: 3, Name: "Sage", Level: 75, Substitute: "Thyme"},
		{Id: 4, Name: "Thyme", Level: 125, Substitute: "Sage"},
	}

	return spices
}

func (fsDal FsSpiceDal) WriteSpice(s Spice) {
	fmt.Println(s)
}
