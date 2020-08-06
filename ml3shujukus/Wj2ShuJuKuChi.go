package ml3shujukus

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"sync"
)

func init() {
	log.Println("初始化连接池开始")
	lianjiechenggong := huoQuShiLi().chuShiHuaChi()
	if !lianjiechenggong {
		log.Println("初始化连接池失败...")
		os.Exit(1)
	}
}

func LianJieChi() *gorm.DB {
	db := huoQuShiLi().huoQuShuJuChi()
	return db
}

type GoHouTaiFuWuChi struct {
}

var danShiLi *GoHouTaiFuWuChi
var suoShiLi sync.Once
var lianJieChi *gorm.DB
var lianJieChiCuoWu error

func huoQuShiLi() *GoHouTaiFuWuChi {
	suoShiLi.Do(func() {
		danShiLi = &GoHouTaiFuWuChi{}
	})
	return danShiLi
}

func (m *GoHouTaiFuWuChi) chuShiHuaChi() bool {
	lianJieChi, lianJieChiCuoWu = gorm.Open("sqlite3", "../hfx.db")
	if lianJieChiCuoWu != nil {
		log.Fatal(lianJieChiCuoWu)
		return false
	}
	lianJieChi.DB().SetMaxOpenConns(50)
	lianJieChi.SingularTable(true)
	return true
}
func (m *GoHouTaiFuWuChi) huoQuShuJuChi() *gorm.DB {
	return lianJieChi
}
