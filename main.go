package main

import (
	"web-frame/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	serviceTree := app.Party("/")
	m := mvc.New(serviceTree)
	m.Handle(new(controller.ExampleController))
	app.Listen(":8080")
}
