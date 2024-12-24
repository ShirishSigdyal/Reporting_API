package models

// Customer represents the customers table.
type Customer struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	SignupDate   string  `json:"signup_date"`
	Location     string  `json:"location"`
	LifetimeValue float64 `json:"lifetime_value"`
}

// Product represents the products table.
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

// Order represents the orders table.
type Order struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	OrderDate  string `json:"order_date"`
	Status     string `json:"status"`
}

// OrderItem represents the order_items table.
type OrderItem struct {
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// Transaction represents the transactions table.
type Transaction struct {
	ID            int     `json:"id"`
	OrderID       int     `json:"order_id"`
	PaymentStatus string  `json:"payment_status"`
	PaymentDate   string  `json:"payment_date"`
	TotalAmount   float64 `json:"total_amount"`
}
