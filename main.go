package main

import (
	"belajar_golang_dependency_injection/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//err := server.ListenAndServe()
	//helper.PanicIfError(err)
}
