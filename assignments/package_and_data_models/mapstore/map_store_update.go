package mapstore

import (
	"assignment1/domain"
	"errors"
)

func (m *MapStore) Update(id string, customer domain.Customer) error {
	if _, hasVal := m.store[id]; hasVal == false {
		return errors.New("'Customer '" + id + "' is not exist.'")
	}
	customer.ID = id
	m.store[id] = customer
	return nil
}
