# Definição do nome do binário
BINARY_NAME=backup

# Comandos
GO=go
GOBUILD=$(GO) build
GORUN=$(GO) run
GOTEST=$(GO) test
GOCLEAN=$(GO) clean
GOFMT=$(GO) fmt
GOLINT=golangci-lint run

# Diretórios
CMD_DIR=cmd
FRONTEND_DIR=frontend

.PHONY: build run runc


## Compila o backend
build:
	@clear
	@echo "🔨 Compilando o backend..."
	cd $(CMD_DIR) && $(GOBUILD) -o $(BINARY_NAME)

## Executa o backend
run:
	@clear
	@echo "🚀 Rodando o backend com config"
	cd $(CMD_DIR) && $(GORUN) main.go

runc:
	@clear
	@rm "cmd/config/config.yaml"
	@echo "🚀 Rodando o backend sem config"
	cd $(CMD_DIR) && $(GORUN) main.go
