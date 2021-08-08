# taptalk-diary
Tap Talk Diary Code Testing

## core tech stacks
- Golang 1.16
- Go fiber Version 2.16
- Gorm
- Swagger for API Documentation

## How to Use This API
# first time setup
before that, please copy from .env.example to .env
# then
$ go mod tidy
# if there is any changed with swagger
$ swag init 
# run the api
$ go run main.go
OR 
$ go build <Your Executed file>
# to open the swagger 
localhost:<yourport>/swagger/

## FYI
# for username can consists email or actual username 
# there are email validation and birthday validation format (described at swagger)
# for doc details are described at swagger
