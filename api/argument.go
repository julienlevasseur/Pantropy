package pantropy

type Argument struct {
	Name 		string 				`json:"name"`
	Required	bool 				`json:"required"`
	options 	map[string]string 	`json:"options"`
}
