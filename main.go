package main

import (
	"web-frame/controller"
	"web-frame/util"

	"github.com/kataras/iris/v12"
)

func main() {
	util.InitConfig()
	app := iris.New()
	stApp := app.Party("/")
	example := &controller.ExampleController{}
	stApp.Get("/", example.Get)
	app.Listen(":8080")
}
