package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"sake_io_api/domain/model"
	"sake_io_api/usecase"
	"strconv"
	"time"
)

type ManagementHandler interface {
	GetManagementAllHandle(c echo.Context) error
	GetManagementHandle(c echo.Context) error
	RegisterManagementHandle(c echo.Context) error
	UpdateManagementHandle(c echo.Context) error
	DeleteManagementHandle(c echo.Context) error
}

type managementHandler struct {
	managementUsecase usecase.ManagementUsecase
}

func NewManagementHandler(mu usecase.ManagementUsecase) ManagementHandler {
	return &managementHandler{
		managementUsecase: mu,
	}
}

func (mh managementHandler) GetManagementAllHandle(c echo.Context) error {
	return c.JSON(http.StatusOK, mh.managementUsecase.GetManagementAll())
}

func (mh managementHandler) GetManagementHandle(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, mh.managementUsecase.GetManagement(user_id))
}

func (mh managementHandler) RegisterManagementHandle(c echo.Context) error {
	var formData model.RegisterManagementType
	formData.User_id, _ = strconv.Atoi(c.Param("id"))
	formData.Sake_name = c.FormValue("sake_name")
	formData.Amount, _ = strconv.Atoi(c.FormValue("amount"))
	formData.Date, _ = time.Parse("2006-01-02", c.FormValue("date"))
	mh.managementUsecase.RegisterManagement(&formData)
	return c.String(http.StatusOK, "success create management")
}

func (mh managementHandler) UpdateManagementHandle(c echo.Context) error {
	var formData model.UpdateManagementType
	formData.Management_id, _ = strconv.Atoi(c.Param("id"))
	formData.Amount, _ = strconv.Atoi(c.FormValue("amount"))
	mh.managementUsecase.UpdateManagement(&formData)
	return c.String(http.StatusOK, "success update management")
}

func (mh managementHandler) DeleteManagementHandle(c echo.Context) error {
	management_id, _ := strconv.Atoi(c.Param("id"))
	mh.managementUsecase.DeleteManagement(management_id)
	return c.String(http.StatusOK, "success delete management")
}
