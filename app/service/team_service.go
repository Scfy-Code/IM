package service

import "github.com/Scfy-Code/IM/app/mapper"

// TeamService 群组数据层服务接口
type TeamService interface {
	CreateTeam(teamID string) bool
	DeleteTeam(teamID string) bool
	UpdateTeam(teamInfo map[string]interface{}) bool
	SelectTeam(talkerID string) map[string]interface{}
	SelectTeams(selfID string) []map[string]interface{}
}
type teamServiceImpl struct {
	teamMapper mapper.TeamMapper
}

func (tsi teamServiceImpl) CreateTeam(teamID string) bool {
	return false
}
func (tsi teamServiceImpl) DeleteTeam(teamID string) bool {
	return false
}
func (tsi teamServiceImpl) UpdateTeam(teamInfo map[string]interface{}) bool {
	return false
}
func (tsi teamServiceImpl) SelectTeam(talkerID string) map[string]interface{} {
	return tsi.teamMapper.SelectTeam(talkerID)
}
func (tsi teamServiceImpl) SelectTeams(selfID string) []map[string]interface{} {
	return tsi.teamMapper.SelectTeams(selfID)
}

// NewTeamService 新建群组数据层服务对象
func NewTeamService() TeamService {
	return teamServiceImpl{
		mapper.NewTeamMapper(),
	}
}
