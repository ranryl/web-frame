// Package classification Service Tree API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
// 	   Version: 0.0.1
// 	   Host: localhost:8080
// 	   BasePath: /node
// 	   Schemes: http, https
// 	   Contact: ruilong.ran@nio.com
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     Security:
//     - api_key: xxxx
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
// swagger:meta
package main

import (
	"github/ranryl/web-frame/controller"
	"github/ranryl/web-frame/util"

	"github.com/kataras/iris/v12"
)

func main() {
	util.InitConfig()
	app := iris.Default()
	app.Get("/", index)
	app.HandleDir("/swagger", iris.Dir("./docs"))
	exampleApp := app.Party("/node")
	ex := controller.NewExampleController()
	exampleApp.Get("/get", ex.GetNode)
	exampleApp.Post("/add", ex.AddNode)
	app.Listen(":8080")
}
func index(ctx iris.Context) {
	ctx.HTML("hello world")
}
