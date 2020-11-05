package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/rioalamanda/go-simple-todo/controller"
	"github.com/rioalamanda/go-simple-todo/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
