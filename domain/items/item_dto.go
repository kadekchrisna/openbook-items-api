package items

// Item struct
type Item struct {
	Id                string      `json:"id"`
	Title             string      `json:"title"`
	Seller            int64       `json:"seller"`
	Description       Description `json:"description"`
	Picture           []Picture   `json:"picture"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:'available_quantity'`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plaintext"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"id"`
}
