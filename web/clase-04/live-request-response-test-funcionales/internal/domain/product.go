package domain

import (
	"errors"
	"time"
)

type Producto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (producto *Producto) Validate() (err error) {
	err = validExpirationDate(producto.Expiration)
	if err != nil {
		return
	}
	return
}

func validExpirationDate(expirationDate string) error {
	const YYYYMMDD string = "02/01/2006"
	_, err := time.Parse(YYYYMMDD, expirationDate)
	if err != nil {
		return errors.New("invalid expiration date format")
	}
	return nil
}
