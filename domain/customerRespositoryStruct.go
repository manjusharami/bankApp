package domain

type CustomerRepositryStruct struct {
	customer []Customer
}

func (s CustomerRepositryStruct) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRespository() CustomerRepositryStruct {
	customer := []Customer{
		{
			Id:       90,
			Name:     "manju",
			City:     "chennai",
			Postcode: "mk9988",
			Status:   "active",
		},
		{
			Id:       190,
			Name:     "sweet",
			City:     "chennai",
			Postcode: "mk988",
			Status:   "active",
		},
	}
	return CustomerRepositryStruct{customer}
}
