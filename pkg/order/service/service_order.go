package repository

import (
	"errors"
	"go_services_lab/models"
	"go_services_lab/pkg/order/repository"
	"strconv"
)

type OrderService struct {
	rep repository.Order
}

func NewOrderService(rep repository.Order) *OrderService {
	return &OrderService{rep: rep}
}

func atoiMap(products map[string]int) (map[int]int, error) {
	ret_map := make((map[int]int))
	for key, val := range products {
		k, err := strconv.Atoi(key)
		if err != nil {
			return ret_map, errors.New("Wronge key for product's ID.")
		} else {
			ret_map[k] = val
		}
	}
	return ret_map, nil
}

func (s *OrderService) Get(id int) (models.Order, error) {
	return s.rep.Get(id)
}

func (s *OrderService) Amount(id int) (float32, error) {
	return s.rep.Amount(id)
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.rep.GetAll()
}

func (s *OrderService) Delete(id int) (int, error) {
	return s.rep.Delete(id)
}

func (s *OrderService) Create(user_id int, products map[string]int) (int, error) {
	pr, err := atoiMap(products)
	if err != nil {
		return 0, err
	}
	return s.rep.Create(user_id, pr)
}
