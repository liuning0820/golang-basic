package models

// Student is a student object
type Student struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func (f Student) FullName(firstName, lastName string) string {
	var result = firstName + lastName
	return result
}
