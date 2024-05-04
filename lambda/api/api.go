package api

import (
	"fmt"
	"lamda-func/database"
	"lamda-func/types"
)


type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler {
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("Request has empty parameters")
	}

	userExists, err := api.dbStore.DoesUserExist(event.Username)

	if err != nil {
		return fmt.Errorf("DynamoDB error: %w", err)
	}

	if userExists {
		return fmt.Errorf("User with username: %s already exists", event.Username) 
	}

	err = api.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("Error registering user: %w", err)
	}

	return nil
}
