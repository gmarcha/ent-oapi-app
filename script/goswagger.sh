#! /bin/bash

### Generate api code for swagger specification.
#goswagger generate server -f ./api/swagger.yml -A rest-api
goswagger generate server -t internal -f ./api/swagger.yml -A rest-api --exclude-main