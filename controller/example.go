package controller

import (
	"net/http"

	"github/ranryl/web-frame/model"
	"github/ranryl/web-frame/service"

	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
)

// swagger:response responseData
type ResponseDataWrapper struct {
	// in: body
	Body responseTreeData
}
type responseTreeData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// model.Example
	Data []model.Example `json:"data"`
}

// swagger:response defaultResponse
type DefaultResponseWrapper struct {
	// in : body
	Body defaultResponse
}

type defaultResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// swagger:parameters addNode
// in: body
// example: {id:1, nodeName: "NIO"}
// required: true
type nodeData struct {
	ParentID uint   `json:"parent_id"`
	NodeName string `json:"node_name"`
}
type ExampleController struct {
	stService service.ExampleSvc
}

func NewExampleController() *ExampleController {
	return &ExampleController{
		stService: service.NewServiceTreeSvc(),
	}
}

// swagger:operation GET /getNode node getNode
// ---
// summary: get node by nodeID
// description:
// parameters:
// - name: id
//   in: query
//   description: node id
//   required: true
//   type: integer
//   format: uint32
// responses:
//   "200":
//     "$ref": "#/responses/responseData"
//   "400":
//     "$ref": "#/responses/defaultResponse"
func (s *ExampleController) GetNode(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil || id <= 0 {
		if err != nil {
			log.Errorf("get id error: %w", err)
		}
		ctx.StopWithJSON(http.StatusOK, defaultResponse{
			Code: http.StatusBadRequest,
			Msg:  "id is invalid",
		})
		return
	}
	node, err := s.stService.GetNodeInfo(uint(id))
	if err != nil {
		log.Errorf("ServiceTreeController getNodeInfo error: %w\n", err)
	}
	ctx.JSON(responseTreeData{
		Code: http.StatusOK, Msg: http.StatusText(http.StatusOK),
		Data: []model.Example{node},
	})
}

// swagger:operation POST /AddNode node addNode
// ---
// produces:
// - application/json
// summary: add node
// description:
// responses:
//   "200":
//     "$ref": "#/responses/defaultResponse"
//   "400":
//     "$ref": "#/responses/defaultResponse"
func (s *ExampleController) AddNode(ctx iris.Context) {
	var node nodeData
	if err := ctx.ReadBody(&node); err != nil {
		bodyData, err := ctx.GetBody()
		if err != nil {
			log.Error("get body data error: %w", err)
		}
		ctx.StopWithJSON(http.StatusOK, defaultResponse{
			Code: http.StatusBadRequest,
			Msg:  "get request data error: " + string(bodyData),
		})
		return
	}
	if node.ParentID == 0 || node.NodeName == "" {
		ctx.StopWithJSON(http.StatusOK, defaultResponse{
			Code: http.StatusBadRequest,
			Msg:  "parentID or nodeName can't by empty",
		})
		return
	}
	err := s.stService.InsertNode(node.ParentID, node.NodeName)
	if err != nil {
		log.Errorf("AddNode error: %w", err)
		ctx.StopWithJSON(http.StatusOK, defaultResponse{
			Code: http.StatusBadRequest,
			Msg:  "AddNode error",
		})
		return
	}
	ctx.JSON(defaultResponse{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK)})
}
