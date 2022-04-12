package controller

import "github.com/kataras/iris/v12"

type ExampleController struct {
	XXX string
}

func (s *ExampleController) Get(ctx iris.Context) {
	ctx.JSON(ExampleController{
		XXX: "TableController",
	})
}
