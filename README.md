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

### Go commands
- atualizar o go.mod:
```bash
go mod tidy
```

### MongoDB
```bash
docker run --name mongodb -d -p 27017:27017 mongo
```

###
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README.md)