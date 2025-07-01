package models

type Person struct {
	name    string
	Age     int
	Birth   int
	Address string
	Phone   string
}

func (p Person) String() string {
	return "Name: " + p.name + ", Age: " + string(p.Age) + ", Birth: " + string(p.Birth) +
		", Address: " + p.Address + ", Phone: " + p.Phone
}

// ChangeName is a method that changes the name of the person.
// It takes a pointer receiver so that it can modify the original Person instance.
// This is useful when you want to change the state of the object.
func (p *Person) ChangeName(newName string) {
	p.name = newName
}

// create constructor for Person
func NewPerson(name string, age int, birth int, address string, phone string) *Person {
	return &Person{
		name:    name,
		Age:     age,
		Birth:   birth,
		Address: address,
		Phone:   phone,
	}
}

func (p *Person) GetName() string {
	return p.name
}
