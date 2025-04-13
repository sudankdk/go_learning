package structs

//struct is very similar to class
type Person struct{
	firstName string //private fields
	// LastName string //public fields
	lastName string
	age int
}

//class ko function banako jun cahi class of object dekhi matra invoke huna sakca
func (p Person) walk () string{
	return p.firstName + "Likes Walking"
}

//setter function
func(p *Person) SetName(name string){  //* hale paxi pass by refrence hunxa natra pass by value hunxa
	p.firstName = name
}

// another way to craete objects  
// also callled constructor
func NewPerson(fn, ln string, age int)Person {
	return Person{
		firstName: fn,
		lastName: ln,
		age: age,
	}
}