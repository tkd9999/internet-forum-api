#internet-forum-api
インターネット掲示板RestAPI(Go/Echo)をクリーンアーキテクチャで実装

##Usage
'''
#create module
go mod init github.com/junshintakeda/internet-forum
#start db
docker compose up -d
#migration
GO_ENV=dev go run ./migrate/migrate.go
#run app
GO_ENV=dev go run main.go
'''

##Refernces
https://github.com/GomaGoma676/echo-rest-api
https://github.com/eldimious/golang-api-showcase
