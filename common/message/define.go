package message

const (
	TypeMessageLogin = "TypeMessageLogin"
	TypeMessageLoginResponse = "TypeMessageLoginResponse"
)

type Message struct {
	Type string `json:"message_type"`// 消息的类型
	Data string `json:"message_data"`// 消息
}

type MessageLogin struct {
	UserId string `json:"user_id"`
	Password string `json:"password"`
}

type MessageLoginResponse struct {
	Code int `json:"code"`
	Info string `json:"info"`
}
