package service

import (
	"strconv"

	"github/ranryl/web-frame/dao"
	"github/ranryl/web-frame/model"

	"gorm.io/gorm"
)

type ExampleSvc interface {
	GetNodeInfo(nodeID uint) (model.Example, error)
	InsertNode(parentID uint, nodeName string) error
	GetOneChilds(nodeID uint) ([]model.Example, error)
	GetAllChilds(nodeID uint) ([]model.Example, error)
}
type serviceTreeSvc struct {
	dao dao.ExampleDao
}

func NewServiceTreeSvc() ExampleSvc {
	return &serviceTreeSvc{
		dao: *dao.NewExampleDao(),
	}
}
func (s *serviceTreeSvc) GetNodeInfo(nodeID uint) (model.Example, error) {
	return s.dao.GetNodeInfo(nodeID)
}

func (s *serviceTreeSvc) GetOneChilds(nodeID uint) ([]model.Example, error) {
	data, err := s.dao.GetNodeInfo(nodeID)
	if err != nil {
		return nil, err
	}
	return s.dao.GetOneChildsByPath(data.Path)
}

func (s *serviceTreeSvc) GetAllChilds(nodeID uint) ([]model.Example, error) {
	data, err := s.dao.GetNodeInfo(nodeID)
	if err != nil {
		return nil, err
	}
	return s.dao.GetAllChildsByPath(data.Path)
}

func (s *serviceTreeSvc) InsertNode(parentID uint, nodeName string) error {
	return s.dao.Transaction(func(tx *gorm.DB) error {
		nodeTree := model.Example{
			NodeName: nodeName,
		}
		result := tx.Create(&nodeTree)
		if result.Error != nil {
			return result.Error
		}
		parentData := model.Example{}
		parentResult := tx.Find(&parentData, gorm.Model{ID: parentID})
		if parentResult.Error != nil {
			return result.Error
		}
		nodeTree.Path = parentData.Path + "/" + strconv.Itoa(int(nodeTree.ID))
		result = tx.Model(&nodeTree).Where("id = ?", nodeTree.ID).Update("path", nodeTree.Path)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
