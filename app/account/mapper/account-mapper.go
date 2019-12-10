package mapper

import (
	"github.com/Scfy-Code/IM/app/account/entity"
	"github.com/Scfy-Code/IM/sys"
)

type accountMapper struct {
	account entity.Entity
}

// NewAccountMapper 新建数据库交互对象
func NewAccountMapper() AccountMapper {
	return accountMapper{
		entity.NewAccount(),
	}
}
func (am accountMapper) SelectAccount(email, password string) map[string]interface{} {
	rows, err := sys.GetSQLClient("US").Query("select email,avatar,signature,nickName,life from account where email=? and password=?", email, password)
	if err != nil {
		return nil
	}
	if rows.Next() {
		rows.Scan(am.account.GetFields()...)
	}
	return am.account.EntityToMap()
}
