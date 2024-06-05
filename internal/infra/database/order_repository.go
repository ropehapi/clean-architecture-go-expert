package database

import (
	"database/sql"
	"github.com/ropehapi/clean-architecture-go-expert/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{DB: DB}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.DB.Prepare("INSERT INTO orders (id, price, tax, final_price) values(?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.DB.QueryRow("SELECT count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
