package ml0gongjus

import (
	"math/rand"
)

var	ShuZiHeDaXieZiMu = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var	ShuZiHeXiaoXieZiMu = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
var	ShuZiHeZiMu = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var	ShuZi = []rune("0123456789")

func SuiJiZiFuChuan(n int, allowedChars ...[]rune) string {
	var letters []rune=allowedChars[0]

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
func SuiJiShuZi(len int) string {
	return SuiJiZiFuChuan(len,ShuZi)
}
func SuiJiDaXieZiMuHeShuZi(len int) string {
	return SuiJiZiFuChuan(len,ShuZiHeDaXieZiMu)
}
func SuiJiXiaoXieZiMuHeShuZi(len int) string {
	return SuiJiZiFuChuan(len,ShuZiHeXiaoXieZiMu)
}
func SuiJiZiMuHeShuZi(len int) string {
	return SuiJiZiFuChuan(len,ShuZiHeZiMu)
}
