package domain

type Animal struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	LegsCount int8   `json:"legs_count"`
}
