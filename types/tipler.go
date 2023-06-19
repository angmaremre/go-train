package types

type Marka struct {
	Id        int
	Name      string
	Is_active bool
}

type Markalar []Marka

func (currentMarka *Marka) UpdateMarkaName(newName string) {
	currentMarka.Name = newName
}
