package result

import (
	"github.com/labstack/echo"
	"net/http"
)

const (
	SUCCESS int = 1
	FAILED  int = 0
)

func Success(ctx echo.Context,message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  message,
		"data": data,
	})
}

func Failed(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  data,
	})
}
