# Etapa 1: Construção
FROM golang:1.23-alpine AS builder

# Definindo o diretório de trabalho
WORKDIR /app

# Copiar o arquivo go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências (go mod tidy)
RUN go mod tidy

# Copiar o restante do código-fonte para o contêiner
COPY . .

# Compilar o binário da aplicação
RUN go build -o /tmp/main ./cmd/api

# Etapa 2: Execução
FROM alpine:latest

# Definindo o diretório de trabalho no contêiner final
WORKDIR /root/

# Copiar o binário compilado da etapa de construção
COPY --from=builder /tmp/main /tmp/main

# Comando para rodar o binário
CMD ["/tmp/main"]