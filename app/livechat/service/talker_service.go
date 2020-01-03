package service

import "github.com/Scfy-Code/IM/app/livechat/mapper"

// TalkerService 好友数据层服务接口
type TalkerService interface {
	CreateTalker(talkerID string) bool
	DeleteTalker(bindID string) bool
	UpdateTalker(talkerInfo map[string]interface{}) bool
	SelectTalker(talkerID string) map[string]interface{}
	SelectTalkers(selfID string) []map[string]interface{}
}
type talkerServiceImpl struct {
	talkerMapper mapper.TalkerMapper
}

func newTalkerServiceImpl() TalkerService {
	return &talkerServiceImpl{
		mapper.NewTalkerMapper("talkerMapper"),
	}
}
func (tsi talkerServiceImpl) CreateTalker(talkerID string) bool {
	return false
}
func (tsi talkerServiceImpl) DeleteTalker(bindID string) bool {
	return tsi.talkerMapper.DeleteTalker(bindID)
}
func (tsi talkerServiceImpl) UpdateTalker(talkerInfo map[string]interface{}) bool {
	return false
}
func (tsi talkerServiceImpl) SelectTalker(talkerID string) map[string]interface{} {
	return tsi.talkerMapper.SelectTalker(talkerID)
}
func (tsi talkerServiceImpl) SelectTalkers(selfID string) []map[string]interface{} {
	return tsi.talkerMapper.SelectTalkers(selfID)
}

// NewTalkerService 新建一个好友数据层服务对象
func NewTalkerService(serviceName string) TalkerService {
	switch serviceName {
	case "talkerService":
		return newTalkerServiceImpl()
	default:
		return newTalkerServiceImpl()
	}
}
