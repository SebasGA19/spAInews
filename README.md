# spAInews

[![Test build and deploy](https://github.com/SebasGA19/spAInews/actions/workflows/pipeline.yaml/badge.svg)](https://github.com/SebasGA19/spAInews/actions/workflows/pipeline.yaml)

## Deploying

Make sure you have `docker` and `docker-cli` installed (`docker desktop` should do the job on Windows).

### Production

It is important to check if the predefined configurations inside `configs/` directory are consistent with your deployment

```shell
docker-compose -f ./docker-compose.yaml up -d --build
```

### Development

#### Frontend development:

1. Setup the API

```shell
docker-compose -f ./compose-dev-api.yaml up -d --build
```

2. First install dependencies:

```shell
cd frontend
npm install
cd ..
```

3. Start the frontend:

```shell
cd frontend
npm run dev
```

#### Backend development:

Start the databases services:

```shell
docker-compose -f ./compose-dev-databases.yaml up -d --build
```

