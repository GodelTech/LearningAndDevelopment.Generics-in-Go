// package main contains an example of interfaces usage in code with generics
package main

import "fmt"

// Stringer is a type which implements method String() string
type Stringer interface {
	String() string
}

// Person describes a person with Firstname and Lastname props
type Person struct {
	Firstname, Lastname string
}

// String is needed to implement a Stringer interface
func (p *Person) String() string {
	return p.Firstname + " " + p.Lastname
}

// Animal describes an animal with Name and Alias properties
type Animal struct {
	Name, Alias string
}

// String is needed to implement a Stringer interface
func (a *Animal) String() string {
	return fmt.Sprintf("%s(%s)", a.Alias, a.Name)
}

func concat[T Stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String() + ","
	}
	return result
}

func main() {
	peoples := []*Person{
		{
			Firstname: "Aliaksandr",
			Lastname:  "Paliavy",
		},
		{
			Firstname: "John",
			Lastname:  "Doe",
		},
	}

	animals := []*Animal {
		{
			Name:  "Cat",
			Alias: "Lily",
		},
		{
			Name:  "Dog",
			Alias: "Bobby",
		},
	}

	fmt.Printf("peoples: %s\n", concat(peoples))
	fmt.Printf("animals: %s\n", concat(animals))
}