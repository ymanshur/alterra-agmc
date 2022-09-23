## GO RESTFUL APP

The day 6 & 7 submission AGMC (by [Alterra](https://www.alterra.id/)) project.

Notable features:

- Includes a multi-stage Dockerfile, which actually build Go binaries (minimum app size).
- Has functional tests for application's business requirements using default Go testing.
- You need to run database app (MySQL) in your local machine first before up the app docker image

### Setup
1. Create `.env` config file, look at [.env.example](./.env.example) for mandatory *key-value*
2. Build app using `docker`

    ```bash
    docker build -t <your_tag> .
    ```
    or using `docker-compose`

    ```bash
    docker-compose build
    ```
    or just *up* published docker image (means [Docker Hub](https://hub.docker.com/)) using following `docker-compose.yml` file

    ```
    version: "3.7"

    services:
    app:
        image: ymanshur/go-restful:${APP_VERSION}
        container_name: go-restful
        ports:
            - "${APP_PORT:-8080}:8080"
        environment:
            JWT_SECRET: ${JWT_SECRET:?err}
            DB_USER: ${DB_USER:?err}
            DB_PASS: ${DB_PASS}
            DB_PORT: ${DB_PORT:?err}
            DB_HOST: host.docker.internal
            DB_NAME: ${DB_NAME:?err}
    ```

2. Run app using docker-compose

    ```bash
    docker-compose up -d
    ```
    <small>Note: `-d` argument means detached mode</small>

### Testing

- Do functional test using following command

    ```bash
    make test
    ```
- To show in html mode, use

    ```bash
    make test mode=html
    ```