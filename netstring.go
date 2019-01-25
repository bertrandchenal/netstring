package netstring

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type NetString struct {
	buffer bytes.Buffer
	err    error
}

// Netstring constructors
func NewNetBytes(values ...[]byte) *NetString {
	var buffer bytes.Buffer
	for _, val := range values {
		buffer.Write(val)
	}
	return &NetString{buffer, nil}
}

func NewNetString(values ...string) *NetString {
	var buffer bytes.Buffer
	for _, val := range values {
		buffer.WriteString(val)
	}
	return &NetString{buffer, nil}
}

// NetString encoding
func (self *NetString) Encode(items ...[]byte) {
	tail := ","
	for _, item := range items {
		head := fmt.Sprintf("%d:", len(item))
		_, err := self.buffer.WriteString(head)
		_, err = self.buffer.Write(item)
		_, err = self.buffer.WriteString(tail)
		if err != nil && self.err != nil {
			self.err = err
			return
		}
	}
}

func (self *NetString) EncodeString(items ...string) {
	for _, item := range items {
		self.Encode([]byte(item))
	}
}

// Netstring decoding
func (self *NetString) Decode() [][]byte {
	head, err := self.buffer.ReadBytes(byte(':'))
	if err == io.EOF {
		return make([][]byte, 0)
	}
	// Read header giving item size
	length, err := strconv.ParseInt(string(head[:len(head)-1]), 10, 32)
	if err != nil {
		self.err = err
		return nil
	}
	// Read payload
	payload := make([]byte, length)
	_, err = self.buffer.Read(payload)
	if err != nil {
		self.err = err
		return nil
	}
	res := [][]byte{payload}
	// Read end delimiter
	delim, err := self.buffer.ReadByte()
	if err != nil {
		self.err = err
		return nil
	}
	if delim != byte(',') {
		self.err = errors.New("Unable de decode netstring, unexpected end of stream")
		return nil
	}
	// Recurse
	tail := self.Decode()
	if self.err != nil {
		return nil
	}

	return append(res, tail...)
}

func (self *NetString) DecodeString() []string {
	res_bytes := self.Decode()
	res := make([]string, len(res_bytes))
	for pos, val := range res_bytes {
		res[pos] = string(val)
	}
	return res
}

func (self *NetString) Bytes() []byte {
	return self.buffer.Bytes()
}

func Decode(in []byte) ([][]byte, error) {
	ns := NewNetBytes(in)
	out := ns.Decode()
	return out, ns.err
}

func DecodeString(in []byte) ([]string, error) {
	ns := NewNetBytes(in)
	out := ns.DecodeString()
	return out, ns.err
}


func Encode(items ...[]byte) ([]byte, error) {
	ns := NewNetBytes()
	ns.Encode(items...)
	return ns.buffer.Bytes(), ns.err
}

func EncodeString(items ...string) ([]byte, error) {
	ns := NewNetString()
	ns.EncodeString(items...)
	return ns.buffer.Bytes(), ns.err
}
