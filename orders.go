// Getting the orders from your woocommerce api
//
// From Jonas Kwiedor <support@pikbat.de>
// For my company J&J Ideenschmiede UG
//
// In this file you can get all orders from the woocommerce api or all orders from a time period.
// All documentation is on my github profile github.com/jojojojonas/woocommercev3

package woocommerce

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Order data for the request
type OrderData struct {
	Domain         string
	ConsumerKey    string
	ConsumerSecret string
	Period         bool
	Start          string
	End            string
}

// For save the data from wordpress
type Order struct {
	ID                 int             `json:"id"`
	ParentID           int             `json:"parent_id"`
	Number             string          `json:"number"`
	OrderKey           string          `json:"order_key"`
	CreatedVia         string          `json:"created_via"`
	Version            string          `json:"version"`
	Status             string          `json:"status"`
	Currency           string          `json:"currency"`
	DateCreated        string          `json:"date_created"`
	DateCreatedGMT     string          `json:"date_created_gmt"`
	DateModified       string          `json:"date_modified"`
	DateModifiedGMT    string          `json:"date_modified_gmt"`
	DiscountTotal      string          `json:"discount_total"`
	DiscountTax        string          `json:"discount_tax"`
	ShippingTotal      string          `json:"shipping_total"`
	ShippiungTax       string          `json:"shippiung_tax"`
	CartTax            string          `json:"cart_tax"`
	Total              string          `json:"total"`
	TotalTax           string          `json:"total_tax"`
	PricesIncludeTax   bool            `json:"prices_include_tax"`
	CustomerID         int             `json:"customer_id"`
	CustomerIPAdress   string          `json:"customer_ip_adress"`
	CustomerUserAgent  string          `json:"customer_user_agent"`
	CustomerNote       string          `json:"customer_note"`
	Billing            Billing         `json:"billing"`
	Shipping           Shipping        `json:"shipping""`
	PaymentMethod      string          `json:"payment_method"`
	PaymentMethodTitle string          `json:"payment_method_title"`
	TransactionID      string          `json:"transaction_id"`
	DatePaid           string          `json:"date_paid"`
	DatePaidGMT        string          `json:"date_paid_gmt"`
	DateCompleted      string          `json:"date_completed"`
	DateCompletedGMT   string          `json:"date_completed_gmt"`
	CartHash           string          `json:"cart_hash"`
	MetaData           []MetaData      `json:"meta_data"`
	LineItems          []LineItems     `json:"line_items"`
	TaxLines           []TaxLines      `json:"tax_lines"`
	ShippingLines      []ShippingLines `json:"shipping_lines"`
	FeeLines           []interface{}   `json:"fee_lines"`
	CouponLines        []interface{}   `json:"coupon_lines"`
	refunds            []interface{}   `json:"refunds"`
	CurrencySymbol     string          `json:"currency_symbol"`
	Links              Links           `json:"_links"`
}

type Billing struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Comapny   string `json:"comapny"`
	Adress1   string `json:"adress_1"`
	Adress2   string `json:"adress_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Shipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Comapny   string `json:"comapny"`
	Adress1   string `json:"adress_1"`
	Adress2   string `json:"adress_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}

type MetaData struct {
	ID    int         `json:"id"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type LineItems struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProductID   int        `json:"product_id"`
	VariationID int        `json:"variation_id"`
	Quantity    int        `json:"quantity"`
	TaxClass    string     `json:"tax_class"`
	Subtotal    string     `json:"subtotal"`
	SubtotalTax string     `json:"subtotal_tax"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
	SKU         string     `json:"sku"`
	Price       float64    `json:"price"`
}

type Taxes struct {
	ID       int    `json:"id"`
	Total    string `json:"total"`
	Subtotal string `json:"subtotal"`
}

type TaxLines struct {
	ID               int        `json:"id"`
	RateCode         string     `json:"rate_code"`
	RateID           int        `json:"rate_id"`
	Label            string     `json:"label"`
	Compound         bool       `json:"compound"`
	TaxTotal         string     `json:"tax_total"`
	ShippingTaxTotal string     `json:"shipping_tax_total"`
	RatePercent      int        `json:"rate_percent"`
	MetaData         []MetaData `json:"meta_data"`
}

type ShippingLines struct {
	ID          int        `json:"id"`
	MethodTitle string     `json:"method_title"`
	MethodID    string     `json:"method_id"`
	InstanceID  string     `json:"instance_id"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
}

type Links struct {
	Self       []Self       `json:"self"`
	Collection []Collection `json:"collection"`
}

type Self struct {
	HREF string `json:"href"`
}

type Collection struct {
	HREF string `json:"href"`
}

var (
	url    string
	page   int = 1
	orders []Order
)

func Orders(data OrderData) ([]Order, error) {

	// Get the url for request
	url = "https://" + data.Domain + "/wp-json/wc/v3/orders?consumer_key=" + data.ConsumerKey + "&consumer_secret=" + data.ConsumerSecret + "&per_page=100"

	// If time period is true
	if data.Period {
		url = url + "&after=" + data.Start + "T00:00:00Z&before=" + data.End + "T24:00:00Z"
	}

	// Loop over the sites
	for {

		// Define request with GET
		response, err := http.Get(fmt.Sprintf("%s&page=%d", url, page))
		if err != nil {
			return nil, err
		}

		// Decode data
		var decode []Order

		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return nil, err
		}

		// Check length of response and break it, when site content is null
		if len(decode) == 0 {
			break
		}

		// Save decode data to orders
		for _, value := range decode {

			// Append to order
			orders = append(orders, value)

		}

		// Add a number to page.
		page++

	}

	// Return
	return orders, nil

}
