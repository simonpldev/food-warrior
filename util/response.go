package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response ...
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response200 ...
func Response200(c echo.Context, message string, data interface{}) error {
	if message == "" {
		message = "Thành công!"
	}

	return c.JSON(http.StatusOK, Response{
		Message: message,
		Data:    data,
	})
}

// Response400 ...
func Response400(c echo.Context, message string) error {
	if message == "" {
		message = "Không hợp lệ!"
	}

	return c.JSON(http.StatusBadRequest, Response{
		Message: message,
	})
}

// Response401 ...
func Response401(c echo.Context, message string) error {
	if message == "" {
		message = "Bạn không có quyền truy cập!"
	}

	return c.JSON(http.StatusUnauthorized, Response{
		Message: message,
	})
}

// Response404 ...
func Response404(c echo.Context, message string) error {
	if message == "" {
		message = "Không tìm thấy kết quả!"
	}

	return c.JSON(http.StatusNotFound, Response{
		Message: message,
	})
}
