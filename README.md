# Simple Crud using Golang
### GoLang
- version go1.23.4

### go.mod
```bash
go mod init github.com/rbaccaglini/simple_crud_golang
```

### Packeges
- go get github.com/joho/godotenv
- go get -u github.com/gin-gonic/gin
- go get github.com/go-playground/validator/v10
- go get -u go.uber.org/zap
- go get go.mongodb.org/mongo-driver/mongo
- go get github.com/google/uuid
- go get -u github.com/golang-jwt/jwt
- go get -u github.com/ory/dockertest/v3

### Go commands
- atualizar o go.mod:
```bash
go mod tidy
```

### MongoDB
```bash
docker run --name mongodb -d -p 27017:27017 mongo
```
### Example of user registration in DB:
```json
{
    "_id": {
        "$oid": "67a260ca7847df2a9696aa3a"
    },
    "password": "54df418471fded5c07f2d338241ba202",
    "email": "test@test.com",
    "name": "Roger",
    "age": 47
}
```

### Mockgen
- installing mockgen
```bash
go install go.uber.org/mock/mockgen@latest
```

- creating user repository mock
```bash
mockgen -source=internal/repositories/user/user_repository_interface.go -destination=test/mocks/user_repository_interface_mock.go -package=mocks
```

- creating user service mock
```bash
mockgen -source=internal/services/user/user_service_interface.go -destination=test/mocks/user_service_interface_mock.go -package=mocks
```

- creating user handler mock
```bash
mockgen -source=internal/handlers/user/user_handler_interface.go -destination=test/mocks/user_handler_interface_mock.go -package=mocks
```

### Uteis
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README.md)
- Commands
```bash
docker container stop $(docker ps -a -q)
docker container rm $(docker ps -a -q)
docker image rm meuprimeirocrudgo
```

### Docker
- Build image
```bash
docker build -t meuprimeirocrudgo .
```

- Run container
```bash
docker container run -d -p 8080:8080 meuprimeirocrudgo
```
