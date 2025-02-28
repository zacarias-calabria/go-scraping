# Utilizamos la imagen oficial de Go en su última versión
FROM golang:latest

# Actualizamos los paquetes e instalamos dependencias necesarias para compilar y depurar
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    make \
 && rm -rf /var/lib/apt/lists/*

# Instalamos Delve, el depurador para Go, para facilitar la integración con el IDE de JetBrains
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Instalamos golangci-lint, una herramienta para validación y análisis de código
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

# Establecemos el directorio de trabajo
WORKDIR /app

# Exponemos puertos necesarios:
# - 40000 para el depurador (Delve)
# - 8080 (por si la aplicación expone algún servicio HTTP)
EXPOSE 40000
EXPOSE 8080

# De momento, mantenemos el contenedor en ejecución para trabajar en modo interactivo
CMD ["tail", "-f", "/dev/null"]
