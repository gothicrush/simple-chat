package utils

import (
	"encoding/binary"
	"encoding/json"
	"github.com/gothicrush/simple-chat/common/message"
	"io"
	"log"
	"net"
)

func ReadPkg(conn net.Conn) ([]byte, error) {

	// 准备接收长度包的缓冲区
	buf := make([]byte, 4)

	_, err := conn.Read(buf[0:4])
	if err != nil && err != io.EOF {
		log.Println("readPkgLenError: ", err)
		return nil, err
	}

	// 获取数据包的长度
	pkgLen := int(binary.BigEndian.Uint32(buf[0:4]))

	// 准备接收实际的数据包的缓冲区

	buf = make([]byte, pkgLen)

	n, err := conn.Read(buf[0:pkgLen])
	if n != pkgLen || (err != nil && err != io.EOF) {
		return nil, err
	}

	return buf,nil
}

func WritePkg(conn net.Conn, data []byte) error {

	pkgLen := uint32(len(data))
	var pkgLenByte []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(pkgLenByte[0:4],pkgLen)

	_, err := conn.Write(pkgLenByte)

	if err != nil {
		return err
	}

	n, err := conn.Write(data)

	if n != int(pkgLen) && err != nil {
		return err
	}

	return nil
}

func BytesToMessage(data []byte) (*message.Message,error) {

	var msg message.Message

	err := json.Unmarshal(data, &msg)

	if err != nil {
		return nil,err
	}

	return &msg,nil
}