package ml0gongjustests

import(
  "testing"
  "log"
  "time"
  "xm1shengcheng/ml0gongjus"
  "xm1shengcheng/ml1changliangs"
)
func TestString2Time(t *testing.T){
  ret := ml0gongjus.String2Time("2020-03-16 20:20:20",ml1changliangs.NyrSfm)
  log.Println(ret)
}
func TestString2TimeNyrSfm(t *testing.T){
  ret := ml0gongjus.String2TimeNyrSfm("2020-03-16 20:20:20")
  log.Println(ret)
}
func TestString2TimeNyr(t *testing.T){
  ret := ml0gongjus.String2TimeNyr("2020-03-16")
  log.Println(ret,"后续记得加上时间格式的正则判断，如果入参格式错误直接返回空")
}
func TestTime2String(t *testing.T){
  ti:=time.Date(9999,12,31,23,59,59,1000,time.Local)
  ret := ml0gongjus.Time2String(ti,ml1changliangs.NyrSfm)
  log.Println(ret)
}
func TestTime2StringNyr(t *testing.T){
  ti:=time.Now()
  ret := ml0gongjus.Time2StringNyr(ti)
  log.Println(ret)
}
func TestTime2StringNyrSfm(t *testing.T){
  ti:=time.Now()
  ret := ml0gongjus.Time2StringNyrSfm(ti)
  log.Println(ret)
}
func TestTime2StringNyrSfmXhx(t *testing.T){
  ti:=time.Now()
  ret := ml0gongjus.Time2StringNyrSfmXhx(ti)
  log.Println(ret)
}
func TestTime2StringNyrXhx(t *testing.T){
  ti:=time.Now()
  ret := ml0gongjus.Time2StringNyrXhx(ti)
  log.Println(ret)
}
func TestTime2StringNyrWu(t *testing.T){
  ti:=time.Now()
  ret := ml0gongjus.Time2StringNyrWu(ti)
  log.Println(ret)
}
