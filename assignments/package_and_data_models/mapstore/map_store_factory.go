package mapstore

import "assignment1/domain"

func NewMapStore() domain.CustomerStore {
	return &MapStore{
		store: make(map[string]domain.Customer)}
}
