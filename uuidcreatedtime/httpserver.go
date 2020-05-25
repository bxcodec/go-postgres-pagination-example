package uuidcreatedtime

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		limitStr := c.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil && limitStr != "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    "BAD_REQUEST_PARAMS",
				"message": "limit parameter is invalid, should be positive integer",
			})
		}
		if limit == 0 {
			limit = 10
		}
		cursor := c.QueryParam("cursor")
		fetchParam := FetchParam{
			Limit:  uint64(limit),
			Cursor: cursor,
		}

		res, nextCursor, err := FetchPayment(c.Request().Context(), db, fetchParam)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": err.Error(),
			})
		}

		c.Response().Header().Add("X-Cursor", nextCursor)
		return c.JSON(http.StatusOK, res)
	}
}
