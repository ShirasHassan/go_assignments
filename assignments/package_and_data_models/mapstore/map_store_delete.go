package mapstore

import "errors"

func (m *MapStore) Delete(id string) error {
	if _, hasVal := m.store[id]; hasVal == false {
		return errors.New("'Customer '" + id + "' is not exist.'")
	}
	delete(m.store, id)
	return nil
}
