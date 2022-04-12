package dao

import (
	"web-frame/model"

	"gorm.io/gorm"
)

type ExampleDao struct {
	db *gorm.DB
}

func NewServiceTreeDao() *ExampleDao {
	return &ExampleDao{
		db: GetDB(),
	}
}
func (s *ExampleDao) GetNodeInfo(nodeID int) (model.Example, error) {
	nodeTree := model.Example{}
	s.db.Select(model.Example{ID: nodeID}).Find(&nodeTree)
	return nodeTree, nil
}
