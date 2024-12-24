package models

// SalesReport represents the structure of the sales report response.
type SalesReport struct {
	Category    string  `json:"category"`
	TotalSales  float64 `json:"total_sales"`
}

// CustomerReport represents the structure of the customer report response.
type CustomerReport struct {
	Segment        string  `json:"segment"`
	TotalCustomers int     `json:"total_customers"`
	AverageOrders  float64 `json:"average_orders"`
}
