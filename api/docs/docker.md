# Docker

## Overview

The docker routes offers a support for docker functionalities (such as: building, pushing images).

## Routes

### POST /v1/docker/build

Build a docker image.

* Content-Type: "application/json"
* Accept: "application/json"

Include `Dockerfile url` and `tag` in the body:

```
{
	"dockerfile": "https://raw.githubusercontent.com/owner/repo/branch/Dockerfile",
	"tag": "owner/image_name:version"
}
```

#### Example response

* Status: 201
* Content-Type: "application/json"

```
{
	"id": "image_id"
}
```

### POST /v1/docker/push

Push a docker image.

* Content-Type: "application/json"
* Accept: "application/json"

Include `image id` in the body:

```
{
	"tag": "owner/image_name:version"
}
```

Credentials for login to docker hub will be expected as environment variables.

#### Example response

* Status: 201
