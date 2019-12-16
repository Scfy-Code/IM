package service

import (
	"github.com/Scfy-Code/IM/app/account/mapper"
	"github.com/Scfy-Code/IM/sys"
)

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
	account := as.accountMapper.SelectAccount(email, password)
	sys.InfoLogger.Println(account)
	return false
}
