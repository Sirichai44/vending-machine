package dtos

var (
	DBProduct = "products"
	DBCoins   = "money"
)

type Coins struct {
	Value   int `json:"value"`
	Quantity int `json:"quatity"`
}

type Product struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	ImageURL  string  `json:"image_url"`
	Price     int     `json:"price"`
	Stock     int     `json:"stock"`
	CreatedAt []uint8 `json:"created_at"`
}

type DtoProduct struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	ImageUrl string `json:"image_url" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Stock    int    `json:"stock" validate:"required"`
}

type BuyProduct struct {
	Product []Products  `json:"product" validate:"required"`
	Pay     int         `json:"pay" validate:"required"`
	Total   int         `json:"total" validate:"required"`
	Value   []ValueCoin `json:"value" validate:"required"`
}

type ValueCoin struct {
	Type  int `json:"type"`
	Count int `json:"count"`
}

type Products struct {
	ID      uint `json:"id"`
	Quantity int  `json:"quantity" validate:"required"`
}
