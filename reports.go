package woocommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// To save reports
type Report struct {
	TotalSales     string      `json:"total_sales"`
	NetSales       string      `json:"net_sales"`
	AverageSales   string      `json:"average_sales"`
	TotalOrders    int         `json:"total_orders"`
	TotalItems     int         `json:"total_items"`
	TotalTax       string      `json:"total_tax"`
	TotalShipping  string      `json:"total_shipping"`
	TotalRefunds   int         `json:"total_refunds"`
	TotalDiscount  string      `json:"total_discount"`
	TotalGroupedBy string      `json:"total_grouped_by"`
	Totals         interface{} `json:"totals"`
	TotalCustomers int         `json:"total_customers"`
	Links          interface{} `json:"_links"`
}

// Get all reports of the last month
func Reports(domain, consumerKey, consumerSecret, period string) (Report, error) {

	// Set url
	url := fmt.Sprintf("https://%s/wp-json/wc/v3/reports/sales?consumer_key=%s&consumer_secret=%s", domain, consumerKey, consumerSecret)

	// Check period parameter
	if len(period) > 0 {
		url = fmt.Sprintf("%s&period=%s", url, period)
	}

	// Define client
	client := &http.Client{}

	// New http request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Report{}, err
	}

	// Send request to woocommerce shop
	response, err := client.Do(request)

	// Defer response body
	defer response.Body.Close()

	// Decode json data
	var decode []Report

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return Report{}, err
	}

	// Return
	return decode[0], nil

}
