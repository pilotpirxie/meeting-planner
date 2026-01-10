# hangout-planner

A simple and open source hangout planner.

## How to run

```sh
# backend server
cd backend
go run ./cmd/api

# to run client in a separate terminal
cd client
yarn dev
```

## How to build client

```sh
cd client
yarn build:copy
```

## How to prepare migrations
```sh
# https://github.com/pressly/goose
brew install goose
goose up
```