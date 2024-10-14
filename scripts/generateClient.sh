#!/usr/bin/env bash

# This script should be run from the root of the repository (./scripts/generateClient.sh)

set -x
set -e

# First remove any existing files that were generated. If we want to edit files post-generation, this section will need to be removed
# rm ./aic/api_*.go
# rm ./aic/model_*.go
# rm -r ./aic/docs/

docker run --rm \
    -v "$PWD:/local" openapitools/openapi-generator-cli:v7.0.1 generate \
    -i /local/spec/openapi.json \
    -g go \
    -o /local/aic \
    --git-host github.com \
    --git-repo-id aic-temp-go-client \
    --git-user-id pingidentity \
    --type-mappings=integer=int64,number=float64 \
    --additional-properties=enumClassPrefix=true,packageName=aic,useOneOfDiscriminatorLookup=true \
    --skip-validate-spec

rm -r aic/test/

# Run any code generators
go mod tidy
go generate ./...

# Clean things up
go fmt ./...
