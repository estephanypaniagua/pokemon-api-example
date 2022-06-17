# Etapa 1: Constructor
FROM golang:1.18 as builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod .
COPY go.sum .

# Instalar dependencias
RUN go mod download

# Copiar archivos de código
COPY . /app

# Construir binario
RUN GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .


# Etapa 2: Ejecutor
FROM alpine:latest

# Instalar dependencias del sistema operativo
RUN apk add --no-cache libc6-compat

# Ejecutar binario
WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]