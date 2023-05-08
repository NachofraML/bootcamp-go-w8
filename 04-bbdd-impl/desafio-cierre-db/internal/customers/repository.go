package customers

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(customers *domain.Customers) (int64, error)
	GetConditionsTotals() ([]*domain.CustomerConditionsTotals, error)
	ReadAll() ([]*domain.Customers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers *domain.Customers) (int64, error) {
	query := `INSERT INTO customers (first_name, last_name, customers.condition) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *repository) GetConditionsTotals() ([]*domain.CustomerConditionsTotals, error) {
	query := `SELECT c.condition as 'Condition', ROUND(SUM(p.price * s.quantity), 2) as 'Total' FROM customers c
					INNER JOIN invoices i ON i.customer_id = c.id
					INNER JOIN sales s ON s.invoice_id = i.id 
					INNER JOIN products p ON p.id = s.product_id 
					GROUP BY c.condition`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customerConditionsTotals := make([]*domain.CustomerConditionsTotals, 0)
	for rows.Next() {
		customerConditionTotal := domain.CustomerConditionsTotals{}
		err = rows.Scan(&customerConditionTotal.Condition, &customerConditionTotal.Total)
		if err != nil {
			return nil, err
		}
		customerConditionsTotals = append(customerConditionsTotals, &customerConditionTotal)
	}
	return customerConditionsTotals, err
}
