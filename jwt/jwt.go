package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	var (
		plainHeader  string
		plainPayload string
		plainSecret  string
	)

	encoder := base64.RawURLEncoding

	plainHeader = `{"alg":"HS256","typ":"JWT"}`
	base64EncodedHeader := encoder.EncodeToString([]byte(plainHeader))
	// fmt.Println(base64EncodedHeader)

	plainPayload = `{"sub":"1234567890","name":"John Doe","iat":1516239022}`
	base64EncodedPayload := encoder.EncodeToString([]byte(plainPayload))
	// fmt.Println(base64EncodedPayload)

	plainSecret = "01234567890123456789012345678901" // 256bit
	signature := sign(base64EncodedHeader, base64EncodedPayload, plainSecret)
	base64EncodedSignature := encoder.EncodeToString(signature)
	// fmt.Println(base64EncodedSignature)

	jwt := jwt(base64EncodedHeader, base64EncodedPayload, base64EncodedSignature)
	fmt.Println(jwt)
}

func sign(header, payload, secret string) []byte {
	builded := header + "." + payload

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(builded))
	return mac.Sum(nil)
}

func jwt(header, payload, signature string) string {
	return header + "." + payload + "." + signature
}
