package main

import (
	"assignment1/domain"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (controller *CustomerController) Add(customer domain.Customer) error {
	return controller.store.Create(customer)
}

func (controller *CustomerController) Modify(customer domain.Customer) error {
	return controller.store.Update(customer.ID, customer)
}

func (controller *CustomerController) Remove(customer domain.Customer) error {
	return controller.store.Delete(customer.ID)
}

func (controller *CustomerController) Get(id string) (domain.Customer, error) {
	return controller.store.GetById(id)
}

func (controller *CustomerController) GetAll() ([]domain.Customer, error) {
	return controller.store.GetAll()
}
