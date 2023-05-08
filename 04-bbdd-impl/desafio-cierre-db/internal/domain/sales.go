package domain

type Sales struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
	Quantity  int `json:"quantity"`
}
