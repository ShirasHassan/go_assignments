package mapstore

import (
	"assignment1/domain"
	"errors"
)

func (m *MapStore) GetById(id string) (domain.Customer, error) {
	if _, hasVal := m.store[id]; hasVal == false {
		return domain.Customer{}, errors.New("'Customer '" + id + "' is not exist.'")
	}
	return m.store[id], nil
}
