package service

import (
	"github.com/Scfy-Code/IM/app/mapper"
)

// IndexService 首页服务接口
type IndexService interface {
	GetTalkerList(selfID string) []map[string]interface{}
	GetTeamList(selfID string) []map[string]interface{}
}
type indexServiceImpl struct {
	indexMapper mapper.IndexMapper
}

func (isi indexServiceImpl) GetTalkerList(selfID string) []map[string]interface{} {
	return isi.indexMapper.GetTalkerList(selfID)
}
func (isi indexServiceImpl) GetTeamList(selfID string) []map[string]interface{} {
	return isi.indexMapper.GetTeamList(selfID)
}

// NewIndexService 创建首页服务实例
func NewIndexService() IndexService {
	return indexServiceImpl{
		mapper.NewIndexMapper(),
	}
}
