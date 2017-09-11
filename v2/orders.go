package main

import (
	"net/http"
	"fmt"
	"time"
)

type Order struct {
	Id    int
	Type  string
	Owner string
	//CostCentre int
	//Stage      string
	//Status     string
	Created time.Time
	Expiry  time.Time
	//Deleted    time.Time
}

type Orders []Order

var orderId int
var orders Orders

func orders_list(w http.ResponseWriter, r *http.Request) {
	json_encoder(w, http.StatusOK, orders)
}

func init() {
	fmt.Println("Init: Orders")

	RepoCreateorder(Order{Type: "ami",})
	RepoCreateorder(Order{Type: "poc",})
	RepoCreateorder(Order{Type: "ami",})
	RepoCreateorder(Order{Type: "dev",})
}

func RepoCreateorder(order Order) Order {
	orderId += 1
	order.Id = orderId
	order.Created = time.Now()
	order.Expiry = time.Now().AddDate(0, 0, 60)

	orders = append(orders, order)
	return order
}


