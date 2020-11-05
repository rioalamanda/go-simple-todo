package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-simple-todo/controller"
	"go-simple-todo/model"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	// log.Fatal(http.ListenAndServe(":3000", mux))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
