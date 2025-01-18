#!/bin/bash

set -euo pipefail
fuse-oapi \
  schemas/auth.openapi.json \
  schemas/deliver.openapi.json \
  schemas/verify.openapi.json \
  schemas/webhook.openapi.json \
  --output orum.openapi.json
kiota generate -l go -c OrumClient -n github.com/HitchPin/go-orum-client -d ./orum.openapi.json -o ./
go get ...
go build ...