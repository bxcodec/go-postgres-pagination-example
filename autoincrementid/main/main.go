package main

import (
	pagination "github.com/bxcodec/go-postgres-pagination-example"
	"github.com/bxcodec/go-postgres-pagination-example/autoincrementid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	db := pagination.InitDB()
	fetchHandler := autoincrementid.FetchHandler(db)
	e := echo.New()
	e.GET("/payments", fetchHandler)
	logrus.Error(e.Start(":9090"))
}
