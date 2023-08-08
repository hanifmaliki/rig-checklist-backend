.PHONY: test run env-vault

export VAULT_HOST ?= http://127.0.0.1:8200
# export VAULT_TOKEN ?= token
export ENVCONSUL_CONFIG ?= ./config.hcl
# export ENVCONSUL_SECRET_PATH ?= kv/secret

test:
	gofmt -w . && go test -v ./pkg/... -cover -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

run:
	gofmt -w . && go run ./cmd/api/main.go

env-vault:
	echo "./envconsul -config="$(ENVCONSUL_CONFIG)" -vault-addr="$(VAULT_HOST)" -once ./executable" > cmd.sh
