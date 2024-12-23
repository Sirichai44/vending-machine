package services

import (
	"database/sql"

	"vending_machine/dtos"
)

type CoinService interface {
	Service[dtos.Coins]
	UpdateCoinQuantity(value int, quantity int) error
}

type coinService struct {
	Service[dtos.Coins]
	mysqlDB *sql.DB
}

func NewCoinService(mySqlConn *sql.DB) CoinService {
	return &coinService{
		Service: NewService[dtos.Coins](mySqlConn, dtos.DBCoins),
		mysqlDB: mySqlConn,
	}
}

func (c *coinService) UpdateCoinQuantity(value int, quantity int) error {
	_, err := c.mysqlDB.Exec("UPDATE money SET quantity = quantity + ? WHERE value = ?", quantity, value)
	return err
}
