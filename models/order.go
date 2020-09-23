package models

import (
	"log"
)

type Order struct {
	ID                int    `json:"orderNumber"`
	OrderDate         string `json:"orderDate"`
	RequiredDate      string `json:"requiredDate"`
	ShippedDate       string `json:"shippedDate"`
	Status            string `json:"status"`
	Comments          string `json:"comments"`
}

func GetLengthOrders(customer_id int) ([]Order, error){
	var orders []Order
	rows, err := DB.Query("SELECT , name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
		return orders, err
	}

	defer rows.Close()
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, 
						 &order.OrderDate,
						 &order.RequiredDate,
						 &order.ShippedDate,
						 &order.Status,
						 &order.Comments)
		if err != nil {
			log.Fatal(err)
			return orders, err
		}

		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return orders, err
	}

	return orders, nil
}