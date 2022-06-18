# pokemon-api-example

Ejemplo básico de un API de pokemon para validar conceptos de CI/CD y Postman

## Rutas

```
GET /
{"message": "Hola mundo"}

GET /pokemon
[
  {"id": 1, "nombre": "pikachu"},
  {"id": 2, "nombre": "charizard"},
  {"id": 3, "nombre": "bulbasaur"},
  {"id": 4, "nombre": "mewto"}
]

GET /pokemon/:id
Con id=1 ->
{
  "id":1,
  "nombre":"pikachu",
  "tipo":"electricidad"
}
```

## Desarrollo

```bash
# Instalar dependencias
go install

# Ejecutar proyecto
go run main.go
```

## Construcción del proyecto

```bash
# Instalar dependencias
go install

# Construir ejecutable/binario
go build

# Ejecutar proyecto
./pokemon-api-example
```

## Docker

```bash
# Construir imagen con docker
docker build . -t estephany16/pokemon-api

# Ejecutar tests con docker compose
docker compose up --exit-code-from newman
```
