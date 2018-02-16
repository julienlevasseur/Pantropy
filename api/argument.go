package main

// Argument definition
type Argument struct {
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
}

// Option definition
type Option struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}
