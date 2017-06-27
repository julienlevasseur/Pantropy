package main

import (
	"net/http"
	"encoding/json"
)

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
		Resource{
			Name: "chef_data_bag_item",
			Arguments: []Argument{
				Argument{
					Name: "name",
					Required: true,
				},
				Argument{
					Name: "content_json",
					Required: true,
				},
			},
		},
		Resource{
			Name: "chef_environment",
			Arguments: []Argument{
				Argument{
					Name: "name",
					Required: true,
				},
				Argument{
					Name: "description",
					Required: false,
				},
				Argument{
					Name: "default_attributes_json",
					Required: false,
				},
				Argument{
					Name: "override_attributes_json",
					Required: false,
				},
				Argument{
					Name: "cookbook_constraints",
					Required: false,
				},
			},
		},
		Resource{
			Name: "chef_node",
			Arguments: []Argument{
				Argument{
					Name: "name",
					Required: true,
				},
				Argument{
					Name: "automatic_attributes_json",
					Required: false,
				},
				Argument{
					Name: "normal_attributes_json ",
					Required: false,
				},
				Argument{
					Name: "default_attributes_json ",
					Required: false,
				},
				Argument{
					Name: "override_attributes_json",
					Required: false,
				},
				Argument{
					Name: "run_list",
					Required: false,
				},
			},
		},
		Resource{
			Name: "chef_role",
			Arguments: []Argument{
				Argument{
					Name: "name",
					Required: true,
				},
				Argument{
					Name: "description",
					Required: false,
				},
				Argument{
					Name: "default_attributes_json ",
					Required: false,
				},
				Argument{
					Name: "override_attributes_json",
					Required: false,
				},
				Argument{
					Name: "run_list",
					Required: false,
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resources)
	//json.MarshalIndent(json.NewEncoder(w).Encode(resources), "", "  ")
}
