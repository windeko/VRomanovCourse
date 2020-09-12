package person

func NewPerson(id int, name string, secret string) *Person {
	return &Person{
		ID:     id,
		Name:   name,
		secret: secret,
	}
}
