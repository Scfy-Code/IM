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
	account, err := as.accountMapper.SelectAccount(email, password)
	sys.WarnLogger.Println(account, err.Error())
	return false
}
func (as accountService) InsertAccount(email, password, password0 string) bool {
	// var (
	// 	account entity.Account = entity.Account{
	// 		time.Now().UnixNano(), email, password, "", "", "",
	// 	}
	// )
	//result := as.accountMapper.InsertAccount(account)
	// if result == 1 {
	// 	return true
	// }
	return false
}
