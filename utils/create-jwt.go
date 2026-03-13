package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type PayLoad struct {
	Sub         int `json:"sub"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"owner"`
}

func CreateJwt(secret string, data PayLoad) (string, error) {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerBytesArr, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64Encode(headerBytesArr)
	payloadBytesArr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64Encode(payloadBytesArr)

	message := headerB64 + "." + payloadB64

	byteSecret := []byte(secret)
	bytemessage := []byte(message)

	h := hmac.New(sha256.New, byteSecret)
	h.Write(bytemessage)
	signature := h.Sum(nil)
	signatureB64 := base64Encode(signature)

	jwtToken := headerB64 + "." + payloadB64 + "." + signatureB64
	return jwtToken, nil

}

func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
