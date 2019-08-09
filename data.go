package gomnik

import (
	"bytes"
	"encoding/binary"
)

type Request struct {
	Header   [4]byte
	RevSN1   [4]byte
	RevSN2   [4]byte
	Sep      [2]byte
	Checksum byte
	Trailer  byte
}

func (r Request) Bytes() (b []byte, err error) {
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, r)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}

func NewRequest(serial int) (req Request) {
	req = Request{Header: [4]byte{0x68, 0x02, 0x40, 0x30}, Sep: [2]byte{0x01, 0x00}, Trailer: 0x16}

	revSn := [4]byte{byte(serial & 0x000000ff), byte((serial & 0x0000ff00) >> 8), byte((serial & 0x00ff0000) >> 16), byte((serial & 0xff000000) >> 24)}
	req.RevSN1 = revSn
	req.RevSN2 = revSn

	var cs byte = 115
	for _, i := range revSn {
		cs += i * 2
	}
	req.Checksum = cs

	return
}

func NewResponse(b []byte) (r Response, err error) {
	return decodeResponse(b)
}

type Response struct {
	Header       [4]byte
	RevSN        [8]byte
	_            [3]byte
	ID           [16]byte
	Temp         uint16
	PVVoltage1   uint16
	PVVoltage2   uint16
	PVVoltage3   uint16
	PVCurrent1   uint16
	PVCurrent2   uint16
	PVCurrent3   uint16
	ACCurrent1   uint16
	ACCurrent2   uint16
	ACCurrent3   uint16
	ACVoltage1   uint16
	ACVoltage2   uint16
	ACVoltage3   uint16
	ACFrequency1 uint16
	ACPower1     uint16
	ACFrequency2 uint16
	ACPower2     uint16
	ACFrequency3 uint16
	ACPower3     uint16
	EToday       uint16
	ETotal       uint32
	HTotal       uint32
}

func decodeResponse(r []byte) (resp Response, err error) {
	buf := bytes.NewBuffer(r)
	err = binary.Read(buf, binary.BigEndian, &resp)
	return
}

/*
	var b []byte
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.LittleEndian, req)
	b = buf.Bytes()

	fmt.Printf("% x\n", b)
*/
