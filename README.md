# spAInews

## Deploying

Make sure you have `docker` and `docker-cli` installed (`docker desktop` should do the job on Windows).

### Production

```shell
TODO: DOCUMENT THIS
```

### Development

#### Frontend development:

1. First install dependencies:

```shell
cd frontend
npm install
cd ..
```

2. Start the APIs:

```shell
./dev-apis.ps1 # This command may require admin rights
```

3. Start the frontend:

```shell
cd frontend
npm start
```

#### Backend development:

1. First start the databases services:

```shell
./dev-databases.ps1 # This command may require admin rights, it may require you to wait some second until databases are initialized
```

2. Start auth API:

```shell
./start-auth-api.ps1
```

3. Start news API:

```shell
./start-news-api.ps1
```

4. Unit testing the APIs can be done by executing:

```shell
go test ./...
```

