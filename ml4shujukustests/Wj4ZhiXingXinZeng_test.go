package ml4shujukustests
import(
  "testing"
  "xm1shengcheng/ml3shujukus"
  "log"
)

func TestZhiXingXinZeng(t *testing.T){
  sql:=""
  values:=[]string{
    "",
  }
  ret := ml3shujukus.ZhiXingXinZeng(sql,values)
  log.Println("Wj3ZhiXingXinZeng_test.go---TestZhiXingXinZeng,ret",ret)
}