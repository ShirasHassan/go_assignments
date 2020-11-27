package mapstore

import (
	"assignment1/domain"
	"errors"
)

func (m *MapStore) Create(customer domain.Customer) error {
	if _, hasVal := m.store[customer.ID]; hasVal == true {
		return errors.New("'Customer '" + customer.ID + "' is already exist.'")
	}
	m.store[customer.ID] = customer
	return nil
}
