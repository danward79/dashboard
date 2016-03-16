package controllers

import (
	"log"

	"github.com/danward79/dashboard/models"
)

//Env ...
type Env struct {
	DB *models.DB
}

//Init ...
func Init(f string) *Env {

	db, err := models.NewDB(f)
	if err != nil {
		log.Panic(err)
	}

	return &Env{db}
}
