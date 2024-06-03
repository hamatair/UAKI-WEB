package middleware

import (
	"UAKI-WEB/entity"
	"UAKI-WEB/internal/service"
	"UAKI-WEB/model"
	"UAKI-WEB/pkg/jwt"
	"UAKI-WEB/pkg/response"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	TimeoutMiddleware() gin.HandlerFunc
	AuthenticateUser(c *gin.Context)
	Authorization(c *gin.Context)
}

type Middleware struct {
	jwtauth jwt.Interface
	service *service.Service
}

func Init(jwtauth jwt.Interface, service *service.Service) Interface {
	return &Middleware{
		jwtauth: jwtauth,
		service: service,
	}
}

// AuthenticateUser implements Interface.
func (m *Middleware) AuthenticateUser(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		response.Error(c, http.StatusUnauthorized, "empty token", errors.New(""))
		c.Abort()
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtauth.ValidateToken(token)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "failed validate token", err)
		c.Abort()
	}
	
	user, err := m.service.UserService.GetUser(model.UserParam{
		ID: userId,
	})
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "failed get user", err)
		c.Abort()
	}

	c.Set("user", user)

	c.Next()}

// PremiumAccess implements Interface.
func (m *Middleware) Authorization(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		response.Error(c, http.StatusForbidden, "failed to authorize user", errors.New(""))
		c.Abort()
	}

	if user.(entity.User).Role != 1 {
		response.Error(c, http.StatusForbidden, "failed to let user", errors.New("user don't have access"))
		c.Abort()
	}

	c.Next()}

func (m *Middleware) TimeoutMiddleware() gin.HandlerFunc {
	timeOut, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeOut)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func testResponse(c *gin.Context) {
	response.Error(c, http.StatusRequestTimeout, "Time Out Limit", errors.New(""))
}
