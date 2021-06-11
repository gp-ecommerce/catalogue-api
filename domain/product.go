package domain

type Product struct {
	Name        string                 `json:"name"`
	Id          string                 `json:"id"`
	Description string                 `json:"description"`
	Feature     map[string]interface{} `json:"feature"`
	Number      uint                   `json:"number"`
}
