package main

import (
	"encoding/binary"
	"log"
	"testing"

	"github.com/robinson/gos7"
)

func Test_s7(t *testing.T) {
	handler := gos7.NewTCPClientHandler("127.0.0.1:1503", 0, 1)
	err1 := handler.Connect()
	defer handler.Close()
	if err1 != nil {
		log.Println(err1.Error())
		return
	}
	client := gos7.NewClient(handler)

	buf := make([]byte, 2)
	buf[0] = 24
	buf[1] = 34
	println("SEND ", binary.BigEndian.Uint16(buf))

	if err := client.AGWriteDB(13, 4, 2, buf); err != nil {
		t.Error(err)
	}
	buf = make([]byte, 2)
	println("DELETE  BUFFER ", binary.BigEndian.Uint16(buf))

	if err := client.AGReadDB(13, 4, 2, buf); err != nil {
		t.Error(err)
	}
	println("READ ", binary.BigEndian.Uint16(buf))

}