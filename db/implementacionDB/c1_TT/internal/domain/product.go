package domain

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Count           int     `json:"count"`
	Price           float64 `json:"price"`
	Warehouse_id    int     `json:"id_warehouse"`
	Warehouse       string  `json:"warehouse_name"`
	WarehouseAdress string  `json:"warehouse_adress"`
}
