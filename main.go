package main

type Student struct {
	Nmae   string
	Number int
	Grade  int
}
type Teacher struct {
	Name string
}

type Person interface {
	getEmail() string
}

func main() {
}
