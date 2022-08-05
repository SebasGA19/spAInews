# API

This documentation container everything related to the .

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

### POST `/confirm`

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

