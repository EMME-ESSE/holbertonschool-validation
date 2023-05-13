# Configuración del binario
BINARY_NAME=awesome-api
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_VERSION=dev
COMMIT=$(shell git rev-parse --short HEAD)

# Comandos del Makefile
.PHONY: build clean run test

build:
	@echo "Compilando el binario $(BINARY_NAME)..."
	go build -ldflags "-s -w -X main.buildTime=$(BUILD_TIME) -X main.buildVersion=$(BUILD_VERSION) -X main.gitCommit=$(COMMIT)" -o $(BINARY_NAME)

clean:
	@echo "Limpiando el proyecto..."
	rm -f $(BINARY_NAME)

run:
	@echo "Iniciando la aplicación..."
	./$(BINARY_NAME)

test:
	@echo "Ejecutando pruebas..."
	go test ./...

# Configuración de ayuda
help:
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos:"
	@echo "  build   Compilar el binario"
	@echo "  clean   Limpiar el proyecto"
	@echo "  run     Iniciar la aplicación"
	@echo "  test    Ejecutar pruebas"
