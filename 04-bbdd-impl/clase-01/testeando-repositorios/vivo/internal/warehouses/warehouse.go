package warehouses

type Warehouse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

type ProductReport struct {
	WarehouseName string `json:"warehouse_name"`
	ProductCount  int    `json:"product_count"`
}
