package main

import (
	"context"
	"fmt"

	"goscrum/server/db"
	"goscrum/server/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(_ context.Context) {
	db := db.DbClient(true)
	defer db.Close()

	err := db.AutoMigrate(&models.Project{}, &models.Participant{}, &models.Question{}, &models.Workspace{}, &models.Answer{}).Error

	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Request Initiated")
	lambda.Start(HandleRequest)
}
