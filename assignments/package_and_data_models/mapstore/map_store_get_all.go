package mapstore

import "assignment1/domain"

func (m MapStore) GetAll() ([]domain.Customer, error) {
	customers := make([]domain.Customer, 0)
	for _, customer := range m.store {
		customers = append(customers, customer)
	}
	return customers, nil
}
