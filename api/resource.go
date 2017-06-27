package pantropy

type Resource struct {
	Name		string `json:"name"`
	Arguments 	[]Argument `json:"arguments"`
}
