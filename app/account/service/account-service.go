package service

import "github.com/Scfy-Code/IM/app/account/mapper"

type accountService struct {
	accountMapper mapper.AccountMapper
}

// NewAccountService 新建账号服务
func NewAccountService() AccountService {
	return accountService{
		mapper.NewAccountMapper(),
	}
}
func (as accountService) SelectAccount(email, password string) bool {
	return false
}
