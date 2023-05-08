package products

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound      = errors.New("product not found in the given repository")
	ErrAlreadyExists = errors.New("product already exists in the given repository")
)

type Repository interface {
	Get(id int) (product *Product, err error)
	Create(product *Product) (err error)
	Update(product *Product) (err error)
	Delete(id int) (err error)
}

func NewMySqlRepositoty(db *sql.DB) Repository {
	return &MySqlRepositoty{db: db}
}

type MySqlRepositoty struct {
	db *sql.DB
}

func (m *MySqlRepositoty) Get(id int) (product *Product, err error) {
	product = &Product{}
	query := `SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?`
	row := m.db.QueryRow(query, id)

	err = row.Scan(&product.ID, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return
	}
	return
}

func (m *MySqlRepositoty) Create(product *Product) (err error) {
	stmt, err := m.db.Prepare(`INSERT INTO products VALUES(null, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		// Cast to MySQL error.
		mysqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return
		}

		// Check the error code.
		switch mysqlError.Number {
		case 1062:
			err = ErrAlreadyExists
		case 1586:
			err = ErrAlreadyExists
			// TODO: Handle more errors.
		}

		return
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}
	product.ID = int(lastId)
	return
}

func (m *MySqlRepositoty) Update(product *Product) (err error) {
	stmt, err := m.db.Prepare(`UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.ID)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}
	return
}

func (m *MySqlRepositoty) Delete(id int) (err error) {
	stmt, err := m.db.Prepare(`DELETE FROM products WHERE id = ?`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}
	return
}
