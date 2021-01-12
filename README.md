# Social network

## Introduction:

> The simple social network web application project which using Golang, React and GraphQL

## How it works?

> Require: Go version 1.13, Node version 1.13

### Server:

- Environments (default):
  ```shell
  PORT=8080
  ENV=local
  LOG_LEVEL=INFO
  LOG_PATH=
  JWT_KEY=secret
  DB_URI=mongodb://localhost/social-network
  ```

- Run the server:
  ```shell
  $ make run-server
  ```
  or
  ```shell
  $ go run cmd/main.go
  ```

### Web client:

- Environments (default):
  ```shell
  REACT_APP_BASE_URL=http://localhost:8080      (server URL)
  ```

- Run the web client:
  ```shell
  $ make run-client
  ```
  or
  ```shell
  $ cd web/
  $ yarn start
  ```

Run all project with Docker: (*coming soon...*)