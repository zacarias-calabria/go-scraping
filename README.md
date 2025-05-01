# Go Scraper

Una aplicación de scraping web escrita en Go que permite extraer información de sitios web de manera eficiente y configurable.

## Características

- Scraping asíncrono de sitios web
- Configuración de límites de velocidad
- Rotación automática de User-Agent
- Soporte para proxy
- Control de tiempo de espera
- Seguimiento de enlaces con profundidad configurable
- Manejo de errores y códigos de estado HTTP

## Requisitos

- Go 1.24 o superior
- Docker (opcional, para ejecución en contenedor)

## Instalación

1. Clona el repositorio:
```bash
git clone https://github.com/tu-usuario/go-scraping.git
cd go-scraping
```

2. Instala las dependencias:
```bash
go mod download
```

## Uso

La aplicación se puede ejecutar directamente o mediante Docker.

### Ejecución directa

```bash
go run main.go -url https://ejemplo.com [-timeout 30s] [-proxy http://proxy:puerto]
```

Parámetros:
- `-url`: URL del sitio a scrapear (requerido)
- `-timeout`: Tiempo máximo de espera (por defecto: 30s)
- `-proxy`: URL del proxy a utilizar (opcional)

### Ejecución con Docker

1. Construye la imagen:
```bash
docker compose build
```

2. Inicia el contenedor:
```bash
docker compose up -d
```

3. Ejecuta el scraper dentro del contenedor:
```bash
docker compose exec app go run main.go -url https://ejemplo.com
```

Para detener el contenedor:
```bash
docker compose down
```

## Estructura del Proyecto

```
.
├── internal/
│   ├── adapters/     # Implementaciones de adaptadores
│   ├── domain/       # Modelos de dominio
│   └── ports/        # Interfaces de puertos
├── main.go          # Punto de entrada de la aplicación
├── go.mod           # Archivo de dependencias
├── go.sum           # Suma de verificación de dependencias
├── Dockerfile       # Configuración de Docker
└── compose.yaml     # Configuración de Docker Compose
```

## Desarrollo

### Pruebas

Para ejecutar las pruebas:
```bash
go test ./...
```

### Depuración

La aplicación está configurada para depuración con Delve. Para depurar:

1. Inicia el servidor de depuración:
```bash
dlv debug
```

2. Conéctate al puerto 40000 desde tu IDE.

## Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo LICENSE para más detalles. 
