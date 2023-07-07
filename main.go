package main

import (
	"net/http"

	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/app"
	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/controller"
	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/helper"
	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/middleware"
	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/repository"
	"github.com/dickidarmawansaputra/belajar-go-RESTfull-API/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.InitDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.Router(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
