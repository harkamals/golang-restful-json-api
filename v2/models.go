package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Order struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`

	Type  string
	Owner string
	//CostCentre int
	//Stage      string
	//Status     string
	Created time.Time
	Expiry  time.Time
	//Deleted    time.Time
}

func (o *Order) getOrder(db *sql.DB) error {
	return db.QueryRow("SELECT id, name, price FROM orders where id=$1", o.Id).Scan(&o.Name, &o.Price)
}

func (o *Order) updateOrder(db *sql.DB) error {
	_, err := db.Exec("UPDATE orders SET name=$1, price=$2 WHERE id=$3", o.Name, o.Price, o.Id)
	return err
}

func (o *Order) deleteOrder(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM orders WHERE id=$1", o.Id)
	return err

}

func (o *Order) createOrder(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO orders (name, price) VALUES ($1, $2) RETURNING id", o.Name, o.Price).Scan(&o.Id)
	if err != nil {
		return err
	}
	return nil
}

func getOrders(db *sql.DB, start, count int) ([]Order, error) {
	rows, err := db.Query("SELECT id, name, price FROM orders LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	orders := []Order{}

	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.Id, &o.Name, &o.Price); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil

}

//type Orders []Order
//
//var orderId int
//var orders Orders
//
//func orders_list(w http.ResponseWriter, r *http.Request) {
//	respondWithJSON(w, http.StatusOK, orders)
//}

func init() {
	fmt.Println("Init: Orders")

	//RepoCreateorder(Order{Type: "ami",})
	//RepoCreateorder(Order{Type: "poc",})
	//RepoCreateorder(Order{Type: "ami",})
	//RepoCreateorder(Order{Type: "dev",})
}

//func RepoCreateorder(order Order) Order {
//	orderId += 1
//	order.Id = orderId
//	order.Created = time.Now()
//	order.Expiry = time.Now().AddDate(0, 0, 60)
//
//	orders = append(orders, order)
//	return order
//}
