package entry

type UserEntry struct {
	ID         string
	RemarkName string
	Avatar     string
	ChatList   []map[string]interface{}
	FriendList []map[string]interface{}
	StartList  []map[string]interface{}
}
