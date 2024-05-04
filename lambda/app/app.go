package app

import (
	"lamda-func/api"
	"lamda-func/database"
)


type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App {
		ApiHandler: apiHandler,
	}
}
