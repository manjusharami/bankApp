package domain

type Customer struct {
	Id       int
	Name     string
	City     string
	Postcode string
	Status   string
}

type CustomerRepositry interface {
	FindAll() ([]Customer, error)
}

 