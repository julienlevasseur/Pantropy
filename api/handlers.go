package pantropy

import (
	"net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	routes := []string{"providers"}
	json.NewEncoder(w).Encode(routes)
}

func ProvidersIndex(w http.ResponseWriter, r *http.Request) {
	providers := []Provider{
		Provider{Name: "aws"},
		Provider{Name: "chef"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(providers); err != nil {
		panic(err)
	}
}

func ProvidersChef(w http.ResponseWriter, r *http.Request) {
	resources := []Resource{
		Resource{
			Name: "chef_data_bag",
			Arguments: []Argument{
				Argument{
					Name: "name",
					Required: true,
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resources)
}