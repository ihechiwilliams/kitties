package kittiesbundle

// Kitty struct
type Kitty struct {
	Name      string `json:"name"`
	Breed     string `json:"breed"`
	BirthDate string `json:"birthDate"`
}

// NewKitty create a new Kitty
func NewKitty(name string, breed string, birthDate string) *Kitty {
	return &Kitty{
		Name:      name,
		Breed:     breed,
		BirthDate: birthDate,
	}
}
