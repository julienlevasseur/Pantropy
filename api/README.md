# Pantropy API

![Build Status](https://travis-ci.org/julienlevasseur/Pantropy.svg?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9e04ed0b0adb438a86304d00abea4342)](https://www.codacy.com/app/julienlevasseur/Pantropy?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=julienlevasseur/Pantropy&amp;utm_campaign=Badge_Grade)
[![Coverage Status](https://coveralls.io/repos/github/julienlevasseur/Pantropy/badge.svg?branch=master)](https://coveralls.io/github/julienlevasseur/Pantropy?branch=master)

## Overview

The Pantropy API expose 2 types of entities :

- The getter routes
- The setter routes

### Getter routes

The role of the getter routes is to provide a declarative structure for the interface forms.

For example, when a user want to create a chef_node resource, the frontend make an HTTP request to `/v1/provider/chef/chef_node` to retrieve these data :

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

Verb | Route | Description
---- | ----- | -----------
GET  | `/v1/` | List the supported features
GET  | `/v1/providers` | List the supported providers
GET  | `/v1/providers/{provider_name}` | List the provider's resources
GET  | `/v1/providers/{provider_name}/{resource_name}` | Return the specific resource
POST | `/v1/json2hcl` | Return HCL version of the JSON passed in data

## Examples

### json2hcl

Request:
```bash
curl -X POST http://127.0.0.1:8080/v1/json2hcl \
-H "Content-Type: application/json" \
-d '{"variable":{"ami":{"default": "foo"}}}'
```

Reply:
```
"variable" "ami" {
  "default" = "foo"
}
```