package rest

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

// for returning JSON bodies
type H map[string]interface{}

// returns the http status code found in the error message
func parseError(c echo.Context, err error) error {
	// if error came from vault, relay it
	errCode := strings.Split(err.Error(), "Code:")
	errMsgs := strings.Split(err.Error(), "*")
	if len(errCode) > 1 && len(errMsgs) > 1 {
		code := 400
		fmt.Sscanf(errCode[1], "%d", &code)
		return c.JSON(code, H{
			"error": "Vault: " + errMsgs[1],
		})
	}

	return c.JSON(http.StatusBadRequest, H{
		"error": err.Error(),
	})
}
