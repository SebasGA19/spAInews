# API

This documentation explains the setup and functioonality of the API.

## Command line tool

```shell
api.elf HOST ENV_FILE
```

### Example

```shell
api.elf "0.0.0.0:5000" ./api.env
```

## Config (DotEnv)

To configure the API client it is important to specify this environment variables. Check `configs/api.env` as reference.

| Variable name      | Example value      | Description                   |
| ------------------ | ------------------ | ----------------------------- |
| REDIS_HOST         | 127.0.0.1          | Redis Host to connect         |
| REDIS_PORT         | 6379               | Redis port                    |
| SQL_USER           | spainews           | MariaDB user                  |
| SQL_PASS           | spainews           | User password                 |
| SQL_HOST           | 127.0.0.1          | SQL host                      |
| SQL_PORT           | 3306               | SQL port                      |
| SQL_DB             | spainews           | SQL database to use           |
| SQL_EXTRA          | parseTime=true     | Extra parameters to parse     |
| SQL_TCP            | true               | Use TCP for the connection    |
| MONGO_USER         | spainews           | MongoDB user                  |
| MONGO_PASS         | spainews           | MongoDB user password         |
| MONGO_HOST         | 127.0.0.1          | MongoDB host                  |
| MONGO_PORT         | 27017              | MongoDB port                  |
| MONGO_DB           | spainews           | MongoDB database              |
| SMTP_HOST          | 127.0.0.1          | SMTP host                     |
| SMTP_PORT          | 25                 | SMTP port                     |
| SMTP_FROM          | dashboard@spainews | SMTP from value               |
| SMTP_USERNAME      | spainews           | SMTP username                 |
| SMTP_PASSWORD      | spainews           | SMTP password                 |
| SMTP_DEV           | true               | Load SMTP in development mode |
| SMTP_DASHBOARD_URL | http://127.0.0.1   | URL to put in the emails sent |

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

Open `docs/api.openapi.yaml` with [Swagger Editor](https://editor.swagger.io/)
