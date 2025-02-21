# Definição do nome do binário
BINARY_NAME=backup
TGZ_NAME=$(BINARY_NAME)-linux-X86_64.tar.gz

# Comandos
GO=go
GOBUILD=$(GO) build
GORUN=$(GO) run
GOTEST=$(GO) test
GOCLEAN=$(GO) clean
GOFMT=$(GO) fmt
GOLINT=golangci-lint run
TAR=tar
TARGZ=$(TAR) -czf
# Diretórios
CMD_DIR=cmd
FRONTEND_DIR=frontend

.PHONY: build run runc


## Compila o backend
build:
	@clear
	@echo "🔨 Compilando o backend..."
	cd $(CMD_DIR) && $(GOBUILD) -o $(BINARY_NAME)
	cd $(CMD_DIR) && $(TARGZ) $(TGZ_NAME) $(BINARY_NAME) 

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
