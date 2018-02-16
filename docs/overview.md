# Overview

Pantropy meant to be a full set of tools to manage application workflow (from dev to prod).

## Components

* Development container - is a docker container to develop code on an isolated space
* Development squad - is a docker compose group of containers to develop a multi part application on isolated spaces.

## Requirements

```bash
go get github.com/fatih/color
go get github.com/stretchr/testify
```

## API

### Infrastructure Level

#### GET /v1/infra

List the supported infrastructure resource types.

#### Example response

* Status: 200
* Content-Type: "application/json"

```
[
	"providers"
]
```

### Application Level