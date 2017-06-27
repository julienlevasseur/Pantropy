package main

type Argument struct {
	Name 		string 		`json:"name"`
	Required	bool 		`json:"required"`
	Options 	[]Option 	`json:"options"`
}

type Option struct {
	Name 		string 	`json:"name"`
	Value 		string 	`json:"value"`
	Required	bool 	`json:"required"`
}