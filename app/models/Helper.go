package models

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func JsonToArray(jsondata []byte) map[string]interface{} {
	var err error
	var msgMapTemplate interface{}
	err = json.Unmarshal(jsondata, &msgMapTemplate)
	msgMap := msgMapTemplate.(map[string]interface{})
	if err != nil {
		panic(err)
	}

	return msgMap
}

func JsonToMultiArray(jsondata []byte) []interface{} {
	var err error
	var f interface{}
	err = json.Unmarshal(jsondata, &f)
	o := f.([]interface{})
	if err != nil {
		panic(err)
	}

	return o
}

func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return bytes, err
	}
	return data, err
}

func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

type Header struct {
	branchid string `header:"branchid" `
}
