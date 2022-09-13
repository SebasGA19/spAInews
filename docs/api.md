# Auth API

This documentation container everything related to the auth API.

## Config

To configure the API client it is important to specify this environment variables:

| Environment variable name | Example value                                                | Description                                                  |
| ------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| REDIS_ADDRESS             | 127.0.0.1:6379                                               | Address of the redis server                                  |
| SMTP_HOST                 | smtp.gmail.com                                               | Host of the SMTP server                                      |
| SMTP_POST                 | 465                                                          | Port of the SMTP server                                      |
| SMTP_FROM                 | spainnews@gmail.com                                          | From email to be used by the client.                         |
| SMTP_USERNAME             | spainnews@gmail.com                                          | SMTP Username                                                |
| SMTP_PASSWORD             | password                                                     | SMTP password                                                |
| MARIA_URL                 | spainews:spainews@tcp(127.0.0.1:3306)/spainews?parseTime=true | This is the URL specified in [The library documentation](https://github.com/go-sql-driver/mysql) used to connect to the MariaDB |
| MONGO_URL                 | mongodb://spainews:spainews@mongo:27017/spainews             | This is the URL specified in [The library documentation](https://github.com/mongodb/mongo-go-driver) used to connect to the MongoDB |
| LISTEN_ADDRESS            | 0.0.0.0:5000                                                 | This is address that the application is going to use for listening and accepting incoming connections |

## Errors

Possible error messages are:

| Code | Error                        | Description                                                  |
| ---- | ---------------------------- | ------------------------------------------------------------ |
| 0    | Internal server error        | Internal server error means something went wrong with our databases connections or other kind OS operation |
| 1    | Username not available       | Means the username is not available                          |
| 2    | Email not available          | Means the email is not available                             |
| 3    | Credentials not submitted    | No credentials given in login                                |
| 4    | Invalid username or password | Invalid username or password                                 |
| 5    | Permission denied            | User doesn't have permission to complete the transaction     |

Error JSON structure is:

```json
{
    "code": 0,
    "message": "error message"
}
```

## Routes

### PUT `/register`

Requests a registration for a new user in the application. On success send confirmation email to address specified.

#### Request

- METHOD: `PUT`
- Headers:
  - `Content-Type: application/json`
- Body:

```json
{
    "username": "john",
    "email": "john@gmail.com",
    "password": "super_password_123"
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Username not available`
- `Email not available`

#### Response

Empty `200 OK` response.

#### Email

User receives an email with the confirmation URL / CODE

### POST `/confirm/registration`

Confirm register request.

#### Request

- Method: `POST`
- Headers:
  - `Confirm-Code: CONFIRMATION_CODE_SENT_BY_EMAIL`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`

#### Response

Empty `200 OK` response.

### GET `/session`

Returns a new session cookie related to the user.

#### Request

- Method: `GET`
- Headers:
  - `Authorization: Basic base64(<username>:<password>)`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Credentials not submitted`

- `Invalid username or password`

#### Response

- Status: 200
- Body:

```json
{
    "session": "COOKIE"
}
```

### DELETE `/session`

Destroys de cookies created by `GET /session`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`

#### Response

- Status: 200

### POST `/password`

Changes the password of the account given the old one and the new one.

#### Request

- Method: `POST`
- Headers:
  - `Session: COOKIE`
  - `Content-Type: application/json`
- Body:

```json
{
    "old-password": "user_old_password",
    "new-password": "user_new_password"
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200

### GET `/account`

Returns account information of the related session given.

#### Request

- Method: `GET`
- Headers:
  - `Session: COOKIE`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200
- Body:

```json
{
    "username": "username",
    "email": "email@mail.com",
    "creation-date": "31/12/2000"
}
```

#### POST `/reset/password`

Requests a password reset, useful in case the user forgets its credentials. It sent an email to the account with the code that will be used to reset the credentials.

#### Request

- Method: `POST`
- Headers:
  - `Content-Type: application/json`
- Body:

```json
{
    "email": "email@mail.com"
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`

#### Response

- Method: 200

#### Email

User receives an email with the reset URL / CODE

#### POST `/confirm/reset/password`

Ensures the reset of the credentials of the user using the CODE provided by email.

#### Request

- Method: `POST`
- Headers:
  - `Confirm-Code: RESET_CODE`
  - `Content-Type: application/json`
- Body:

```json
{
    "new-password": "new_user_password"
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200

### POST `/email`

Requests an update for the email address. Send to the new email a confirmation URL / CODE

#### Request

- Method: `POST`
- Headers:
  - `Session: COOKIE`
  - `Content-Type: application/json`
- Body:

```json
{
    "password": "password",
    "new-email": "email@mail.com"
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Method: 200

#### Email

User receives an email with the confirm email URL / CODE

#### POST `/confirm/email`

Ensures the new email provided if owned by the user

#### Request

- Method: `POST`
- Headers:
  - `Confirm-Code: RESET_CODE`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200

### GET `/words`

Return the configured key words for the current account.

#### Request

- Method: `GET`
- Headers:
  - `Session COOKIE`

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200
- Headers:
  - `Content-Type: application/json`
- Body:

```json
{
    "automatic": true,
    "words": ["word-1", "word-2", "..."]
}
```

### POST `/words`

Updates the selected words for the application.

#### Request

- Method: `POST`
- Headers:
  - `Session: COOKIE`
  - `Content-Type: application/json`
- Body:

```json
{
    "automatic": true,
    "words": ["word-1", "word-2", "..."]
}
```

#### Possible Errors

Any non `200` response is considered an error.

- `Internal server error`
- `Permission denied`

#### Response

- Status: 200

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
    "number-of-results": 10,
    "news": [
        {
            "title": "",
            "description": "",
            "maintext": "",
            "authors": [],
            "category": "",
            "date_publish": "31/12/2022",
            "source_domain": "",
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
    "start-date": "31/12/2022",
    "end-date": "31/12/2022",
    "sources": ["apnews.com", "nytimes.com", "..."],
    "maintext-words": ["Colombia", "Venezuela", "..."],
    "title-words": ["Colombia", "Venezuela", "..."],
    "old-first": true,
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
    "number-of-results": 10,
    "news": [
        {
            "title": "",
            "description": "",
            "maintext": "",
            "authors": [],
            "category": "",
            "date_publish": "31/12/2022",
            "source_domain": "",
            "url": ""
        },
        "..."
    ]
}
```

