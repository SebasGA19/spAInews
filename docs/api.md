# API

This documentation container everything related to the .

## Config

To configure the API client it is important to specify this environment variables:

| Environment variable name | Example value                                    | Description                                                                                                                         |
|---------------------------|--------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| MONGO_URL                 | mongodb://spainews:spainews@mongo:27017/spainews | This is address that the application is going to use for listening and accepting incoming connections                               |
| LISTEN_ADDRESS            | 0.0.0.0:5000                                     | This is the URL specified in [The library documentation](https://github.com/mongodb/mongo-go-driver) used to connect to the MongoDB |

## Routes

...