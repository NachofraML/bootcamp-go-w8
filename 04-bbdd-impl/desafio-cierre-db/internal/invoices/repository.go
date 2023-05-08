package invoices

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	Update(invoices *domain.Invoices) error
	GetInvoiceProducts(id int) ([]*domain.InvoiceProducts, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) Update(invoices *domain.Invoices) error {
	query := `UPDATE invoices SET datetime = ?, customer_id = ?, total = ? WHERE id = ?`
	_, err := r.db.Exec(query, &invoices.Datetime, &invoices.CustomerId, &invoices.Total, &invoices.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetInvoiceProducts(id int) ([]*domain.InvoiceProducts, error) {
	query := `SELECT p.price, s.quantity FROM invoices i 
			  INNER JOIN sales s ON s.invoice_id = i.id 
			  INNER JOIN products p ON p.id = s.product_id 
			  WHERE i.id = ?`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoiceProducts := make([]*domain.InvoiceProducts, 0)

	for rows.Next() {
		invoiceProduct := domain.InvoiceProducts{}
		err = rows.Scan(&invoiceProduct.Price, &invoiceProduct.Quantity)
		if err != nil {
			return nil, err
		}
		invoiceProducts = append(invoiceProducts, &invoiceProduct)
	}
	return invoiceProducts, nil
}
