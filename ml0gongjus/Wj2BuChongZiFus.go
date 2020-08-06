package ml0gongjus

import (
	"strings"
	"xm1shengcheng/ml1changliangs"
)

func BuChongChuan(geShu int,ziFu string)string {
	if(geShu < ml1changliangs.Sz1){
		return ml1changliangs.KongZiFu
	}
	retshuzu := []string{}
	for i := ml1changliangs.Sz0 ;i < geShu ; i++{
		retshuzu = append(retshuzu, ziFu)
	}
	return strings.Join(retshuzu, ml1changliangs.KongZiFu)
}

func ZuoBuChong(geShu int,ziFu string,daiBuChong string) string {
	buChong:=BuChongChuan(geShu,ziFu)
	ret := buChong+daiBuChong
	return ret
}
func YouBuChong(geShu int,ziFu string,daiBuChong string) string {
	buChong:=BuChongChuan(geShu,ziFu)
	ret := daiBuChong+buChong
	return ret
}
func ZuoBuQi(zongWeiShu int,ziFu string,daiBuChong string) string {
	daiBuChongWeiShu:=len(daiBuChong)
	if(daiBuChongWeiShu >= zongWeiShu){
		return daiBuChong
	}
	geShu:=zongWeiShu-daiBuChongWeiShu
	buChong:=BuChongChuan(geShu,ziFu)
	ret := buChong+daiBuChong
	return ret
}
func YouBuQi(zongWeiShu int,ziFu string,daiBuChong string) string {
	daiBuChongWeiShu:=len(daiBuChong)
	if(daiBuChongWeiShu >= zongWeiShu){
		return daiBuChong
	}
	geShu:=zongWeiShu-daiBuChongWeiShu
	buChong:=BuChongChuan(geShu,ziFu)
	ret := daiBuChong+buChong
	return ret
}

func ShuZiZiFuBu1Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz1,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu2Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz2,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu3Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz3,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu4Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz4,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu5Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz5,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu6Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz6,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu7Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz7,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu8Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz8,ml1changliangs.Zf0,ziFu)
}
func ShuZiZiFuBu9Ge0(ziFu string)string{
	return YouBuChong(ml1changliangs.Sz9,ml1changliangs.Zf0,ziFu)
}
