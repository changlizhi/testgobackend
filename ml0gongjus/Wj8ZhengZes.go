package ml0gongjus

import (
	"log"
	"regexp"
	"xm1shengcheng/ml1changliangs"
)

func ShiFouPiPei(zhengZeBiaoDaShi,ziFu string)string{
	re, err := regexp.Compile(zhengZeBiaoDaShi)
	if err != nil {
		log.Println("正则表达式错误", err)
		return ml1changliangs.Zf0
	}
	ret := re.MatchString(ziFu)
	if ret == true {
		return ml1changliangs.Zf1
	}
	return ml1changliangs.Zf0
}

func ShouJiHaoPiPei(shoujihao string) string {
	sjh := `^1\d{10}$`
	return ShiFouPiPei(sjh,shoujihao)
}
