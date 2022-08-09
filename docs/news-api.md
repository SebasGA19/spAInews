# News API

This documentation container everything related to the auth API.

## Setup

...? Connection host? Connection username/password? connection database?

## Errors

Possible error messages are:

| Code | Error                 | Description                                                  |
| ---- | --------------------- | ------------------------------------------------------------ |
| 0    | Internal server error | Internal server error means something went wrong with our databases connections or other kind OS operation |

Error JSON structure is:

```json
{
    "code": 0,
    "message": "error message"
}
```

## Routes

### GET `/news/{PAGE}`

Requests the latest news obtained from the crawler.

It will return 10 or less news as results

#### Request

- URI:
  - `PAGE`: Integer with the page number, used for the pagination functionality
- Method: `GET`

#### Possible Errors

- `Internal server error`

#### Response

- Status: `200`
- Headers:
  - `Content-Type: application/json`
- Body:

```json
{
    "current-page": 1,
    "total-pages": 10,
    "number-of-results": 10,
    "news": [
        {
          "authors": [],
          "date_download": null,
          "date_modify": null,
          "date_publish": "",
          "description": "",
          "filename": "",
          "image_url": "",
          "language": "en",
          "localpath": null,
          "source_domain": "",
          "maintext": "",
          "title": "",
          "title_page": null,
          "title_rss": null,
          "url": ""
        },
        "..."
    ]
}
```

### GET `/search/{PAGE}`

#### Request

- URI:
  - `PAGE`: Integer with the page number, used for the pagination functionality
- Method: `GET`
- Headers:
  - `Content-Type: application/json`
- Body:



```json
{
    "filter": {
        "start-date": "DAY/MONTH/YEAR",
        "end-date": "DAY/MONTH/YEAR",
        "sources": ["CNN", "RT", "NYTIMES", "..."]
    },
    "sort": "newest"
}
```

`sort`: Could be `new-first` or `old-first` 

#### Possible Errors

- `Internal server error`

#### Response

- Status: `200`
- Headers:
  - `Content-Type: application/json`
- Body:

```json
{
    "current-page": 1,
    "total-pages": 10,
    "number-of-results": 10,
    "news": [
        {
          "authors": [],
          "date_download": null,
          "date_modify": null,
          "date_publish": "",
          "description": "",
          "filename": "",
          "image_url": "",
          "language": "en",
          "localpath": null,
          "source_domain": "",
          "maintext": "",
          "title": "",
          "title_page": null,
          "title_rss": null,
          "url": ""
        },
        "..."
    ]
}
```