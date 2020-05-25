package autoincrementid

import (
	"database/sql"
	"fmt"
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
				"message": "limit parameter is invalid",
			})
		}
		if limit == 0 {
			limit = 10
		}

		cursorStr := c.QueryParam("cursor")
		cursor, err := strconv.Atoi(cursorStr)
		if err != nil && cursorStr != "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    "BAD_REQUEST_PARAMS",
				"message": "cursor parameter is invalid",
			})
		}

		fetchParam := FetchParam{
			Limit:    uint64(limit),
			CursorID: uint64(cursor),
		}

		res, nextCursor, err := FetchPayment(c.Request().Context(), db, fetchParam)
		if err != nil {
			return err
		}

		c.Response().Header().Add("X-NextCursor", fmt.Sprintf("%d", nextCursor))
		return c.JSON(http.StatusOK, res)
	}
}
