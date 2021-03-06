package data

type CreateRequest struct {
	Username  string
	Password  string
	Level     uint64
	QQ        string
	Wechat    string
	Cellphone string
	Email     string
	Age       uint64
}

type DeletesRequest struct {
	Ids []uint64
}

type UpdateRequest struct {
	Id uint64
}

type GetsRequest struct {
	Ids       []uint64
	Username  string
	Level     uint64
	QQ        string
	Wechat    string
	Cellphone string
	Email     string
	Age       uint64
}

type AuthRequest struct {
	Username string
	Password string
}
