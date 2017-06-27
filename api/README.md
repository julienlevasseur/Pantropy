# Pantropy API

## Overview

The Pantropy API expose 2 types of entities :

- The getter routes
- The setter routes

### Getter routes

The role of the getter routes is to provide a declarative structure for the interface forms.

For example, when a user want to create a chef_node resource, the frontend make an HTTP request to `/provider/chef/chef_node` to retrieve these data :

```javascript
{
  "name":"chef_node",
  "arguments": [
  	{
  	  "name": "name",
  	  "required": true,
  	  "description": "The unique name to assign to the node.",
  	  "options": null
  	},
  	{
  	  "name": "automatic_attributes_json",
  	  "required": false,
  	  "description": "String containing a JSON-serialized object containing the automatic attributes for the node.",
  	  "options": null
  	},
  	{
  	  "name": "normal_attributes_json ",
  	  "required": false,
  	  "description": "String containing a JSON-serialized object containing the normal attributes for the node.",
  	  "options": null
  	},
  	{
  	  "name": "default_attributes_json ",
  	  "required": false,
  	  "description": "String containing a JSON-serialized object containing the default attributes for the node.",
  	  "options": null
  	},
  	{
  	  "name": "override_attributes_json",
  	  "required": false,
  	  "description": "String containing a JSON-serialized object containing the override attributes for the node.",
  	  "options": null
  	},
  	{
  	  "name": "run_list",
  	  "required": false,
  	  "description": "List of strings to set as the run list for the node.",
  	  "options": null
  	}
  ]
}
```

Those datas are used to construct the form (whic field must be displayed, which field are mandatory, what is its description .. )

### Setter routes

The role of the setter routes is to provide a feature to generate an HCL file from a previously created resource (via the UI).

## Routes

Route | Description
----- | -----------
`/` | List the supported features
`/providers` | List the supported providers
`/providers/{provider_name}` | List the provider's resources
`/providers/{provider_name}/{resource_name}` | Return the specific resource