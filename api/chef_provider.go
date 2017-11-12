package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var chefResources = map[string]Resource{
	"chef_data_bag": Resource{
		Name: "chef_data_bag",
		Arguments: []Argument{
			Argument{
				Name: "name",
				Required: true,
			},
		},
	},
	"chef_data_bag_item": Resource{
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
	"chef_environment": Resource{
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
	"chef_node": Resource{
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
	"chef_role": Resource{
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

// ProvidersChef : Define the chef provider function
func ProvidersChef(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(chefResources); err != nil {
		panic(err)
	}
}

// ResourceChef : Define the chef resource function
func ResourceChef(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ResourceName := vars["ResourceName"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(chefResources[ResourceName]); err != nil {
		panic(err)
	}
}
