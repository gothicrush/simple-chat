package message

import (
	"encoding/json"
)

func PackageToMessage(data []byte) (*Message,error) {

	var msg Message

	err := json.Unmarshal(data, &msg)

	if err != nil {
		return nil,err
	}

	return &msg,nil
}

func PackageToMessageLogin(data []byte) (*MessageLogin, error) {

	// 创建MessageLogin结构体
	var messageLogin MessageLogin

	// 反序列化为MessageLogin
	err := json.Unmarshal(data, &messageLogin)

	if err != nil {
		return nil, err
	}

	return &messageLogin, nil
}

func PackageToMessageLoginResponse(data []byte) (*MessageLoginResponse, error) {

	// 创建MessageLoginResponse结构体
	var messageLoginResponse MessageLoginResponse

	// 反序列化为MessageLoginResponse
	err := json.Unmarshal(data, &messageLoginResponse)

	if err != nil {
		return nil, err
	}

	return &messageLoginResponse, nil
}

func MessageToPackage(msg *Message) ([]byte, error) {

	messagePackage, err := json.Marshal(*msg)

	if err != nil {
		return nil, err
	}

	return messagePackage, nil
}

func MessageLoginToPackage(msglog *MessageLogin) ([]byte, error) {

	messageLoginPackage, err := json.Marshal(*msglog)

	if err != nil {
		return nil, err
	}

	return messageLoginPackage, nil

}

func MessageLoginResponseToPackage(msglogret *MessageLoginResponse) ([]byte, error) {

	messageLoginReturnPackage, err := json.Marshal(*msglogret)

	if err != nil {
		return nil, err
	}

	return messageLoginReturnPackage, nil
}
