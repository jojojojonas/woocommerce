# WooCommerce
This is the first version of a WooCommerce library in GO lang. Currently the functionality is not yet complete. You can currently only read the orders. Either all orders ever made or in a certain time period.

## Install
First you have to install the package:

```console
go get github.com/jojojojonas/woocommerce
```

## How to use?
As already mentioned, in the current version you can only read out orders. But hopefully we will be able to extend it with time. It is planned. Only currently we only need the orders.

### Get all orders
In order to get all orders delivered back you should proceed as follows. We have built structs exactly the way WooCommerce uses them in the API.

```go
response, err := woocommerce.Orders(woocommerce.OrderData("shop.test.de", "ck_", "cs_", false, "", ""))
if err != nil {
    fmt.Println("The following error occurred while loading all orders: ", err)
}
```

### Get orders in a period
To find orders in a certain period of time we simply add a few parameters to the function. Now we add the start date and the end date.

```go 
response, err := woocommerce.Orders(woocommerce.OrderData("shop.test.de", "ck_", "cs_", true, "2020-11-01", "2020-11-30"))
if err != nil {
    fmt.Println("The following error occurred while loading the orders in the period: ", err)
}
```

### Get reports
If you want to read out the reports, you can do this as follows. For the period there are the following parameters: week, month, last_month and year.

```go
response, err := woocommerce.Reports("shop.test.de", "ck_", "cs_", "period")
if err != nil {
	fmt.Println("The following error occurred while loading the reports: ", err)
}
```