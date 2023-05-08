package warehouses

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound      = errors.New("warehouse not found in the given repository")
	ErrAlreadyExists = errors.New("warehouse already exists in the given repository")
)

type Repository interface {
	Get(id int) (warehouse *Warehouse, err error)
	GetAll() (warehouses []*Warehouse, err error)
	ReportProducts(id int) (productReport *ProductReport, err error)
	Create(warehouse *Warehouse) (err error)
	Update(warehouse *Warehouse) (err error)
	Delete(id int) (err error)
}

func NewMySqlRepositoty(db *sql.DB) Repository {
	return &MySqlRepositoty{db: db}
}

type MySqlRepositoty struct {
	db *sql.DB
}

func (m *MySqlRepositoty) GetAll() (warehouses []*Warehouse, err error) {
	query := `SELECT id, name, adress, telephone, capacity FROM warehouses`
	row, err := m.db.Query(query)

	for row.Next() {
		warehouse := &Warehouse{}
		err = row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
		if err != nil {
			return
		}
		warehouses = append(warehouses, warehouse)
	}
	return
}

func (m *MySqlRepositoty) Get(id int) (warehouse *Warehouse, err error) {
	warehouse = &Warehouse{}
	query := `SELECT id, name, adress, telephone, capacity FROM warehouses WHERE id = ?`
	row := m.db.QueryRow(query, id)

	err = row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Address, &warehouse.Telephone, &warehouse.Capacity)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return
	}
	return
}

func (m *MySqlRepositoty) ReportProducts(id int) (productReport *ProductReport, err error) {
	productReport = &ProductReport{}
	query := `SELECT w.name, COUNT(p.id) FROM warehouses w 
    		  INNER JOIN products p ON p.id_warehouse = w.id
              WHERE w.id = ?             
              GROUP BY w.name`
	row := m.db.QueryRow(query, id)

	err = row.Scan(&productReport.WarehouseName, &productReport.ProductCount)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return
	}
	return
}

func (m *MySqlRepositoty) Create(warehouse *Warehouse) (err error) {
	stmt, err := m.db.Prepare(`INSERT INTO warehouses VALUES(null, ?, ?, ?, ?)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(warehouse.Name, warehouse.Address, warehouse.Telephone, warehouse.Capacity)
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
	warehouse.ID = int(lastId)
	return
}

func (m *MySqlRepositoty) Update(warehouse *Warehouse) (err error) {
	stmt, err := m.db.Prepare(`UPDATE warehouses SET name = ?, adress = ?, telephone = ?, capacity = ? WHERE id = ?`)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(warehouse.Name, warehouse.Address, warehouse.Telephone, warehouse.Capacity, warehouse.ID)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}
	return
}

func (m *MySqlRepositoty) Delete(id int) (err error) {
	stmt, err := m.db.Prepare(`DELETE FROM warehouses WHERE id = ?`)
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
