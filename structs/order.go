package structs

// OrderView : Architecture rendered to view
type OrderView struct {
	OrderID      int      `json:"orderID"`
	ProductsCode []string `json:"productsCode"`
}
