package domain

type Customers struct {
	Id        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition bool   `json:"condition"`
}

type CustomerConditionsTotals struct {
	Condition bool    `json:"condition"`
	Total     float64 `json:"total"`
}
