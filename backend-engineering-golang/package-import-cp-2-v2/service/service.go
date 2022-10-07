package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, err := s.database.GetProductByname(productName)
	if err != nil {
		return err
	}

	if quantity < 1 {
		return errors.New("invalid quantity")
	}

	cartItems, err := s.database.Load()
	if err != nil {
		return err
	}

	cartItems = append(cartItems, entity.CartItem{
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
	})

	err = s.database.Save(cartItems)
	return err
}

func (s *Service) RemoveCart(productName string) error {
	product, err := s.database.GetProductByname(productName)
	if err != nil {
		return err
	}

	cartItems, err := s.database.Load()
	if err != nil {
		return err
	}

	found := false
	for _, v := range cartItems {
		if v.ProductName == product.Name {
			found = true
		}
	}

	if !found {
		return errors.New("product not found")
	}

	for i, cartItem := range cartItems {
		if cartItem.ProductName == product.Name {
			cartItems = append(cartItems[:i], cartItems[i+1:]...)
		}
	}

	err = s.database.Save(cartItems)
	return err
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	err := s.database.Save([]entity.CartItem{})
	return err
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	products := s.database.GetProductData()
	return products, nil
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {
	cartItems, err := s.database.Load()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	var totalPrice int
	for _, cartItem := range cartItems {
		totalPrice += cartItem.Price * cartItem.Quantity
	}

	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	paymentInformation := entity.PaymentInformation{
		ListProduct: cartItems,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      money - totalPrice,
	}

	err = s.database.Save([]entity.CartItem{})
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	return paymentInformation, nil
}
