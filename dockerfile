### Stage 1: Baixar dependencias e compilar o binário
FROM golang:1.22-alpine as stage1

WORKDIR /go/src/app

# Copia o código da aplicação e compila o binário
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

### Stage 2:
FROM scratch

# Copia apenas o binário gerado no stage anterior
COPY --from=stage1 /go/src/app/main /

# Define o ponto de entrada para o container como /main.
# O binario será executado quando o container for iniciado.
ENTRYPOINT [ "/main" ]
