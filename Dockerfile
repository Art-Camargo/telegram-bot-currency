# Etapa 1: build (com Go instalado)
FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

# Etapa 2: imagem final (sem Go instalado)
FROM debian:bullseye-slim

WORKDIR /opt/app

# Copia só o binário para a imagem final
COPY --from=builder /app/app .

# (Opcional) adiciona um usuário não root
RUN useradd -m appuser
USER appuser

# Expõe porta, se necessário (ex: 8080)
# EXPOSE 8080

ENTRYPOINT ["./app"]
