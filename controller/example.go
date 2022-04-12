package controller

import (
	"github.com/kataras/iris/v12"
)

type ExampleController struct {
	XXX string
}

func (s *ExampleController) Get(ctx iris.Context) {
	ctx.JSON(ExampleController{
		XXX: "TableController",
	})
	// ex := service.NewExampleSvc()
	// node, err := ex.GetNodeInfo(1)
	// if err != nil {
	// 	log.Errorf("ExampleController getNodeInfo error: %w\n", err)
	// }
	// ctx.JSON(node)
}
