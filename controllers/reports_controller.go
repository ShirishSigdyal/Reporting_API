package controllers

import (
	"encoding/json"
	"net/http"
	"reporting-api/services"
)

type ReportsController struct {
	service *services.ReportService
}

func NewReportsController(service *services.ReportService) *ReportsController {
	return &ReportsController{service: service}
}

func (c *ReportsController) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/reports/sales", c.GetSalesReport)
	mux.HandleFunc("/reports/customers", c.GetCustomerReport)
	return mux
}

func (c *ReportsController) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	response, err := c.service.GenerateSalesReport(r.Context(), r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (c *ReportsController) GetCustomerReport(w http.ResponseWriter, r *http.Request) {
	response, err := c.service.GenerateCustomerReport(r.Context(), r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
