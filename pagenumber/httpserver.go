package pagenumber

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil && pageStr != "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    "BAD_REQUEST_PARAMS",
				"message": "page parameter is invalid, should be positive integer",
			})
		}

		fetchParam := FetchParam{
			PageNumber: uint64(page),
		}

		if pageStr == "" {
			fetchParam.PageNumber = 1
		}

		res, nextpage, err := FetchPayment(c.Request().Context(), db, fetchParam)
		if err != nil {
			return err
		}

		c.Response().Header().Add("X-NextPage", fmt.Sprintf("%d", nextpage))
		return c.JSON(http.StatusOK, res)
	}
}
