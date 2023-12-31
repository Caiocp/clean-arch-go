package database

import (
	"database/sql"

	"github.com/caiocp/clean-arch-go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	stmt, err := o.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order

	rows, err := o.Db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order entity.Order

		err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
