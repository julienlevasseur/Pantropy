package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var chefResources = map[string]Resource{
	"chef_data_bag": {
		Name: "chef_data_bag",
		Arguments: []Argument{
			{
				Name:     "name",
				Required: true,
			},
		},
	},
	"chef_data_bag_item": {
		Name: "chef_data_bag_item",
		Arguments: []Argument{
			{
				Name:     "name",
				Required: true,
			},
			{
				Name:     "content_json",
				Required: true,
			},
		},
	},
	"chef_environment": {
		Name: "chef_environment",
		Arguments: []Argument{
			{
				Name:     "name",
				Required: true,
			},
			{
				Name:     "description",
				Required: false,
			},
			{
				Name:     "default_attributes_json",
				Required: false,
			},
			{
				Name:     "override_attributes_json",
				Required: false,
			},
			{
				Name:     "cookbook_constraints",
				Required: false,
			},
		},
	},
	"chef_node": {
		Name: "chef_node",
		Arguments: []Argument{
			{
				Name:     "name",
				Required: true,
			},
			{
				Name:     "automatic_attributes_json",
				Required: false,
			},
			{
				Name:     "normal_attributes_json ",
				Required: false,
			},
			{
				Name:     "default_attributes_json ",
				Required: false,
			},
			{
				Name:     "override_attributes_json",
				Required: false,
			},
			{
				Name:     "run_list",
				Required: false,
			},
		},
	},
	"chef_role": {
		Name: "chef_role",
		Arguments: []Argument{
			{
				Name:     "name",
				Required: true,
			},
			{
				Name:     "description",
				Required: false,
			},
			{
				Name:     "default_attributes_json ",
				Required: false,
			},
			{
				Name:     "override_attributes_json",
				Required: false,
			},
			{
				Name:     "run_list",
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
