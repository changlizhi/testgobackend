package ml0gongjustests

import(
  "testing"
  "xm1shengcheng/ml0gongjus"
  "log"
)

func TestDuQuMoXingJson(t *testing.T){
  ret := ml0gongjus.DuQuMoXingJson()
  for _,r:=range ret{
    log.Println("Wj10MoxingJson_test.go---TestDuQuMoXingJson,r",r)
  }
}
