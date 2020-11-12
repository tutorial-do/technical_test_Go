package models

// Comprador Representation of buyer
type Comprador struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Producto Representation of products
type Producto struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Transaccion Representation of transactions
type Transaccion struct {
	ID        int    `json:"id"`
	BuyerID   int    `json:"buyer_id"`
	IP        string `json:"ip"`
	Device    string `json:"device"`
	ProductID int    `json:"product_id"`
}
