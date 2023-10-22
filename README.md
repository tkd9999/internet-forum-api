# internet-forum-api
インターネット掲示板RestAPI(Go/Echo)をクリーンアーキテクチャで実装（詳細について記事作成予定）

## Architecture
```
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── db
│   └── db.go
├── migrate
│   └── migrate.go
├── controller
│   ├── post_controller.go
│   ├── thread_controller.go
│   └── user_controller.go
├── models
│   ├── post.go
│   ├── thread.go
│   └── user.go
├── repository
│   ├── post_repository.go
│   ├── thread_repository.go
│   └── user_repository.go
├── router
│   └── router.go
├── usecase
│   ├── post_usecase.go
│   ├── thread_usecase.go
│   └── user_usecase.go
└── validator
    ├── post_validator.go
    ├── thread_validator.go
    └── user_validator.go
```

## Usage
```
#create module
go mod init github.com/junshintakeda/internet-forum
#start db
docker compose up -d
#migration
GO_ENV=dev go run ./migrate/migrate.go
#run app
GO_ENV=dev go run main.go
```

## Refernces
1. https://github.com/GomaGoma676/echo-rest-api
2. https://github.com/eldimious/golang-api-showcase
