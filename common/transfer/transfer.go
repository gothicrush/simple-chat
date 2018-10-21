package transfer

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Transfer struct {
	Conn net.Conn
}

func (tr *Transfer)AcceptPackage() ([]byte, error) {

	// 准备接收长度包的缓冲区
	buf := make([]byte, 4)

	// 从连接中读取数据
	_, err := tr.Conn.Read(buf[0:4])

	if err != nil && err != io.EOF {
		log.Println("transfer.AcceptPackage Error 001 : ", err)
		return nil, err
	}

	// 获取数据包的长度
	pkgLen := int(binary.BigEndian.Uint32(buf[0:4]))

	// 准备接收实际数据包的缓冲区
	buf = make([]byte, pkgLen)

	// 从连接中读取数据
	n, err := tr.Conn.Read(buf[0:pkgLen])
	if n != pkgLen || (err != nil && err != io.EOF) {
		log.Println("transfer.AcceptPackage Error 002 : ", err)
		return nil, err
	}

	return buf,nil
}

func (tr *Transfer)SendPackage(data []byte) error {

	// 获取实际数据包的长度
	packageLength := len(data)

	// 将实际数据包长度以 []byte 表示
	pkgLenByte := make([]byte, 4)
	binary.BigEndian.PutUint32(pkgLenByte[0:4],uint32(packageLength))

	// 通过连接发送长度数据包
	_, err := tr.Conn.Write(pkgLenByte)

	if err != nil {
		log.Println("transfer.SendPackage Error 001 : ", err)
		return err
	}

	// 通过连接发送实际数据包
	n, err := tr.Conn.Write(data)

	if n != packageLength && err != nil {
		log.Println("transfer.SendPackage Error 002 : ", err)
		return err
	}

	return nil
}
