package handler

import (
	"net/http"

	model "github.com/gamepkw/shopping-web-auth-microservice/internal/models"
	authService "github.com/gamepkw/shopping-web-auth-microservice/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService authService.AuthService
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

func NewAuthHandler(e *echo.Echo, as authService.AuthService) {
	handler := &AuthHandler{
		authService: as,
	}

	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)
}

func (a *AuthHandler) Register(c echo.Context) error {
	var registerRequest model.LoginRequest

	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if registerRequest.Username == "" {
		logrus.Errorf("[Register] Invalid Username")
		return c.JSON(http.StatusBadRequest, Response{Message: "Invalid Username", Body: nil})
	}

	if registerRequest.Password == "" {
		logrus.Errorf("[Register] Empty Password")
		return c.JSON(http.StatusBadRequest, Response{Message: "Empty Password", Body: nil})
	}

	return c.JSON(http.StatusCreated, Response{Message: "Register successful", Body: nil})
}

func (a *AuthHandler) Login(c echo.Context) error {
	var loginRequest model.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if loginRequest.Username == "" {
		logrus.Errorf("[Register] Invalid Username")
		return c.JSON(http.StatusBadRequest, Response{Message: "Invalid Username", Body: nil})
	}

	if loginRequest.Password == "" {
		logrus.Errorf("[Register] Empty Password")
		return c.JSON(http.StatusBadRequest, Response{Message: "Empty Password", Body: nil})
	}

	return c.JSON(http.StatusCreated, Response{Message: "Login successful", Body: nil})
}

var TimestampFormat = "2006-01-02 15:04:05"

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
