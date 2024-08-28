# Nome do comando de execução
RUN_CMD = go run main.go wire_gen.go

# Diretório do servidor
SERVER_DIR = cmd/ordersystem

# Nome do executável
APP_NAME = desafio-clean-architecture-go-expert

# Regra padrão para rodar o servidor
.PHONY: run
run:
	cd $(SERVER_DIR) && $(RUN_CMD)

# Regra para limpar (não é obrigatória mas pode ser útil)
.PHONY: clean
clean:
    # Aqui você pode adicionar comandos para limpar build artifacts, logs, etc.
	@echo "Clean is not implemented yet."

# Regra para instalar dependências (opcional)
.PHONY: deps
deps:
	go mod tidy

# Regra para testar (opcional)
.PHONY: test
test:
	go test ./...

# Regra para build (opcional)
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o $(APP_NAME) $(SERVER_DIR)/main.go $(SERVER_DIR)/wire_gen.go

# Regra para rodar o build
.PHONY: start
start: build
	./$(APP_NAME)