package dao

import (
	"web-frame/model"

	"gorm.io/gorm"
)

type ExampleDao struct {
	db *gorm.DB
}

func NewExampleDao() *ExampleDao {
	return &ExampleDao{
		db: GetDB(),
	}
}
func (s *ExampleDao) GetNodeInfo(nodeID int) (model.Example, error) {
	nodeTree := model.Example{}
	s.db.Find(&nodeTree, model.Example{ID: nodeID})
	return nodeTree, nil
}
