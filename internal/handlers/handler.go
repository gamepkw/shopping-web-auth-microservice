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
	Message string `json:"error_message"`
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
	var request model.RegisterRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid request"})
	}

	if request.Username == "" {
		logrus.Errorf("[Register] Invalid Username")
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid Username"})
	}

	if request.Password == "" {
		logrus.Errorf("[Register] Empty Password")
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Empty Password"})
	}

	if len(request.Password) < 8 {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Password length must be at least 8 characters"})
	}

	ctx := c.Request().Context()
	if err := a.authService.Register(ctx, request); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, Response{Message: "Register successful", Body: nil})
}

func (a *AuthHandler) Login(c echo.Context) error {
	var loginRequest model.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid request"})
	}

	if loginRequest.Username == "" {
		logrus.Errorf("[Register] Invalid Username")
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Invalid Username"})
	}

	if loginRequest.Password == "" {
		logrus.Errorf("[Register] Empty Password")
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Empty Password"})
	}

	ctx := c.Request().Context()
	token, err := a.authService.Login(ctx, loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: err.Error(), Body: nil})
	}

	return c.JSON(http.StatusCreated, Response{Message: "Login successful", Body: model.LoginResponse{Token: token}})
}

var TimestampFormat = "2006-01-02 15:04:05"
