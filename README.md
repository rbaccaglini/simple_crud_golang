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
### Mockgen
- installing mockgen
```bash
go install go.uber.org/mock/mockgen@latest
```

- creating user repository mock
```bash
mockgen -source=internal/repositories/user/user_repository_interface.go -destination=test/mocks/user_repository_interface_mock.go -package=mocks
```

- creating user domain mock
```bash
mockgen -source=src/model/user_domain_interface.go -destination=src/test/mocks/user_domain_interface_mock.go -package=mocks
```

- creating user domain service mock
```bash
mockgen -source=src/model/service/user_interface.go -destination=src/test/mocks/user_interface_mock.go -package=mocks
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
