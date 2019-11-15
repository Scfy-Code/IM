package service

import "github.com/Scfy-Code/IM/sys/mapper"

// TeamService 群组数据层服务接口
type TeamService interface {
	CreateTeam(string) bool
	DeleteTeam(string) bool
	UpdateTeam(map[string]interface{}) bool
	SelectTeam(string) map[string]interface{}
	SelectTeams(string) []map[string]interface{}
}
type teamServiceImpl struct {
	teamMapper mapper.TeamMapper
}

func newTeamServiceImpl() TeamService {
	return &teamServiceImpl{
		mapper.NewTeamMapper("teamMapper"),
	}
}

func (tsi teamServiceImpl) CreateTeam(teamID string) bool {
	return false
}
func (tsi teamServiceImpl) DeleteTeam(bindID string) bool {
	return tsi.teamMapper.DeleteTeam(bindID)
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
func NewTeamService(serviceName string) TeamService {
	switch serviceName {
	case "teamService":
		return newTeamServiceImpl()
	default:
		return newTeamServiceImpl()
	}
}
