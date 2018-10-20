package message

const (
	TypeOfLoginMsg = "LoginMsgType"
	TypeOfLoginResMsg = "LoginResMsgType"
)

type Message struct {
	Type string `json:"message_type"`// 消息的类型
	Data string `json:"message_data"`// 消息
}

type MsgLogin struct {
	UserId string `json:"user_id"`
	Password string `json:"password"`
	UesrName string `json:"user_name"`
}

type MsgLoginRes struct {
	Code int `json:"code"`
	ErrorInfo string `json:"error_info"`
}
