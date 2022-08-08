## Borrador Docs news-api

### /news

- Filtrar por fechas
- Filtrar por fuentes
- Paginacion

### /search

- Filtrar por fechas
- Filtrar por fuentes
- Paginacion

## Ejemplo /news

- URI: /news/5
- Method: GET
- Header:
  - `Content-Type: application/json`
- Body:

```json
{
    "start-date": "7/08/2022",
    "end-date": "10/08/2022",
    "sources": ["cnn", "nytimes", "eltiempo"]
}
```

## Ejemplo /search

- URI: /search/5
- Method: GET
- Header:
  - `Content-Type: application/json`
- Body:

```json
{
    "query": "elecciones 2022",
    "start-date": "7/08/2022",
    "end-date": "10/08/2022",
    "sources": ["cnn", "nytimes", "eltiempo"]
}
```

