package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"strings"
)

type DNSHeader struct {
	ID             uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

type DNSQuestion struct {
	Name  []byte
	Type  uint16
	Class uint16
}

func convertHeaderToBytes(header DNSHeader) []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:2], header.ID)
	binary.BigEndian.PutUint16(buf[2:4], header.Flags)
	binary.BigEndian.PutUint16(buf[4:6], header.NumQuestions)
	binary.BigEndian.PutUint16(buf[6:8], header.NumAnswers)
	binary.BigEndian.PutUint16(buf[8:10], header.NumAuthorities)
	binary.BigEndian.PutUint16(buf[10:12], header.NumAdditionals)
	return buf
}

func convertQuestionToBytes(question DNSQuestion) []byte {
	return append(question.Name,
		byte(question.Type>>8), byte(question.Type),
		byte(question.Class>>8), byte(question.Class),
	)
}

func encodeDNSName(domainName string) []byte {
	encoded := []byte{}
	for _, part := range strings.Split(domainName, ".") {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, []byte(part)...)
	}
	encoded = append(encoded, 0)
	return encoded
}

const type_A = 1
const class_IN = 1

func buildQuery(domainName string, recordType uint16) []byte {
	name := encodeDNSName(domainName)
	id := rand.Intn(65536)
	recursionDesired := 1 << 8
	header := DNSHeader{
		ID:           uint16(id),
		NumQuestions: 1,
		Flags:        uint16(recursionDesired),
	}
	question := DNSQuestion{
		Name:  name,
		Type:  recordType,
		Class: class_IN,
	}
	return append(convertHeaderToBytes(header), convertQuestionToBytes(question)...)

}

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	query := buildQuery("www.example.com", type_A)
	_, err = conn.Write(query)
	if err != nil {
		panic(err)
	}

	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q\n", response)
}
