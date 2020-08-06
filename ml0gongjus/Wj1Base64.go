package ml0gongjus

import (
	"encoding/base64"
	"log"
)

func Base64BianMa(ZiFu string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(ZiFu))
	return encodeString
}
func Base64JieMa(ZiFu string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(ZiFu)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(decodeBytes)
}
func Base64UrlBianMa(ZiFu string) string {
	ret := base64.URLEncoding.EncodeToString([]byte(ZiFu))
	return ret
}
func Base64UrlJiema(ZiFu string) string {
	decodeBytes, err := base64.URLEncoding.DecodeString(ZiFu)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(decodeBytes)
}
