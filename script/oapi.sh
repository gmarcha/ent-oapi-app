#! /bin/bash

oapi-codegen -generate gin ent/openapi.json > internal/api/gen/gin.go
# oapi-codegen -generate types ent/openapi.json > internal/api/gen/types.go