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
type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(p Payload, secret string) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", nil
	}
	headerB64 := Base64UrlEncode(byteArrHeader)
	byteArrPayload, err := json.Marshal(p)
	if err != nil {
		return "", nil
	}
	payloadB64 := Base64UrlEncode(byteArrPayload)
	message := headerB64 + "." + payloadB64
	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	sign := h.Sum(nil)
	signB64 := Base64UrlEncode(sign)
	jwt := headerB64 + "." + payloadB64 + "." + signB64
	return jwt, nil

}


func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
