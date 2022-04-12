package service

import (
	"web-frame/dao"
	"web-frame/model"
)

type Example interface {
	GetNodeInfo(nodeID int) (model.Example, error)
}
type exampleSvc struct {
	dao *dao.ExampleDao
}

func NewExampleSvc() Example {
	return &exampleSvc{
		dao: dao.NewExampleDao(),
	}
}
func (s *exampleSvc) GetNodeInfo(nodeID int) (model.Example, error) {
	return s.dao.GetNodeInfo(nodeID)
}
