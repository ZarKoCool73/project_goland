# Etapa 1: Build
FROM golang:1.25.4-alpine AS builder

# Instala git y herramientas necesarias
RUN apk add --no-cache git

# Establece el directorio de trabajo
WORKDIR /app

# Copia go.mod y go.sum
COPY go.mod ./

# Copia todo el proyecto
COPY . .

# Descarga dependencias
RUN go mod download

# Compila la aplicación
RUN go build -o go-api main.go

# Etapa 2: Runtime
FROM alpine:latest

# Directorio de trabajo
WORKDIR /app

# Copia el binario desde el builder
COPY --from=builder /app/go-api .

# Expone el puerto
EXPOSE 3000

# Comando por defecto
CMD ["./go-api"]
