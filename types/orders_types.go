package types


type OrderStore interface {
	CreateOrder(OrdersAdd) error
	GetOrders()([]Orders,error)
}

type OrdersAdd struct {
	UserId  int    `json:"user_id" validation:"required"`
	Total   float64    `json:"total" validation:"required"`
	Status  string `json:"status" validation:"requred"`
	Address string `json:"address" validation:"required"`
}

type Orders struct {
	Id        int    `json:"id" validation:"required"`
	UserId    int    `json:"user_id" validation:"required"`
	Total     float64    `json:"total" validation:"required"`
	Status    string `json:"status" validation:"requred"`
	Address   string `json:"address" validation:"required"`
	CreatedAt string `json:"created_at" validation:"required"`
}
