package dao

import (
	"github/ranryl/web-frame/model"

	"gorm.io/gorm"
)

type ExampleDao struct {
	*gorm.DB
}

func NewExampleDao() *ExampleDao {
	return &ExampleDao{
		GetDB(),
	}
}
func (s *ExampleDao) GetNodeInfo(nodeID uint) (model.Example, error) {
	nodeTree := model.Example{}
	result := s.Find(&nodeTree, gorm.Model{ID: nodeID})
	return nodeTree, result.Error
}

// GetOneChilds 获取当前节点的直系子节点（包含当前节点）
func (s *ExampleDao) GetOneChildsByPath(path string) ([]model.Example, error) {
	data := make([]model.Example, 0)
	result := s.Where("path = ? or path regexp concat(?, '/', '[0-9]$')", path, path).Find(&data)
	return data, result.Error
}

// GetAllChildsByPath 获取当前节点的所有子节点 （包含当前节点）
func (s *ExampleDao) GetAllChildsByPath(path string) ([]model.Example, error) {
	data := make([]model.Example, 0)
	result := s.Where("path like ?", path+"%").Find(&data)
	return data, result.Error
}

func (s *ExampleDao) InsertNode(path, nodeName string) (uint, error) {
	nodeTree := model.Example{
		Path:     path,
		NodeName: nodeName,
	}
	result := s.Create(&nodeTree)
	return nodeTree.ID, result.Error
}
