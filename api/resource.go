package main

// Resource definition
type Resource struct {
	Name      string     `json:"name"`
	Arguments []Argument `json:"arguments"`
}
