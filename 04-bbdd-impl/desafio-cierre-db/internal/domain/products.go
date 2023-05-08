package domain

type Products struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type TopProducts struct {
	Description string `json:"description"`
	Total       int    `json:"total"`
}
