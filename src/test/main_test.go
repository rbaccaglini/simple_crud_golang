package tests

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
	"github.com/rbaccaglini/simple_crud_golang/src/model/repository"
	"github.com/rbaccaglini/simple_crud_golang/src/model/service"
	"github.com/rbaccaglini/simple_crud_golang/src/test/connection"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserController  controller.UserControllerInterface
	Database        *mongo.Database
	closeConnection func()
)

func TestMain(m *testing.M) {
	if err := os.Setenv("MONGODB_USER_DB", "test_users"); err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	if err := os.Setenv("MONGODB_USER_DB_COLLECTION", "test_users"); err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	if err := os.Setenv("JWT_SECRET_KEY", "123456"); err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	code := m.Run()

	// Ensure closeConnection is called after all tests
	log.Println("Clearing environment variables and closing Docker connection")
	os.Clearenv()
	if closeConnection != nil {
		closeConnection()
	}

	os.Exit(code)
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}

func VerifyToken(tokenValue string) bool {
	secret := os.Getenv("JWT_SECRET_KEY")
	t, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("invalid token")
	})
	if err != nil {
		return false
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return false
	}

	return true
}
