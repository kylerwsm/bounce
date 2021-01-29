#!/usr/bin/env bash

curl --request GET http://localhost:8080/kylerwsm

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"link":"https://kylerwsm.com"}' \
  http://localhost:8080/kylerwsm

curl --request GET http://localhost:8080/kylerwsm

curl --request DELETE http://localhost:8080/kylerwsm

curl --request GET http://localhost:8080/kylerwsm
