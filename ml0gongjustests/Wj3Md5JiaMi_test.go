package ml0gongjustests

import (
  "testing"
  "xm1shengcheng/ml0gongjus"
  "log"
)
func TestMd5JiaMi(t *testing.T){
  ret := ml0gongjus.Md5JiaMi("123abc")
  log.Println(ret,len(ret))
}
