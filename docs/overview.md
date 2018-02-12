# Overview

Pantropy meant to be a full set of tools to manage application workflow (from dev to prod).

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