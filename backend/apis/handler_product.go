package apis

import (
	"context"
	"fmt"
	"log"

	"vending_machine/dtos"
	"vending_machine/services"

	"github.com/gofiber/fiber/v2"
)

type product struct {
	srvProduct services.ProductService
	srvMoney   services.CoinService
}

func NewHandlerProduct(f fiber.Router, srvProduct services.ProductService, srvMoney services.CoinService) {
	product := product{srvProduct: srvProduct, srvMoney: srvMoney}

	g := f.Group(apiVersion).Group("product")

	g.Get("", HandleResponse(product.GetAllList))
	g.Post("buy", HandleBodyParser(product.BuyProduct))
}

func (p *product) GetAllList(c *fiber.Ctx) (*dtos.Context, error) {
	products, err := p.srvProduct.FindAll(c.Context())
	if err != nil {
		return NewContext(fiber.StatusInternalServerError, "Failed to get all product", nil), err
	}

	result := []dtos.DtoProduct{}

	for _, product := range products {
		result = append(result, dtos.DtoProduct{
			ID:       int(product.ID),
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			ImageUrl: product.ImageURL,
		})
	}

	log.Println("Product", "GetProduct", "successfully")
	return NewContext(fiber.StatusOK, "Get all product successfully", result), nil
}

func (p *product) BuyProduct(dto dtos.BuyProduct, c *fiber.Ctx) (*dtos.Context, error) {
	change, err := p.calculateChange(c.Context(), dto.Pay, dto.Total)
	if err != nil {
		return NewContext(fiber.StatusInternalServerError, "Error calculating change", nil), err
	}

	err = p.updateMoneyInDB(c.Context(), dto.Pay)
	if err != nil {
		return NewContext(fiber.StatusInternalServerError, "Error updating money in DB", nil), err
	}

	for _, product := range dto.Product {
		err = p.srvProduct.UpdateProduct(int(product.ID), product.Quantity)
		if err != nil {
			return NewContext(fiber.StatusInternalServerError, "Error updating product in DB", nil), err
		}
	}

	log.Println("Product", "BuyProduct", "successfully")
	return NewContext(fiber.StatusOK, "Buy product successfully", change), nil
}

func (p *product) calculateChange(c context.Context, pay, total int) (map[int]int, error) {
	money, err := p.srvMoney.FindAll(c)
	if err != nil {
		return nil, err
	}

	// sort money
	for i := 0; i < len(money); i++ {
		for j := i + 1; j < len(money); j++ {
			if money[i].Value < money[j].Value {
				money[i], money[j] = money[j], money[i]
			}
		}
	}

	change := pay - total
	changeMap := make(map[int]int)

	for _, m := range money {
		if change <= 0 {
			break
		}
		if m.Quantity > 0 && change >= m.Value {
			numCoins := change / m.Value
			if numCoins > m.Quantity {
				numCoins = m.Quantity
			}
			change -= numCoins * m.Value
			changeMap[m.Value] = numCoins
		}
	}

	if change > 0 {
		return nil, fmt.Errorf("not enough coins to give change")
	}

	return changeMap, nil
}

func (p *product) updateMoneyInDB(c context.Context, pay int) error {
	money, err := p.srvMoney.FindAll(c)
	if err != nil {
		return err
	}

	for _, m := range money {
		if pay <= 0 {
			break
		}
		if pay >= m.Value {
			numCoins := pay / m.Value
			pay -= numCoins * m.Value
			err := p.srvMoney.UpdateCoinQuantity(m.Value, numCoins)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
