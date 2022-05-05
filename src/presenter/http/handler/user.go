package handler

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"sake_io_api/domain/model"
	"sake_io_api/usecase"
	"strconv"
)

type UserHandler interface {
	GetUserAllHandle(c echo.Context) error
	GetUserHandle(c echo.Context) error
	RegisterUserHandle(c echo.Context) error
	UpdateUserHandle(c echo.Context) error
	DeleteUserHandle(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

func (uh userHandler) GetUserAllHandle(c echo.Context) error {
	return c.JSON(http.StatusOK, uh.userUsecase.GetUserAll())
}

func (uh userHandler) GetUserHandle(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, uh.userUsecase.GetUser(user_id))
}

func (uh userHandler) RegisterUserHandle(c echo.Context) error {
	var formData model.RegisterUserType
	if err := c.Bind(&formData); err != nil {
		logrus.Error(err)
	}
	uh.userUsecase.RegisterUser(&formData)
	return c.String(http.StatusOK, "success create user")
}

func (uh userHandler) UpdateUserHandle(c echo.Context) error {
	var formData model.UpdateUserType
	formData.User_id, _ = strconv.Atoi(c.Param("id"))
	formData.Name = c.FormValue("name")
	formData.Mail = c.FormValue("mail")
	formData.Level, _ = strconv.Atoi(c.FormValue("level"))
	formData.Pass = c.FormValue("pass")
	uh.userUsecase.UpdateUser(&formData)
	return c.String(http.StatusOK, "success update user")
}

func (uh userHandler) DeleteUserHandle(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("id"))
	uh.userUsecase.DeleteUser(user_id)
	return c.String(http.StatusOK, "success delete user")
}
