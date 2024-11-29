# Etapa 1: Base para desenvolvimento
FROM golang:1.23-alpine AS dev

# Instalar ferramentas necessárias
RUN apk add --no-cache git curl && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s && \
    mv ./bin/air /usr/local/bin/air

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar o restante do código
COPY . .

# Instalar o Air para hot-reload
RUN [ ! -f .air.toml ] && air init || echo ".air.toml já existe, ignorando 'air init'"

# Expor a porta padrão
EXPOSE 3000

# Comando padrão do contêiner (inicia o Air)
CMD ["air"]