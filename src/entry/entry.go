package entry

// IndexEntry 首页
type IndexEntry struct {
	UserEntry
	TalkList   []UserEntry
	FriendList []UserEntry
}

// UserEntry 用户
type UserEntry struct {
	ID         string
	Account    string
	Password   string
	Email      string
	RemarkName string
	Avatar     string
}
