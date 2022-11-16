package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	// carts.id, carts.product_id, products.name, carts.quantity, carts.total_price
	var carts []model.JoinCart
	err := c.db.Table("carts").Select("carts.id, carts.product_id, products.name, carts.quantity, carts.total_price").Joins("JOIN products ON carts.product_id = products.id").Scan(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c *CartRepository) AddCart(product model.Product) error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var cart model.Cart

	// check if product exists in the carts
	err := tx.Where("product_id = ?", product.ID).First(&cart).Error
	if err != nil {
		// so the product doesn't exist in the carts
		disc := product.Price * product.Discount / 100
		cart.ProductID = product.ID
		cart.Quantity = 1
		cart.TotalPrice = product.Price - disc
		err = tx.Create(&cart).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		// update the product stock
		err = tx.Model(&model.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()
		return nil
	}

	disc := product.Price * product.Discount / 100

	cart.Quantity += 1
	cart.TotalPrice += product.Price - disc
	err = tx.Model(&model.Cart{}).Where("id = ?", cart.ID).Updates(cart).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// and update product stock
	err = tx.Model(&model.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock-1).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	// delete cart and add the stock back
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var cart model.Cart
	err := tx.Where("id = ?", id).First(&cart).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var product model.Product
	err = tx.Where("id = ?", productID).First(&product).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&cart).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&model.Product{}).Where("id = ?", productID).Update("stock", product.Stock+int(cart.Quantity)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	return c.db.Model(&model.Cart{}).Where("id = ?", id).Updates(cart).Error
}
