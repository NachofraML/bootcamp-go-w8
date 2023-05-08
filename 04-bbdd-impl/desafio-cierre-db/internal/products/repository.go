package products

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(product *domain.Products) (int64, error)
	ReadAll() ([]*domain.Products, error)
	GetTopProducts(limit int) ([]*domain.TopProducts, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product *domain.Products) (int64, error) {
	query := `INSERT INTO products (description, price) VALUES (?, ?)`
	row, err := r.db.Exec(query, &product.Description, &product.Price)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Products, error) {
	query := `SELECT id, description, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]*domain.Products, 0)
	for rows.Next() {
		product := domain.Products{}
		err = rows.Scan(&product.Id, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *repository) GetTopProducts(limit int) ([]*domain.TopProducts, error) {
	query := `SELECT p.description, SUM(s.quantity) as 'Total' FROM products p
    			INNER JOIN sales s ON s.product_id = p.id
				GROUP BY p.description
    			ORDER BY SUM(s.quantity) DESC
				LIMIT ?`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	topProducts := make([]*domain.TopProducts, 0)
	for rows.Next() {
		topProduct := domain.TopProducts{}
		err = rows.Scan(&topProduct.Description, &topProduct.Total)
		if err != nil {
			return nil, err
		}
		topProducts = append(topProducts, &topProduct)
	}
	return topProducts, nil
}
