package main

import (
	"AlifTask/api"
	"AlifTask/base"
	"AlifTask/models"
	"fmt"
	"net/http"
)

func main() {
	if err := base.Connect(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Database connected!")
	if err := models.Init(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Models initialized!")
	http.ListenAndServe(":2001", api.MainRouter())
}
