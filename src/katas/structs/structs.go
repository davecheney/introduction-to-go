package structs

type Pet struct {
	Name string
	Kind string
}

func MyPet() Pet {
	var p Pet
	p.Name = "Frankie"
	p.Kind = "dog"
	return p
}
