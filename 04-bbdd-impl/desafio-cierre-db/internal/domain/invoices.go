package domain

type Invoices struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	CustomerId int     `json:"customer_id"`
	Total      float64 `json:"total"`
}

type InvoiceProducts struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
