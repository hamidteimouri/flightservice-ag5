package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseUnprocessableEntity(c echo.Context, resp Response) error {
	return c.JSON(http.StatusUnprocessableEntity, resp)
}

func ResponseOK(c echo.Context, resp Response) error {
	return c.JSON(http.StatusOK, resp)
}

func ResponseNotFound(c echo.Context, resp Response) error {
	return c.JSON(http.StatusNotFound, resp)
}

func ResponseUnauthorized(c echo.Context, resp Response) error {
	return c.JSON(http.StatusUnauthorized, resp)
}

func ResponseInternalError(c echo.Context, resp Response) error {
	return c.JSON(http.StatusInternalServerError, resp)
}
