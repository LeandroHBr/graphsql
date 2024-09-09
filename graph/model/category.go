package model

type Category struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	IDCategory  string  `json:"id_category"`
}
