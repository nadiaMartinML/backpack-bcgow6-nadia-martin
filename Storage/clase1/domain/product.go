package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	TypeProduct string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
	WarehouseId int     `json:"warehouse_id"`
}
