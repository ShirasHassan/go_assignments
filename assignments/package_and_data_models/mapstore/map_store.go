package mapstore

import (
	"assignment1/domain"
	"errors"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() domain.CustomerStore {
	return &MapStore{
		store: make(map[string]domain.Customer)}
}

func (m *MapStore) Create(customer domain.Customer) error {
	if _, hasVal := m.store[customer.ID]; hasVal == true {
		return errors.New("'Customer '" + customer.ID + "' is already exist.'")
	}
	m.store[customer.ID] = customer
	return nil
}

func (m *MapStore) Update(id string, customer domain.Customer) error {
	if _, hasVal := m.store[id]; hasVal == false {
		return errors.New("'Customer '" + id + "' is not exist.'")
	}
	customer.ID = id
	m.store[id] = customer
	return nil
}

func (m *MapStore) Delete(id string) error {
	if _, hasVal := m.store[id]; hasVal == false {
		return errors.New("'Customer '" + id + "' is not exist.'")
	}
	delete(m.store, id)
	return nil
}

func (m MapStore) GetAll() ([]domain.Customer, error) {
	customers := make([]domain.Customer, 0)
	for _, customer := range m.store {
		customers = append(customers, customer)
	}
	return customers, nil
}


func (m *MapStore) GetById(id string) (domain.Customer, error) {
	if _, hasVal := m.store[id]; hasVal == false {
		return domain.Customer{}, errors.New("'Customer '" + id + "' is not exist.'")
	}
	return m.store[id], nil
}
