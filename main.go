package main

import (
	"golang_cms/configs"
	"golang_cms/routes"
)

func main() {

	//run database
	configs.ConnectDB()

	r := routes.BannerRoute()

	r.Run()
}
